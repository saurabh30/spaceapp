FROM golang:1.19

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o api .

EXPOSE 4000


CMD ["./api"]