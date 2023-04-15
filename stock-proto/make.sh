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
      pb_file=${proto_file/.proto/.pb.go}

      if [ ! -e $pb_file ] || [ $proto_file -nt $pb_file ]; then
          if [ ! -d ../${pb_package}/${dirName} ];then
            mkdir ../${pb_package}/${dirName}
            echo "../${pb_package}/${dirName}目录不存在，创建文件夹"
          fi

          echo "Compiling $proto_file -> $pb_file ..."

          # 输出到 ../${pb_package}
          # 一般都会使用source_relative,否则，就会按照option go_package里面的那个路径生成
          #  --proto_path 表示引用外部的proto
          $PROTOC --proto_path=${dirName} --proto_path=. --go_out=../${pb_package}/${dirName} --go_opt=paths=source_relative $proto_file --plugin=protoc-gen-go=../../bin/protoc-gen-go.exe
      fi
  done
done