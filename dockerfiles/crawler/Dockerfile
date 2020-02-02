FROM centos:7
MAINTAINER Takayuki Kaneko <taka.knk.11011@gmail.com>

RUN yum install -y sudo which file wget curl make swig xz patch openssl awk git

# GO言語のインストール
RUN yum install -y epel-release
RUN yum install -y golang

RUN mkdir -p /var/www/go/bin/
RUN mkdir -p /var/www/go/pkg/
RUN mkdir -p /var/www/go/src/

ENV GOPATH /var/www/go
ENV PATH $GOPATH/bin/:$PATH

# MeCab自体と形態素解析用の辞書をインストール
RUN yum install -y https://packages.groonga.org/centos/groonga-release-latest.noarch.rpm
RUN yum install -y mecab mecab-ipadic mecab-devel

# MeCabの為の拡張辞書のインストール (https://github.com/neologd/mecab-ipadic-neologd)
# 辞書が /usr/lib64/mecab/dic/mecab-ipadic-neologd にインストールされる
RUN git clone --depth 1 https://github.com/neologd/mecab-ipadic-neologd.git
RUN /mecab-ipadic-neologd/bin/install-mecab-ipadic-neologd -n -y

# Goからdotenvを使う為のライブラリ
RUN go get github.com/joho/godotenv
# ORM
RUN go get github.com/jinzhu/gorm
RUN go get github.com/t-tiger/gorm-bulk-insert
# DI Containerツール
RUN go get go.uber.org/dig
# MySQLドライバ
RUN go get github.com/go-sql-driver/mysql
# スクレイピングライブラリ
RUN go get github.com/PuerkitoBio/goquery
# GO言語からMeCabを使う為のライブラリ
ENV CGO_LDFLAGS "-L/usr/lib64 -lmecab"
ENV CGO_CFLAGS "-I/usr/include"
RUN go get github.com/bluele/mecab-golang

WORKDIR /var/www/go/src/search_engine_project/crawler/
CMD ["/sbin/init"]

EXPOSE 3306
