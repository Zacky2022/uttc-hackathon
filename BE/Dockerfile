FROM golang:1.18 as builder
COPY ./controller ./
COPY ./dao ./
COPY ./model ./
COPY ./usecase ./
COPY ./main.go ./
ENV MYSQL_USER=MYSQL_USER, MYSQL_PWD=MYSQL_PWD, MYSQL_HOST=MYSQL_HOST, MYSQL_DATABASE=MYSQL_DATABASE
CMD ["go","run","main.go"]
