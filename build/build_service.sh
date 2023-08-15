#!/bin/sh
echo create folder for build package ...
if [ ! -e package ]; then
  mkdir package
fi

echo build the code ...
cd ../service
#添加参数CGO_ENABLED=0，关闭CGO,是为了是编译后的程序可以在alpine中运行
CGO_ENABLED=0 go build
cd ../build

echo remove last package if exist
if [ ! -e package/service ]; then
  mkdir package/service
fi

if [ -e package/service/gwprjservice ]; then
  rm -rf package/service/gwprjservice
fi

mv ../service/gwprj ./package/service/gwprjservice

echo gwprjservice build over.