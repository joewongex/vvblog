# alpine不支持go二进制文件 
# slim没有wget
FROM node:17.5.0
ADD go1.17.7.linux-amd64.tar.gz /usr/local/
ENV PATH=/usr/local/go/bin:/root/go/bin:$PATH
RUN npm config set -g registry http://mirrors.cloud.tencent.com/npm/ && \  
  go env -w GO111MODULE=on && \
  go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct && \
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin  