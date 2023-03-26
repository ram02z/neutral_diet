FROM golang:1.20.2-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build cmd/app/main.go

EXPOSE $PORT

CMD [ "./main" ]
