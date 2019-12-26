FROM centos:7
MAINTAINER Takayuki Kaneko <taka.knk.11011@gmail.com>

RUN yum install -y sudo which

# GO言語のインストール
RUN yum install -y epel-release
RUN yum install -y golang

RUN mkdir -p /var/www/go/bin/
RUN mkdir -p /var/www/go/pkg/
RUN mkdir -p /var/www/go/src/

ENV GOPATH /var/www/go
ENV PATH $GOPATH/bin/:$PATH

# Goからdotenvを使う為のライブラリ
RUN go get github.com/joho/godotenv
# MySQLドライバ
RUN go get github.com/go-sql-driver/mysql
# スクレイピングライブラリ
RUN go get github.com/PuerkitoBio/goquery

WORKDIR /var/www/go/src/search_engine_project/crawler/
CMD ["/sbin/init"]

EXPOSE 3306
