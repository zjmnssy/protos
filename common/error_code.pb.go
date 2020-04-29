// Code generated by protoc-gen-go. DO NOT EDIT.
// source: error_code.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 错误码, 说明:
// 1. 错误码命名规则为: 服务类型+错误说明
type ErrorCode int32

const (
	// 常规 0
	ErrorCode_Success ErrorCode = 0
	// web 200 - 999
	ErrorCode_WebSuccess ErrorCode = 200
	// web 1000 - 1099
	ErrorCode_WebError               ErrorCode = 1000
	ErrorCode_WebMethodError         ErrorCode = 1001
	ErrorCode_WebRouteNotFound       ErrorCode = 1002
	ErrorCode_WebCookieOntFound      ErrorCode = 1003
	ErrorCode_WebCookieIllegal       ErrorCode = 1004
	ErrorCode_WebMissingOfCSRF       ErrorCode = 1005
	ErrorCode_WebIllegalAccessOfCSRF ErrorCode = 1006
	// web 1100 - 1199
	ErrorCode_SystemError                ErrorCode = 1100
	ErrorCode_SystemBusy                 ErrorCode = 1101
	ErrorCode_JSONMarshalError           ErrorCode = 1102
	ErrorCode_JSONUnmarshalError         ErrorCode = 1103
	ErrorCode_CredentialMissing          ErrorCode = 1104
	ErrorCode_CredentialParseError       ErrorCode = 1105
	ErrorCode_GetInfoFromCredentialError ErrorCode = 1106
	ErrorCode_CredentialExpired          ErrorCode = 1107
	ErrorCode_UsernameOrPasswordError    ErrorCode = 1108
	ErrorCode_RecommendCodeError         ErrorCode = 1109
	ErrorCode_VerificationCodeError      ErrorCode = 1110
	ErrorCode_UsernameHasExists          ErrorCode = 1111
	ErrorCode_PhoneHasExists             ErrorCode = 1112
	ErrorCode_EmailHasExists             ErrorCode = 1113
	// 通信 1200 - 1299
	ErrorCode_NetworkCommonError       ErrorCode = 1200
	ErrorCode_NetworkDeadline          ErrorCode = 1201
	ErrorCode_NetworkConnectionRefused ErrorCode = 1202
	ErrorCode_NetworkDestUnreachable   ErrorCode = 1203
	// sql 1300 - 1399
	ErrorCode_SqlCommonError         ErrorCode = 1300
	ErrorCode_SqlCertificationFailed ErrorCode = 1301
	ErrorCode_SqlPermissionDeny      ErrorCode = 1302
	ErrorCode_SqlQueryError          ErrorCode = 1303
	ErrorCode_SqlScanError           ErrorCode = 1304
	ErrorCode_SqlExecError           ErrorCode = 1305
	ErrorCode_SqlTxBeginError        ErrorCode = 1306
	ErrorCode_SqlTxRoobackError      ErrorCode = 1307
	ErrorCode_SqlTxCommitError       ErrorCode = 1308
)

var ErrorCode_name = map[int32]string{
	0:    "Success",
	200:  "WebSuccess",
	1000: "WebError",
	1001: "WebMethodError",
	1002: "WebRouteNotFound",
	1003: "WebCookieOntFound",
	1004: "WebCookieIllegal",
	1005: "WebMissingOfCSRF",
	1006: "WebIllegalAccessOfCSRF",
	1100: "SystemError",
	1101: "SystemBusy",
	1102: "JSONMarshalError",
	1103: "JSONUnmarshalError",
	1104: "CredentialMissing",
	1105: "CredentialParseError",
	1106: "GetInfoFromCredentialError",
	1107: "CredentialExpired",
	1108: "UsernameOrPasswordError",
	1109: "RecommendCodeError",
	1110: "VerificationCodeError",
	1111: "UsernameHasExists",
	1112: "PhoneHasExists",
	1113: "EmailHasExists",
	1200: "NetworkCommonError",
	1201: "NetworkDeadline",
	1202: "NetworkConnectionRefused",
	1203: "NetworkDestUnreachable",
	1300: "SqlCommonError",
	1301: "SqlCertificationFailed",
	1302: "SqlPermissionDeny",
	1303: "SqlQueryError",
	1304: "SqlScanError",
	1305: "SqlExecError",
	1306: "SqlTxBeginError",
	1307: "SqlTxRoobackError",
	1308: "SqlTxCommitError",
}

