FROM golang:1.18 as builder
COPY ./controller ./
COPY ./dao ./
COPY ./model ./
COPY ./usecase ./
COPY ./main.go ./
ENV SSLCA = /
CMD ["go","run","main.go"]
