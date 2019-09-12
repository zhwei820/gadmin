package base

import (
	"github.com/gogf/gf-jwt"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/utils/code"
	"github.com/hailaz/gadmin/utils/context_log"
	"strconv"
	"strings"
)

type BaseController struct {
}

type BaseResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorMap map[string]string  // 后端错误信息， 放在Data里返回

func GetErrorMapForValid(validErrors map[string]map[string]string) ErrorMap {
	var res = make(ErrorMap, 0)
	for key, item:= range validErrors {
		for _, msg:= range item {
			res[key] = msg
			break
		}
	}
	return res
}

// Response API返回
//
// createTime:2019年04月25日 11:32:47
// author:hailaz
func Response(r *ghttp.Request, rs BaseResult) {
	_ = r.Response.WriteJson(rs)
	if rs.Code < 600 && rs.Code > 0 {
		r.Response.Status = rs.Code
	}
}

// Success 返回成功
//
// createTime:2019年04月25日 11:41:44
// author:hailaz
func Success(r *ghttp.Request, data interface{}) {
	Response(r, BaseResult{Code: code.RESPONSE_SUCCESS, Message: "success", Data: data})
}

// Fail 返回失败
//
// createTime:2019年04月25日 11:43:34
// author:hailaz
func Fail(r *ghttp.Request, errCode int, msg string, Data ...interface{}) {
	if msg == "" {
		msg = "fail"
	}
	if len(Data) > 0 {
		Response(r, BaseResult{Code: errCode, Message: msg, Data: Data[0]})
	} else {
		Response(r, BaseResult{Code: errCode, Message: msg})
	}

}

// funcName 获取当前用户
//
// createTime:2019年05月13日 10:01:17
// author:hailaz
func GetUser(r *ghttp.Request) *model.GadminUser {
	claims := jwt.ExtractClaims(r)
	user, _ := model.GetUserByName(claims["username"].(string))
	return user
}

func GetUserId(r *ghttp.Request) int {
	addu := GetUser(r)
	var addUserId = 0
	if addu != nil {
		addUserId = addu.Id
	}
	return addUserId
}

func ReqDebug(r *ghttp.Request, msg string, v ...interface{}) {
	r.GetParam("ctx").Val().(*context_log.Context).Debug(msg, v...)
}

func ReqInfo(r *ghttp.Request, msg string, v ...interface{}) {
	r.GetParam("ctx").Val().(*context_log.Context).Info(msg, v...)
}

func ReqWarning(r *ghttp.Request, msg string, v ...interface{}) {
	r.GetParam("ctx").Val().(*context_log.Context).Warning(msg, v...)
}

func ReqError(r *ghttp.Request, msg string, v ...interface{}) {
	r.GetParam("ctx").Val().(*context_log.Context).Error(msg, v...)
}

var (
	CONTAINS  = "__contains"
	ICONTAINS = "__icontains"
	RANGE     = "__range"
	IN        = "__in"
	GTE       = "__gte"
	LTE       = "__lte"
	GT        = "__gt"
	LT        = "__lt"
)
var FilterKeys = []string{
	CONTAINS,
	ICONTAINS,
	RANGE,
	IN,
	GTE,
	LTE,
	GT,
	LT,
}
var FilterMap = map[string]string{
	CONTAINS:  " like BINARY ? ",
	ICONTAINS: " like ? ",
	RANGE:     " between ? AND ? ",
	IN:        " in (?) ",
	GTE:       " >= ? ",
	LTE:       " <= ? ",
	GT:        " > ? ",
	LT:        " < ? ",
}

func GetWhereFromQuerys(querys map[string]interface{}) map[string]interface{} {
	wheres := make(map[string]interface{}, 0)
	for key := range querys {
		for _, kk := range FilterKeys {
			if strings.Contains(key, kk) {
				val := querys[key]

				if _, ok := querys[key].(string); ok {
					if CONTAINS == kk || ICONTAINS == kk {
						val = "%" + querys[key].(string) + "%"
					}
					if RANGE == kk || IN == kk {
						val = strings.Split(querys[key].(string), ",")
					}
				}

				wheres[key[0:len(key)-len(kk)]+FilterMap[kk]] = val
				break
			}
		}
	}
	return wheres
}

func GetWhereFromRequest(r *ghttp.Request, strKeys, intKeys, searchKeys []string) map[string]interface{} {
	retQuery := make(map[string]interface{}, 0)
	query := r.GetQueryMap()
	if strKeys != nil {

		for _, key := range strKeys {
			if item, ok := query[key]; ok {
				retQuery[key] = item
			}
		}
	}
	if intKeys != nil {
		for _, key := range intKeys {
			if item, ok := query[key]; ok {
				if strings.Contains(key, RANGE) || strings.Contains(key, IN) {
					items := strings.Split(item, ",")
					if len(items) == 2 {
						ii, _ := strconv.Atoi(items[0])
						jj, _ := strconv.Atoi(items[1])
						retQuery[key] = []int{ii, jj}
					}
				} else {
					ii, _ := strconv.Atoi(item)
					retQuery[key] = ii

				}
			}
		}
	}
	retWhere := GetWhereFromQuerys(retQuery)

	if item, ok := query["search"]; ok && searchKeys != nil {
		searchKey := make([]string, 0)
		searchValue := make([]interface{}, 0)
		for _, key := range searchKeys {
			searchKey = append(searchKey, key+" like ? ")
			searchValue = append(searchValue, "%"+item+"%")
		}
		retWhere[strings.Join(searchKey, " or ")] = searchValue
	}
	return retWhere
}