var ErrorCode_value = map[string]int32{
	"Success":                    0,
	"WebSuccess":                 200,
	"WebError":                   1000,
	"WebMethodError":             1001,
	"WebRouteNotFound":           1002,
	"WebCookieOntFound":          1003,
	"WebCookieIllegal":           1004,
	"WebMissingOfCSRF":           1005,
	"WebIllegalAccessOfCSRF":     1006,
	"SystemError":                1100,
	"SystemBusy":                 1101,
	"JSONMarshalError":           1102,
	"JSONUnmarshalError":         1103,
	"CredentialMissing":          1104,
	"CredentialParseError":       1105,
	"GetInfoFromCredentialError": 1106,
	"CredentialExpired":          1107,
	"UsernameOrPasswordError":    1108,
	"RecommendCodeError":         1109,
	"VerificationCodeError":      1110,
	"UsernameHasExists":          1111,
	"PhoneHasExists":             1112,
	"EmailHasExists":             1113,
	"NetworkCommonError":         1200,
	"NetworkDeadline":            1201,
	"NetworkConnectionRefused":   1202,
	"NetworkDestUnreachable":     1203,
	"SqlCommonError":             1300,
	"SqlCertificationFailed":     1301,
	"SqlPermissionDeny":          1302,
	"SqlQueryError":              1303,
	"SqlScanError":               1304,
	"SqlExecError":               1305,
	"SqlTxBeginError":            1306,
	"SqlTxRoobackError":          1307,
	"SqlTxCommitError":           1308,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c5513ac0a8e17e40, []int{0}
}

func init() {
	proto.RegisterEnum("common.ErrorCode", ErrorCode_name, ErrorCode_value)
}

func init() { proto.RegisterFile("error_code.proto", fileDescriptor_c5513ac0a8e17e40) }

