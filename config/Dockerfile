FROM golang:1.23
MAINTAINER Eduard Camacho<edu.kmxo@gmail.com>

RUN apt update && apt -y install lsof

RUN apt install -y tzdata

RUN useradd -m developer --uid=1000 && echo "developer:docker" | chpasswd

WORKDIR /go/src

