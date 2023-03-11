#!/usr/bin/env bash
cd $(dirname $0)

PROTOC=protoc
if (which protoc.exe) then
    PROTOC=protoc.exe
fi

version=`$PROTOC --version`

for proto_file in protos/*/*.proto; do
    ## 替换路径
    pb_file=${proto_file/protos/pkg}
    pb_file=${pb_file/.proto/.pb.go}
    echo $proto_file
    echo $pb_file
    if [ ! -e $pb_file ] || [ $proto_file -nt $pb_file ]; then

        echo "Compiling $proto_file -> $pb_file ..."
        $PROTOC $proto_file --plugin=protoc-gen-go=/usr/local/bin/protoc-gen-go --go_out=plugins=grpc:.
    fi
done
