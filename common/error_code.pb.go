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
	ErrorCode_WebError          ErrorCode = 1000
	ErrorCode_WebMethodError    ErrorCode = 1001
	ErrorCode_WebRouteNotFound  ErrorCode = 1002
	ErrorCode_WebCookieOntFound ErrorCode = 1003
	ErrorCode_WebCookieIllegal  ErrorCode = 1004
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
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0x49, 0x72, 0x13, 0x31,
	0x14, 0x86, 0x61, 0x13, 0xb7, 0x5f, 0x08, 0x51, 0x44, 0x06, 0x08, 0x50, 0xec, 0x59, 0xb0, 0xe1,
	0x06, 0x71, 0x1c, 0x08, 0x55, 0x49, 0x8c, 0x3b, 0xc6, 0x4b, 0x4a, 0xdd, 0xfd, 0x6c, 0xab, 0xac,
	0xd6, 0x8b, 0x25, 0x75, 0xc5, 0xbe, 0x07, 0x73, 0x38, 0x04, 0xc3, 0x25, 0x58, 0x30, 0xcf, 0xdc,
	0x80, 0xe9, 0x10, 0x94, 0x5a, 0x9e, 0x8a, 0x65, 0x7f, 0xef, 0x7f, 0xff, 0x9b, 0x5a, 0xc0, 0xd0,
	0x18, 0x32, 0x77, 0x53, 0xca, 0xf0, 0xda, 0x91, 0x21, 0x47, 0x7c, 0x21, 0xa5, 0x3c, 0x27, 0x7d,
	0xf5, 0x64, 0x01, 0xaa, 0x75, 0x1f, 0xac, 0x51, 0x86, 0x7c, 0x11, 0x2a, 0x71, 0x91, 0xa6, 0x68,
	0x2d, 0x3b, 0xc5, 0x97, 0x01, 0xda, 0x98, 0x4c, 0xbe, 0x5f, 0x9d, 0xe6, 0x4b, 0x10, 0xb5, 0x31,
	0x29, 0xd5, 0xec, 0x67, 0x85, 0x9f, 0x83, 0xb3, 0x6d, 0x4c, 0xf6, 0xd0, 0xf5, 0x28, 0x0b, 0xf0,
	0x57, 0x85, 0xaf, 0x01, 0x6b, 0x63, 0xd2, 0xa4, 0xc2, 0xe1, 0x3e, 0xb9, 0x1d, 0x2a, 0x74, 0xc6,
	0x7e, 0x57, 0xf8, 0x3a, 0xac, 0xb4, 0x31, 0xa9, 0x11, 0xf5, 0x25, 0x1e, 0xe8, 0x31, 0xff, 0x33,
	0x91, 0x07, 0xbe, 0xab, 0x14, 0x76, 0x85, 0x62, 0x7f, 0x2b, 0x9c, 0xc1, 0x62, 0x3c, 0xb2, 0x0e,
	0xf3, 0xe0, 0xfb, 0x3a, 0xf2, 0xcd, 0x04, 0xb2, 0x55, 0xd8, 0x11, 0x7b, 0x13, 0xf9, 0xcc, 0x5b,
	0xf1, 0xc1, 0xfe, 0x9e, 0x30, 0xb6, 0x27, 0x54, 0xd0, 0xbd, 0x8d, 0xf8, 0x06, 0x70, 0x8f, 0x5b,
	0x3a, 0x9f, 0x0f, 0xbc, 0x8b, 0x7c, 0x07, 0x35, 0x83, 0x19, 0x6a, 0x27, 0x85, 0xda, 0x93, 0xd6,
	0x4a, 0xdd, 0x65, 0xef, 0x23, 0x7e, 0x01, 0x56, 0x67, 0xbc, 0x21, 0x8c, 0xc5, 0x90, 0xf2, 0x21,
	0xe2, 0x57, 0x60, 0xf3, 0x06, 0xba, 0x5d, 0xdd, 0xa1, 0x1d, 0x43, 0xf9, 0x4c, 0x15, 0x04, 0x1f,
	0xff, 0xf3, 0xac, 0x0f, 0x8f, 0xa4, 0xc1, 0x8c, 0x7d, 0x8a, 0xf8, 0x25, 0xd8, 0x68, 0x59, 0x34,
	0x5a, 0xe4, 0x78, 0x60, 0x1a, 0xc2, 0xda, 0x63, 0x32, 0xe3, 0x15, 0x7d, 0x2e, 0x5b, 0x6c, 0xa2,
	0x5f, 0x3f, 0xea, 0xcc, 0x6f, 0x3d, 0x04, 0xbe, 0x44, 0x7c, 0x13, 0xd6, 0xee, 0xa0, 0x91, 0x1d,
	0x99, 0x0a, 0x27, 0x49, 0xcf, 0x62, 0x5f, 0xcb, 0x52, 0x13, 0xcb, 0x9b, 0xc2, 0xd6, 0x87, 0xd2,
	0x3a, 0xcb, 0xbe, 0x45, 0xfe, 0x08, 0x8d, 0x1e, 0xe9, 0x39, 0xf8, 0xbd, 0x84, 0xf5, 0x5c, 0x48,
	0x35, 0x83, 0x3f, 0xca, 0xb2, 0xfb, 0xe8, 0x8e, 0xc9, 0xf4, 0x6b, 0xe5, 0xe9, 0x83, 0xf5, 0xb3,
	0x2a, 0x5f, 0x85, 0xe5, 0x71, 0x60, 0x1b, 0x45, 0xa6, 0xa4, 0x46, 0xf6, 0xbc, 0xca, 0x2f, 0xc3,
	0xf9, 0xa9, 0x5c, 0x6b, 0x4c, 0x7d, 0x47, 0x4d, 0xec, 0x14, 0x16, 0x33, 0xf6, 0xa2, 0xca, 0x2f,
	0xc2, 0xfa, 0x34, 0xc9, 0xba, 0x96, 0x36, 0x28, 0xd2, 0x9e, 0x48, 0x14, 0xb2, 0x97, 0x55, 0x5f,
	0x3f, 0x1e, 0xa8, 0xf9, 0x32, 0xf7, 0xc0, 0x67, 0x78, 0x88, 0xc6, 0x4d, 0x27, 0xdc, 0x11, 0x52,
	0x61, 0xc6, 0xee, 0x83, 0x1f, 0x2f, 0x1e, 0xa8, 0x06, 0x9a, 0xdc, 0x5f, 0x86, 0xf4, 0x36, 0xea,
	0x11, 0x7b, 0x00, 0x9c, 0xc3, 0x52, 0x3c, 0x50, 0xb7, 0x0b, 0x34, 0xa3, 0x60, 0xf4, 0x10, 0xf8,
	0x0a, 0x9c, 0x89, 0x07, 0x2a, 0x4e, 0xc5, 0xd8, 0xfb, 0xd1, 0x04, 0xd5, 0x87, 0x98, 0x06, 0xf4,
	0x18, 0xfc, 0x54, 0xf1, 0x40, 0x1d, 0x0e, 0xb7, 0xb0, 0x2b, 0xc7, 0xc2, 0x27, 0x93, 0x3a, 0x87,
	0xc3, 0x26, 0x51, 0x22, 0xd2, 0x7e, 0xe0, 0x27, 0xe0, 0xff, 0xa6, 0x92, 0xfb, 0x9e, 0xa5, 0x0b,
	0xf8, 0x29, 0x24, 0x0b, 0xe5, 0x63, 0xb9, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x77, 0x37, 0x94,
	0x54, 0x40, 0x03, 0x00, 0x00,
}