var fileDescriptor_c5513ac0a8e17e40 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xcb, 0x72, 0xd3, 0x3e,
	0x14, 0xc6, 0xff, 0xff, 0x4d, 0xed, 0x9c, 0x52, 0xaa, 0x8a, 0x5e, 0xa0, 0xc0, 0xb0, 0x67, 0xc1,
	0x86, 0x27, 0xa0, 0x69, 0x02, 0x65, 0xa6, 0x4d, 0x88, 0x5b, 0xbc, 0x64, 0x64, 0xfb, 0xa4, 0xd1,
	0x54, 0x96, 0x1a, 0x49, 0x9e, 0x26, 0xef, 0xc1, 0x1d, 0x1e, 0x82, 0xcb, 0x4b, 0xb0, 0x00, 0xca,
	0xfd, 0xf2, 0x06, 0xdc, 0x9f, 0x81, 0x91, 0x65, 0x27, 0x19, 0x96, 0xf9, 0x9d, 0xef, 0x7c, 0xe7,
	0x3b, 0x39, 0x32, 0x10, 0xd4, 0x5a, 0xe9, 0x5b, 0xa9, 0xca, 0xf0, 0xd2, 0xa1, 0x56, 0x56, 0xd1,
	0xb9, 0x54, 0xe5, 0xb9, 0x92, 0x17, 0x8f, 0xe7, 0xa0, 0xd1, 0x72, 0xc5, 0xa6, 0xca, 0x90, 0xce,
	0x43, 0x10, 0x15, 0x69, 0x8a, 0xc6, 0x90, 0xff, 0xe8, 0x22, 0x40, 0x8c, 0x49, 0xfd, 0xfb, 0xc5,
	0xff, 0x74, 0x01, 0xc2, 0x18, 0x93, 0x52, 0x4d, 0xbe, 0x05, 0xf4, 0x14, 0x9c, 0x8c, 0x31, 0xd9,
	0x46, 0x3b, 0x50, 0x99, 0x87, 0xdf, 0x03, 0xba, 0x02, 0x24, 0xc6, 0xa4, 0xa7, 0x0a, 0x8b, 0x3b,
	0xca, 0xb6, 0x55, 0x21, 0x33, 0xf2, 0x23, 0xa0, 0xab, 0xb0, 0x14, 0x63, 0xd2, 0x54, 0xea, 0x80,
	0x63, 0x47, 0x56, 0xfc, 0x67, 0x2d, 0xf7, 0x7c, 0x4b, 0x08, 0xdc, 0x67, 0x82, 0xfc, 0xaa, 0xf1,
	0x36, 0x37, 0x86, 0xcb, 0xfd, 0x4e, 0xbf, 0x19, 0xf5, 0xda, 0xe4, 0x77, 0x40, 0xcf, 0xc2, 0x6a,
	0x8c, 0x49, 0xa5, 0xbb, 0x52, 0xe6, 0xaa, 0x8a, 0x7f, 0x02, 0x4a, 0x60, 0x3e, 0x1a, 0x1b, 0x8b,
	0xb9, 0xcf, 0xf2, 0x32, 0x74, 0x0b, 0x78, 0xb2, 0x51, 0x98, 0x31, 0x79, 0x15, 0x3a, 0xdb, 0xeb,
	0x51, 0x67, 0x67, 0x9b, 0x69, 0x33, 0x60, 0xc2, 0xeb, 0x5e, 0x87, 0x74, 0x0d, 0xa8, 0xc3, 0x7b,
	0x32, 0x9f, 0x2d, 0x1c, 0x87, 0x2e, 0x75, 0x53, 0x63, 0x86, 0xd2, 0x72, 0x26, 0xaa, 0x34, 0xe4,
	0x4d, 0x48, 0xcf, 0xc0, 0xf2, 0x94, 0x77, 0x99, 0x36, 0xe8, 0x5b, 0xde, 0x86, 0xf4, 0x02, 0xac,
	0x5f, 0x45, 0xbb, 0x25, 0xfb, 0xaa, 0xad, 0x55, 0x3e, 0x55, 0x79, 0xc1, 0xbb, 0x7f, 0x3c, 0x5b,
	0xa3, 0x43, 0xae, 0x31, 0x23, 0xef, 0x43, 0x7a, 0x0e, 0xd6, 0xf6, 0x0c, 0x6a, 0xc9, 0x72, 0xec,
	0xe8, 0x2e, 0x33, 0xe6, 0x48, 0xe9, 0xea, 0x6f, 0xfd, 0x50, 0x46, 0xec, 0xa1, 0x3b, 0x19, 0xca,
	0xcc, 0x5d, 0xca, 0x17, 0x3e, 0x86, 0x74, 0x1d, 0x56, 0x6e, 0xa2, 0xe6, 0x7d, 0x9e, 0x32, 0xcb,
	0x95, 0x9c, 0xd6, 0x3e, 0x95, 0xa3, 0x6a, 0xcb, 0x6b, 0xcc, 0xb4, 0x46, 0xdc, 0x58, 0x43, 0x3e,
	0x87, 0xee, 0x70, 0xdd, 0x81, 0x92, 0x33, 0xf0, 0x4b, 0x09, 0x5b, 0x39, 0xe3, 0x62, 0x0a, 0xbf,
	0x96, 0x63, 0x77, 0xd0, 0x1e, 0x29, 0x7d, 0xd0, 0x2c, 0x9f, 0x8b, 0xb7, 0x7e, 0xd2, 0xa0, 0xcb,
	0xb0, 0x58, 0x15, 0x36, 0x91, 0x65, 0x82, 0x4b, 0x24, 0x4f, 0x1b, 0xf4, 0x3c, 0x9c, 0x9e, 0xc8,
	0xa5, 0xc4, 0xd4, 0x25, 0xea, 0x61, 0xbf, 0x30, 0x98, 0x91, 0x67, 0x0d, 0x77, 0xbe, 0x49, 0x93,
	0xb1, 0x7b, 0x52, 0x23, 0x4b, 0x07, 0x2c, 0x11, 0x48, 0x9e, 0x37, 0xdc, 0xfc, 0x68, 0x28, 0x66,
	0xc7, 0xdc, 0x06, 0xd7, 0xe1, 0x20, 0x6a, 0x3b, 0xd9, 0xb0, 0xcd, 0xb8, 0xc0, 0x8c, 0xdc, 0x01,
	0xb7, 0x5e, 0x34, 0x14, 0x5d, 0xd4, 0xb9, 0xbb, 0x8c, 0x92, 0x9b, 0x28, 0xc7, 0xe4, 0x2e, 0x50,
	0x0a, 0x0b, 0xd1, 0x50, 0xdc, 0x28, 0x50, 0x8f, 0xbd, 0xd1, 0x3d, 0xa0, 0x4b, 0x70, 0x22, 0x1a,
	0x8a, 0x28, 0x65, 0x95, 0xf7, 0xfd, 0x1a, 0xb5, 0x46, 0x98, 0x7a, 0xf4, 0x00, 0xdc, 0x56, 0xd1,
	0x50, 0xec, 0x8e, 0x36, 0x70, 0x9f, 0x57, 0xc2, 0x87, 0xf5, 0x9c, 0xdd, 0x51, 0x4f, 0xa9, 0x84,
	0xa5, 0x07, 0x9e, 0x3f, 0x02, 0xf7, 0x9a, 0x4a, 0xee, 0x32, 0x73, 0xeb, 0xf1, 0x63, 0x48, 0xe6,
	0xca, 0x0f, 0xec, 0xf2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0x1b, 0xcf, 0xc8, 0x74, 0x03,
	0x00, 0x00,
}
