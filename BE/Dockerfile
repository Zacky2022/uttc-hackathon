FROM golang:1.18 as builder
COPY ./controller ./
COPY ./dao ./
COPY ./model ./
COPY ./usecase ./
COPY ./main.go ./
COPY ./go.mod ./
COPY ./go.sum ./
CMD ["go","run","main.go"]
