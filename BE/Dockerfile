FROM golang:1.18 as builder
COPY ["./controller/", "./dao/", "./model/", "./usecase/", "./main.go/", "home/"]
CMD ["go","run","main.go"]
