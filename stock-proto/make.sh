#!/usr/bin/env bash
cd $(dirname $0)

protoDir="protos"
pb_package="pkg"

PROTOC=protoc
if (which protoc.exe) then
    PROTOC=protoc.exe
fi


if [ -e ${pb_package} ]
then
echo "${pb_package} 已存在"
else
mkdir ${pb_package}
echo "${pb_package}目录不存在，创建文件夹"
fi


version=`$PROTOC --version`

cd ${protoDir}

for proto_file in */*.proto; do
    ## 替换路径
    pb_file=${proto_file/protos/pkg}
    pb_file=${pb_file/.proto/.pb.go}
    echo $proto_file
    echo $pb_file
    if [ ! -e $pb_file ] || [ $proto_file -nt $pb_file ]; then

        echo "Compiling $proto_file -> $pb_file ..."
        # 输出到 ../${pb_package}
        # 一般都会使用source_relative,否则，就会按照option go_package里面的那个路径生成
        # -I 表示包含全部
        $PROTOC  -I $proto_file --plugin=protoc-gen-go=../../bin/protoc-gen-go.exe --go_out=../${pb_package} --go_opt=paths=source_relative
    fi
done
