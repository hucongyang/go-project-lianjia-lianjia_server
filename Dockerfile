FROM scratch

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /Users/hucy/Code/GoProject/lianjia/lianjia_server
COPY . /Users/hucy/Code/GoProject/lianjia/lianjia_server

EXPOSE 8080
CMD ["./lianjia_server"]