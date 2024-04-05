FROM golang:1.19.1
WORKDIR /go/src
ENV GOPROXY=https://goproxy.cn,direct
COPY . .
ENV GOOS=linux
RUN go build -o dispatcher ./main.go
EXPOSE 16444
#启动容器时候指定运行的程序:可执行文件
#RUN chown -R 777 
CMD ["./dispatcher"]