# protos
proto files

# 说明
- 1、请使用项目根目录下的compile_go.sh脚本编译此项目的go语言版本。
- 2、关于proto文件引用  
(1) 引用自己当前目录外其他目录的proto文件时  
(2) 引用自己当前目录下的proto文件时与(1)一样  
(3) 引用自己当前目录下其他目录的proto文件时  
import的路径都需要从此项目的proto目录的下一级写起，示例：import "common/error_code.proto";  


