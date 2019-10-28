# protos
proto files

# 说明
- 1、使用compile_go.sh脚本编译此项目。
- 2、proto文件引用时，import的路径需要从此项目的根目录写起，示例：import "protos/private/common/error_code.proto";。
- 3、其他项目引用此protos项目时，protos项目可以放置在任意目录，然后在引用此protos项目的mod文件中，增加一条本地包引用即可，示例：replace protos => /home/nssy/Work/4-zjmnssy/protos
