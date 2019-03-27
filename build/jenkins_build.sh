#!/bin/bash

# 1: Declare to project root
echo "Step 1: Declare ENV ..."
export PROJECT_NAME=HelloTencent
export APOLLO_ENV=$HOME/apollo/env
export PROJECT_PATH=$APOLLO_ENV/$PROJECT_NAME

echo "exporting environment valuables..."
echo "PROJECT_NAME = $PROJECT_NAME"
echo "APOLLO_ENV = $APOLLO_ENV"
echo "PROJECT_PATH = $PROJECT_PATH"

echo ""
echo "---------------------------------------"
echo ""

# 2: Move to project source root folder
echo "Step 2: Move to project source root folder..."
cd $PROJECT_NAME

echo ""
echo "---------------------------------------"
echo ""

# 3: Add dependencies
echo "Step 3.1: get go-bindata"
go get -u github.com/jteeuwen/go-bindata/...

echo "Step 3.2: get go-bindata-assetfs"
go get github.com/elazarl/go-bindata-assetfs/...

echo "Step 3.3: get protoc-gen-go"
go get -u github.com/golang/protobuf/protoc-gen-go

echo "Step 3.4: get protoc-gen-grpc-gateway"
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

echo "Step 3.5: get protoc-gen-swagger"
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

echo "Step 3.6: Add GOPATH to PATH"
export PATH=$GOPATH/bin:$PATH

echo ""
echo "---------------------------------------"
echo ""

# 4: Compile proto file
echo "Step 4.1: Compile proto file at api/*.proto ..."
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  api/*.proto

echo "Step 4.2: Compile reverse proxy file at api/*.proto ..."
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  api/*.proto

echo "Step 4.3: Compile swagger file at api/*.proto ..."
protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --swagger_out=logtostderr=true:.   api/*.proto

echo "Step 4.4: Compile swagger to go file at third-party/swagger-ui/* ..."
go-bindata --nocompress -pkg swagger -o examples/ui/data/swagger/datafile.go third-party/swagger-ui/...

echo ""
echo "---------------------------------------"
echo ""

# 5: Build main file
echo "Step 5.1: Build main file at internal/main.go ..."
go build -o HelloTencent internal/main.go

echo "Step 5.2: Build common service client file at examples/common_service_client.go ..."
go build -o CommonClient examples/common_service_client.go

echo "Step 5.3: Build common service client file at examples/common_service_web.go ..."
go build -o CommonWeb examples/common_service_web.go

echo ""
echo "---------------------------------------"
echo ""

# 6: Make target package
echo "Step 6: Make target package ..."
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

# 7: Copy related files to target folder
echo "Step 7: Copy related files to target folder..."
mv $PROJECT_NAME target/$PROJECT_NAME/bin/$PROJECT_NAME
mv CommonClient target/$PROJECT_NAME/bin/CommonClient
mv CommonWeb target/$PROJECT_NAME/bin/CommonWeb
cp configs/* target/$PROJECT_NAME/configs/
cp api/* target/$PROJECT_NAME/api/

echo ""
echo "---------------------------------------"
echo ""