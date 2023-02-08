FROM golang:1.19-alpine3.16

WORKDIR /usr/src/${PROJECT_NAME}

COPY . .

RUN go mod download
RUN go build -o transactions main.go

CMD ["./transactions"]
