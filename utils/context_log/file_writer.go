package context_log

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

const logFileNameFormat = "%s.%4d-%02d-%02d.log"

// FileWriter 日志实现Writer
type FileWriter struct {
	maxSize  int64
	maxNum   int
	fileName string
	filePath string
	file     *os.File
	writer   io.Writer
	mu       sync.Mutex
	ch       chan []byte
}

// NewFileWriter 新建一个日志writer，并启动三个goroutine来 rotate, check, flush
func NewFileWriter(fileName string, maxSize int64, maxNum int) (io.Writer, error) {
	y, m, d := time.Now().Date()
	path := fmt.Sprintf(logFileNameFormat, fileName, y, m, d)
	file, e := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if e != nil {
		return nil, e
	}
	writer := &FileWriter{fileName: fileName, filePath: path, file: file, writer: file, ch: make(chan []byte, 10000), maxSize: maxSize, maxNum: maxNum}
	go writer.rotate()
	go writer.flush()
	go writer.check()
	return writer, nil
}

// Write 异步channel写日志
func (w *FileWriter) Write(p []byte) (int, error) {
	buf := make([]byte, len(p))
	copy(buf, p)
	select {
	case w.ch <- buf:
		//(2769322, 1)       //log写入成功
		//(33520635, len(p)) //log写入channel字节数
		return len(buf), nil
	default:
		//(2769322, 1) //chan满，写入失败
		return 0, errors.New("chan full, drop")
	}
}

// check 每分钟检查一下日志文件是否存在，运维误删log文件但是进程一直在打日志，fd会一直存在，需要关闭。超过maxSize自动rotate
func (w *FileWriter) check() {
	for {
		time.Sleep(time.Minute)

		w.mu.Lock()
		fileInfo, err := os.Stat(w.filePath)
		if os.IsNotExist(err) {
			//(3216535, 1) //日志已被误删除，重新创建新日志文件
			file, e := os.OpenFile(w.filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
			if e == nil {
				w.file.Close()
				w.file = file
				w.writer = file
			}
			w.mu.Unlock()
			continue
		}
		if w.maxSize > 0 && fileInfo.Size() > w.maxSize {
			//(3216536, 1) //日志文件超过最大size
			name := path.Base(w.filePath) + ".full."
			files, _ := ioutil.ReadDir(path.Dir(w.filePath))
			var minNum = 1000000
			var maxNum = 0
			var totalNum = 0
			for _, f := range files {
				if strings.Contains(f.Name(), name) {
					totalNum++
					s := strings.Split(f.Name(), ".") // going.2018-05-22.log.full.1.log
					if len(s) > 5 {
						n, _ := strconv.Atoi(s[len(s)-2])
						if n > maxNum {
							maxNum = n
						}
						if n < minNum {
							minNum = n
						}
					}
				}
			}
			w.file.Close()
			//rename log file
			name = fmt.Sprintf("%s.full.%d.log", w.filePath, maxNum+1)
			err := os.Rename(w.filePath, name)
			if err != nil {
				//(3216537, 1) //Rename重命名日志文件失败
				fmt.Printf("rename file path:%s fail:%s\n", w.filePath, err)
			}
			file, err := os.OpenFile(w.filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
			if err != nil {
				//(3216538, 1) //创建日志文件失败
				fmt.Printf("open file path:%s fail:%s\n", w.filePath, err)
			}
			if err == nil {
				w.file = file
				w.writer = file
			}
			if totalNum >= w.maxNum {
				//(3216539, 1) //大日志文件个数超过20个
				//remove oldest log file
				name = fmt.Sprintf("%s.full.%d.log", w.filePath, minNum)
				err := os.Remove(name)
				if err != nil {
					//(3216540, 1) //Remove删除老日志文件失败
					fmt.Printf("remove file path:%s fail:%s\n", name, err)
				}
			}
		}
		w.mu.Unlock()
	}
}

// rotate 按天更新日志文件名
func (w *FileWriter) rotate() {
	for {
		now := time.Now()
		y, m, d := now.Add(24 * time.Hour).Date()
		nextDay := time.Date(y, m, d, 0, 0, 0, 0, now.Location())
		tm := time.NewTimer(time.Duration(nextDay.UnixNano() - now.UnixNano() - 100))
		<-tm.C
		w.mu.Lock()
		path := fmt.Sprintf(logFileNameFormat, w.fileName, y, m, d)
		file, e := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if e == nil {
			w.file.Close()
			w.file = file
			w.writer = file
			w.filePath = path
		}
		w.mu.Unlock()
	}
}

// flush 刷新日志到磁盘中
func (w *FileWriter) flush() {
	for {
		log := <-w.ch
		w.mu.Lock()
		w.writer.Write(log)
		w.mu.Unlock()
		//(33520636, 1)        //log写入磁盘个数
		//(33520637, len(log)) //log写入磁盘字节数
	}
}
