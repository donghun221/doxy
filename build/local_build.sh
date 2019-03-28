#!/bin/bash

# 0: check env
if [ -z $1 ];then
	echo "usage:mac|linux|windows"
	exit -1
fi

if [ $1 == "windows"];then
	GS="windows"
elif [ $1 == "mac" ];then
	GS="darwin"
else
	GS="linux"
fi

echo "machine env: "$GS
# 1: Declare to project root
echo "Step 1: Declare ENV ..."
export PROJECT_NAME=HelloTencent
export APOLLO_ENV=$HOME/apollo/env
export PROJECT_PATH=$APOLLO_ENV/$PROJECT_NAME

echo "exporting environment valuables..."
echo "PROJECT_NAME = $PROJECT_NAME"
echo "APOLLO_ENV = $APOLLO_ENV"
echo "PROJECT_PATH = $PROJECT_PATH"

# 2: Move to project source root folder
echo "Step 2: Move to project source root folder..."
cp -r ../../$PROJECT_NAME $GOPATH/src/$PROJECT_NAME
cd $GOPATH/src/$PROJECT_NAME


# 3: Compile proto file
echo "Step 3.1: Compile proto file at api/*.proto ..."
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  api/*.proto

echo "Step 3.2: Compile reverse proxy file at api/*.proto ..."
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  api/*.proto

echo "Step 3.3: Compile swagger file at api/*.proto ..."
protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --swagger_out=logtostderr=true:.   api/*.proto

echo "Step 3.4: Compile swagger to go file at third-party/swagger-ui/* ..."
go-bindata --nocompress -pkg swagger -o examples/ui/data/swagger/datafile.go third-party/swagger-ui/...

# 4: Build main file
echo "Step 4.1: Build main file at internal/main.go ..."
GOOS=$GS GOARCH=amd64 go build -o HelloTencent internal/main.go

echo "Step 4.2: Build common service client file at examples/common_service_client.go ..."
GOOS=$GS GOARCH=amd64 go build -o CommonClient examples/common_service_client.go

echo "Step 4.3: Build common service client file at examples/common_service_web.go ..."
GOOS=$GS GOARCH=amd64 go build -o CommonWeb examples/common_service_web.go

# 5: Make target package
echo "Step 5: Make target package ..."
rm -rf target

mkdir target \
      target/$PROJECT_NAME
mkdir target/$PROJECT_NAME/bin \
      target/$PROJECT_NAME/cmd \
      target/$PROJECT_NAME/tools \
      target/$PROJECT_NAME/env \
      target/$PROJECT_NAME/init \
      target/$PROJECT_NAME/scripts \
      target/$PROJECT_NAME/var \
      target/$PROJECT_NAME/logs \
      target/$PROJECT_NAME/configs \
      target/$PROJECT_NAME/api \
      target/$PROJECT_NAME/assets

# 6: Copy related files to target folder
echo "Step 6: Copy related files to target folder..."
mv $PROJECT_NAME target/$PROJECT_NAME/bin/$PROJECT_NAME
mv CommonClient target/$PROJECT_NAME/bin/CommonClient
mv CommonWeb target/$PROJECT_NAME/bin/CommonWeb
cp configs/* target/$PROJECT_NAME/configs/
cp api/* target/$PROJECT_NAME/api/

# 7: Copy to apollo env
echo "Step 7: Copy to apollo env..."
cp -r target/ $APOLLO_ENV/
