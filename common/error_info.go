package common

import fmt "fmt"

// ErrorMsg 错误提示信息表
var ErrorMsg map[ErrorCode]string

func init() {
	ErrorMsg = make(map[ErrorCode]string)
	for k := range ErrorCode_name {
		ErrorMsg[ErrorCode(k)] = fmt.Sprintf("内部错误(%d)", k)
	}

	ErrorMsg[ErrorCode_Success] = "success"

	ErrorMsg[ErrorCode_WebSuccess] = "success"
	ErrorMsg[ErrorCode_WebMethodError] = "method error"
	ErrorMsg[ErrorCode_WebRouteNotFound] = "route not found"
	ErrorMsg[ErrorCode_WebCookieOntFound] = "cookie missing"
	ErrorMsg[ErrorCode_WebCookieIllegal] = "cookie illegal"

	ErrorMsg[ErrorCode_NetworkCommonError] = "network error"
	ErrorMsg[ErrorCode_NetworkDeadline] = "network deadline"
	ErrorMsg[ErrorCode_NetworkConnectionRefused] = "connection refused"
	ErrorMsg[ErrorCode_NetworkDestUnreachable] = "dest unreachable"
}
