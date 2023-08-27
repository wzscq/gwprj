#!/bin/sh
echo create folder for build package ...
if [ ! -e package ]; then
  mkdir package
fi

if [ ! -e package/web ]; then
  mkdir package/web
fi

echo build the code ...
cd ../function
npm install
npm run build
cd ../build

echo remove last package if exist
if [ -e package/web/function ]; then
  rm -rf package/web/function
fi

mv ../function/build ./package/web/function

echo functionpage package build over.
