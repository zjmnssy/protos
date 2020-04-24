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

	ErrorMsg[ErrorCode_SystemBusy] = "系统繁忙"
	ErrorMsg[ErrorCode_CredentialMissing] = "请先登录"
	ErrorMsg[ErrorCode_CredentialParseError] = "拒绝访问"
	ErrorMsg[ErrorCode_GetInfoFromCredentialError] = "拒绝访问"
	ErrorMsg[ErrorCode_CredentialExpired] = "已过期，请重新登录"
	ErrorMsg[ErrorCode_UsernameOrPasswordError] = "用户名或密码错误"
	ErrorMsg[ErrorCode_RecommendCodeError] = "推荐码错误"
	ErrorMsg[ErrorCode_VerificationCodeError] = "验证码错误"
	ErrorMsg[ErrorCode_UsernameHasExists] = "用户名已存在"
	ErrorMsg[ErrorCode_PhoneHasExists] = "手机号码已注册过"
	ErrorMsg[ErrorCode_EmailHasExists] = "邮箱已注册过"

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
