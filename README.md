# protos
proto files

# 说明
- 1、请使用项目根目录下的compile_go.sh脚本编译此项目的go语言版本。
- 2、关于proto文件引用  
(1) 引用自己当前目录外其他目录的proto文件时  
import的路径需要从此项目的根目录下一级写起，示例：import "common/error_code.proto";  
(2) 引用自己当前目录下的proto文件时  
import时直接写此问价，示例：import "client.proto";  
(3) 引用自己当前目录下其他目录的proto文件时  
import的路径需要从此项目的根目录下一级写起，示例：import "common/error_code.proto";  
- 3、其他项目引用此protos项目  
protos项目可以放置在任意目录，然后在引用此protos项目的mod文件中，增加一条本地包引用即可，示例：replace 包名 => /home/nssy/Work/4-zjmnssy/protos
- 4、添加新的proto文件时，需要将新的proto文件的编译命令添加到对应的编译文件中，注意，各语言版本的编译脚本都要添加。  


