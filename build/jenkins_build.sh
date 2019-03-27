#!/bin/bash

########## 1: Export GoPath
echo "Step 1: Exporting GOPATH ..."
export PATH=/var/jenkins_home/go/bin:$PATH

echo ""
echo "---------------------------------------"
echo ""

########## 2: Declare to ENV 
echo "Step 2: Exporting ENV ..."
export PROJECT_DIR=/var/jenkins_home/tools/org.jenkinsci.plugins.golang.GolangInstallation/go_1.12/src
export PROJECT_NAME=HelloTencent
export PROJECT_PATH=$PROJECT_DIR/$PROJECT_NAME

echo "PROJECT_DIR = $PROJECT_DIR"
echo "PROJECT_NAME = $PROJECT_NAME"
echo "PROJECT_PATH = $PROJECT_PATH"

echo ""
echo "---------------------------------------"
echo ""

########## 3: Move to PROJECT_DIR
echo "Step 3: Move to project directory ..."
cd $PROJECT_DIR

echo ""
echo "---------------------------------------"
echo ""

########## 4: Clear PROJECT_PATH
echo "Step 4: Clear project path ..."
rm -rf $PROJECT_PATH

echo ""
echo "---------------------------------------"
echo ""

########## 5: Git Clone
echo "Step 5: Git Cloning ..."
git clone https://github.com/donghun221/HelloTencent.git

echo ""
echo "---------------------------------------"
echo ""

########## 6: Compile proto file
echo "Step 6.1: Compile proto file at api/*.proto ..."
/var/jenkins_home/go/bin/protoc -I/usr/local/include -I. \
  -I/var/jenkins_home/go/src \
  -I/var/jenkins_home/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  api/*.proto

echo "Step 6.2: Compile reverse proxy file at api/*.proto ..."
/var/jenkins_home/go/bin/protoc -I/usr/local/include -I. \
  -I/var/jenkins_home/go/src \
  -I/var/jenkins_home/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  api/*.proto

echo "Step 6.3: Compile swagger file at api/*.proto ..."
/var/jenkins_home/go/bin/protoc -I/usr/local/include -I.   -I/var/jenkins_home/go/src   -I/var/jenkins_home/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --swagger_out=logtostderr=true:.   api/*.proto

echo "Step 6.4: Compile swagger to go file at third-party/swagger-ui/* ..."
/var/jenkins_home/go/bin/go-bindata --nocompress -pkg swagger -o examples/ui/data/swagger/datafile.go third-party/swagger-ui/...

echo ""
echo "---------------------------------------"
echo ""


########## 7: Build main file
echo "Step 7.1: Build main file at internal/main.go ..."
go build -o HelloTencent internal/main.go

echo "Step 7.2: Build common service client file at examples/common_service_client.go ..."
go build -o CommonClient examples/common_service_client.go

echo "Step 7.3: Build common service client file at examples/common_service_web.go ..."
go build -o CommonWeb examples/common_service_web.go

echo ""
echo "---------------------------------------"
echo ""

# 8: Make target package
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
