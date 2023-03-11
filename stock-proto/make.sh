#!/usr/bin/env bash

protos_dir="./protos"
domin="github.com"
org="RavenHuo"
# 这里是路径
project_name="go-proto/stock-proto"
pb_package="pkg"

## batch clean old code
# gsed -i "s/^\(package\) \(.*;\)/\1 nonoapi;/" protos/**/*.proto
sed -i "/option go_package.*/d" protos/**/*.proto
files=$(ls -l ${protos_dir} |awk '/^d/ {print $NF}')


echo "==============================================="
echo "Compile"
echo "==============================================="
for dirname in $files
do
  echo "change compile package option => ${domin}/${org}/${project_name}/${pb_package}/${dirname}"
  res=$(sed -i "s?^\(package .*;\)?\1\\noption go_package = \"${domin}\\/${org}\\/${project_name}\\/${pb_package}\\/${dirname}\";?" ${protos_dir}/${dirname}/*.proto)
done