package utils

// IsStringInSlice description
//
// createTime:2019年05月21日 15:50:15
// author:hailaz
func IsStringInSlice(str string, strs []string) bool {
	for _, item := range strs {
		if item == str {
			return true
		}
	}
	return false
}
