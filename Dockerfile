FROM golang:1.18.1 AS development

WORKDIR /app
COPY go.mod go.sum ./
COPY main.go ./

RUN go build -o /server

EXPOSE 8000

CMD ["/server"]