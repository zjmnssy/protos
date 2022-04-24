package go_proto_pkg

import (
	"fmt"

	"go_proto_pkg/common"
)

// ErrorMsgZh 中文错误提示信息表
var ErrorMsgZh map[common.ErrorCode]string

// ErrorMsgZh 英文错误提示信息表
var ErrorMsgEn map[common.ErrorCode]string

func init() {
	ErrorMsgZh = make(map[common.ErrorCode]string)
	ErrorMsgEn = make(map[common.ErrorCode]string)

	for k := range common.ErrorCode_name {
		ErrorMsgZh[common.ErrorCode(k)] = fmt.Sprintf("内部错误(%d)", k)
		ErrorMsgEn[common.ErrorCode(k)] = fmt.Sprintf("internal error(%d)", k)
	}

	ErrorMsgZh[common.ErrorCode_Success] = "成功"

	ErrorMsgZh[common.ErrorCode_WebSuccess] = "成功"
	ErrorMsgZh[common.ErrorCode_WebMethodError] = "web method错误"
	ErrorMsgZh[common.ErrorCode_WebRouteNotFound] = "未找到匹配的路由"
	ErrorMsgZh[common.ErrorCode_WebCookieOntFound] = "cookie缺失"
	ErrorMsgZh[common.ErrorCode_WebCookieIllegal] = "cookie错误"
	ErrorMsgZh[common.ErrorCode_WebMissingOfCSRF] = fmt.Sprintf("非法访问(%d)", common.ErrorCode_WebMissingOfCSRF)
	ErrorMsgZh[common.ErrorCode_WebIllegalAccessOfCSRF] = fmt.Sprintf("非法访问(%d)", common.ErrorCode_WebIllegalAccessOfCSRF)

	ErrorMsgZh[common.ErrorCode_SystemBusy] = "系统繁忙, 请稍后重试"
	ErrorMsgZh[common.ErrorCode_CredentialMissing] = "请先登录"
	ErrorMsgZh[common.ErrorCode_CredentialParseError] = fmt.Sprintf("拒绝访问(%d)", common.ErrorCode_CredentialParseError)
	ErrorMsgZh[common.ErrorCode_GetInfoFromCredentialError] = fmt.Sprintf("拒绝访问(%d)", common.ErrorCode_GetInfoFromCredentialError)
	ErrorMsgZh[common.ErrorCode_CredentialExpired] = "凭据已过期,请重新登录"
	ErrorMsgZh[common.ErrorCode_UsernameOrPasswordError] = "用户名或密码错误"
	ErrorMsgZh[common.ErrorCode_RecommendCodeError] = "推荐码错误"
	ErrorMsgZh[common.ErrorCode_VerificationCodeError] = "验证码错误"
	ErrorMsgZh[common.ErrorCode_UsernameHasExists] = "用户名已存在"
	ErrorMsgZh[common.ErrorCode_PhoneHasExists] = "该手机号码已注册"
	ErrorMsgZh[common.ErrorCode_EmailHasExists] = "该邮箱已注册"
	ErrorMsgZh[common.ErrorCode_UsernameFormatError] = "用户名格式不符合要求"
	ErrorMsgZh[common.ErrorCode_EmailFormatError] = "邮箱格式不符合要求"
	ErrorMsgZh[common.ErrorCode_PhoneFormatError] = "手机号格式不符合要求"
	ErrorMsgZh[common.ErrorCode_ParameterError] = "请求参数错误"
	ErrorMsgZh[common.ErrorCode_ConfigurationExists] = "配置已存在或冲突"

	ErrorMsgZh[common.ErrorCode_NetworkCommonError] = "网络错误"
	ErrorMsgZh[common.ErrorCode_NetworkDeadline] = "网络超时"
	ErrorMsgZh[common.ErrorCode_NetworkConnectionRefused] = " 网络连接被拒绝"
	ErrorMsgZh[common.ErrorCode_NetworkDestUnreachable] = "无法连接到目标主机"

	/****************************************************************************************/

	ErrorMsgEn[common.ErrorCode_Success] = "success"

	ErrorMsgEn[common.ErrorCode_WebSuccess] = "success"
	ErrorMsgEn[common.ErrorCode_WebMethodError] = "web method invalid"
	ErrorMsgEn[common.ErrorCode_WebRouteNotFound] = "route missing"
	ErrorMsgEn[common.ErrorCode_WebCookieOntFound] = "cookie missing"
	ErrorMsgEn[common.ErrorCode_WebCookieIllegal] = "cookie invalid"
	ErrorMsgEn[common.ErrorCode_WebMissingOfCSRF] = fmt.Sprintf("invalid access(%d)", common.ErrorCode_WebMissingOfCSRF)
	ErrorMsgEn[common.ErrorCode_WebIllegalAccessOfCSRF] = fmt.Sprintf("invalid access(%d)", common.ErrorCode_WebIllegalAccessOfCSRF)

	ErrorMsgEn[common.ErrorCode_SystemBusy] = "system busy, please try again later"
	ErrorMsgEn[common.ErrorCode_CredentialMissing] = "please login fist"
	ErrorMsgEn[common.ErrorCode_CredentialParseError] = fmt.Sprintf("deny access(%d)", common.ErrorCode_CredentialParseError)
	ErrorMsgEn[common.ErrorCode_GetInfoFromCredentialError] = fmt.Sprintf("deny access(%d)", common.ErrorCode_GetInfoFromCredentialError)
	ErrorMsgEn[common.ErrorCode_CredentialExpired] = "credentials have expired, please log in again"
	ErrorMsgEn[common.ErrorCode_UsernameOrPasswordError] = "account or password error"
	ErrorMsgEn[common.ErrorCode_RecommendCodeError] = "referral code error"
	ErrorMsgEn[common.ErrorCode_VerificationCodeError] = "verification code error"
	ErrorMsgEn[common.ErrorCode_UsernameHasExists] = "account already exists"
	ErrorMsgEn[common.ErrorCode_PhoneHasExists] = "the mobile number is already registered"
	ErrorMsgEn[common.ErrorCode_EmailHasExists] = "the email is already registered"
	ErrorMsgEn[common.ErrorCode_UsernameFormatError] = "username format does not meet requirements"
	ErrorMsgEn[common.ErrorCode_EmailFormatError] = "email format does not meet the requirements"
	ErrorMsgEn[common.ErrorCode_PhoneFormatError] = "The phone number format does not meet the requirements"
	ErrorMsgEn[common.ErrorCode_ParameterError] = "request parameter error"
	ErrorMsgEn[common.ErrorCode_ConfigurationExists] = "configuration already exists or conflicts"

	ErrorMsgEn[common.ErrorCode_NetworkCommonError] = "network error"
	ErrorMsgEn[common.ErrorCode_NetworkDeadline] = "network timeout"
	ErrorMsgEn[common.ErrorCode_NetworkConnectionRefused] = "connection refused"
	ErrorMsgEn[common.ErrorCode_NetworkDestUnreachable] = "unable to connect to target host"
}
