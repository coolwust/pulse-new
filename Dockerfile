FROM docker.io/fedora:latest

MAINTAINER Xianghan Wang <coolwust@gmail.com>

ENV GO_VERSION   1.6.2
ENV GO_URL       https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz
ENV GOROOT       /opt/go
ENV GOPATH       /go
ENV NODE_VERSION v6.1.0
ENV NODE_URL     https://nodejs.org/dist/${NODE_VERSION}/node-${NODE_VERSION}-linux-x64.tar.xz
ENV NODE_ROOT    /opt/node
ENV PATH         $GOROOT/bin:$NODE_ROOT/bin:$PATH
ENV WORKDIR      /go/src/github.com/coldume/pulse

RUN set -ex \
		&& dnf update -y \
		&& dnf install -y \
			git \
			tar \
			xz \
			curl \
			make \
			gem \
		&& gem install \
			sass \
		&& cd /opt \
		&& curl $GO_URL | tar -zxf - \
		&& curl $NODE_URL | tar -Jxf - \
		&& mv node* node \
		&& node/bin/npm install -g \
			typescript \
			pug-cli \
			typings \
			uglify-js \
		&& mkdir -p $WORKDIR

WORKDIR $WORKDIR

EXPOSE 80
