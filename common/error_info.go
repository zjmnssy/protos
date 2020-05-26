package common

import fmt "fmt"

// ErrorMsg 错误提示信息表
var ErrorMsg map[ErrorCode]string

func init() {
	ErrorMsg = make(map[ErrorCode]string)
	for k := range ErrorCode_name {
		ErrorMsg[ErrorCode(k)] = fmt.Sprintf("内部错误(%d)", k)
	}

	ErrorMsg[ErrorCode_Success] = "成功"

	ErrorMsg[ErrorCode_WebSuccess] = "成功"
	ErrorMsg[ErrorCode_WebMethodError] = "method error"
	ErrorMsg[ErrorCode_WebRouteNotFound] = "route not found"
	ErrorMsg[ErrorCode_WebCookieOntFound] = "cookie missing"
	ErrorMsg[ErrorCode_WebCookieIllegal] = "cookie illegal"
	ErrorMsg[ErrorCode_WebMissingOfCSRF] = "非法访问"
	ErrorMsg[ErrorCode_WebIllegalAccessOfCSRF] = "非法访问"

	ErrorMsg[ErrorCode_SystemBusy] = "系统繁忙"
	ErrorMsg[ErrorCode_CredentialMissing] = "请先登录"
	ErrorMsg[ErrorCode_CredentialParseError] = "拒绝访问"
	ErrorMsg[ErrorCode_GetInfoFromCredentialError] = "拒绝访问"
	ErrorMsg[ErrorCode_CredentialExpired] = "凭据已过期，请重新登录"
	ErrorMsg[ErrorCode_UsernameOrPasswordError] = "用户名或密码错误"
	ErrorMsg[ErrorCode_RecommendCodeError] = "推荐码错误"
	ErrorMsg[ErrorCode_VerificationCodeError] = "验证码错误"
	ErrorMsg[ErrorCode_UsernameHasExists] = "用户名已存在"
	ErrorMsg[ErrorCode_PhoneHasExists] = "该手机号码已注册"
	ErrorMsg[ErrorCode_EmailHasExists] = "该邮箱已注册"
	ErrorMsg[ErrorCode_UsernameFormatError] = "用户名格式不符合要求"
	ErrorMsg[ErrorCode_EmailFormatError] = "邮箱格式不符合要求"
	ErrorMsg[ErrorCode_PhoneFormatError] = "手机号格式不符合要求"
	ErrorMsg[ErrorCode_ParameterError] = "请求参数错误"
	ErrorMsg[ErrorCode_ConfigurationExists] = "配置已存在或冲突"

	ErrorMsg[ErrorCode_NetworkCommonError] = "网络错误"
	ErrorMsg[ErrorCode_NetworkDeadline] = "网络超时"
	ErrorMsg[ErrorCode_NetworkConnectionRefused] = " 网络连接被拒绝"
	ErrorMsg[ErrorCode_NetworkDestUnreachable] = "无法连接到目标主机"
}
