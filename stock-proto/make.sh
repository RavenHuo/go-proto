#!/usr/bin/env bash
cd $(dirname $0)

protoDir="protos"
pb_package="pkg"

PROTOC=protoc
if (which protoc.exe) then
    PROTOC=protoc.exe
fi


if [ -d ${pb_package} ]
then
echo "${pb_package} 已存在"
else
mkdir ${pb_package}
echo "${pb_package}目录不存在，创建文件夹"
fi


version=`$PROTOC --version`

cd ${protoDir}
echo "cd ${protoDir}"

for dirName in */;do
  echo "${dirName}"
  for proto_file in ${dirName}*.proto; do
      ## 替换路径

        if [ ! -d ../${pb_package}/${dirName} ];then
          mkdir ../${pb_package}/${dirName}
          echo "../${pb_package}/${dirName}目录不存在，创建文件夹"
        fi

        echo "Compiling $proto_file -> $pb_file ..."

        # --proto_path=${dirName}：指定了 .proto 文件所在的路径，${dirName} 是一个变量，它表示当前目录的名称，即 .proto 文件所在的目录。
        #--proto_path=.：指定了搜索 .proto 文件的路径，即当前目录。
        #--go_out=plugins=grpc:../${pb_package}/${dirName}：指定了生成 Go 语言代码的输出路径。其中 plugins=grpc 表示同时生成 gRPC 代码，../${pb_package}/${dirName} 表示输出路径，${pb_package} 是一个变量，表示该 .proto 文件的包名。
        #--go_opt=paths=source_relative：指定了生成的 Go 代码中的 import 路径是相对路径。
        #$proto_file：指定了要编译的 .proto 文件的名称。
        #--plugin=protoc-gen-go=../../bin/protoc-gen-go.exe：指定了使用的插件，protoc-gen-go 是一个生成 Go 语言代码的插件，../../bin/protoc-gen-go.exe 是插件的路径。
#        $PROTOC --proto_path=${dirName} --proto_path=. --go_out=plugins=grpc:../${pb_package}/${dirName}  --go-grpc_out=. --go-grpc_opt=paths=source_relative $proto_file
        $PROTOC --proto_path=${dirName} --proto_path=. --go_out=plugins=grpc:../${pb_package}/${dirName}  --go_opt=paths=source_relative $proto_file
#        $PROTOC --proto_path=${dirName} --proto_path=. --go_out=plugins=grpc:../${pb_package}/${dirName}  --go_opt=paths=source_relative $proto_file --plugin=protoc-gen-go=../../bin/protoc-gen-go.exe

  done
done