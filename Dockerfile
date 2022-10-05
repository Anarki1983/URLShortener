#基礎映像檔資訊
FROM golang:1.19 as builder
#維護者資訊
MAINTAINER liam anarki1983@gmail.com

#映像檔操作指令
RUN  mkdir -p app
#移動工作目錄到/app
WORKDIR  /app
#將本地檔案移到app目錄內
COPY . .

#編譯 server 執行檔並放在目前工作目錄內
RUN go mod tidy
RUN go build -o webserver ./webservice.go

FROM golang:1.19 as runner
WORKDIR /app
COPY --from=builder /app/webserver .

USER root
ENV TZ=Asia/Taipei
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT ["./webserver"]