FROM golang:1.19-alpine

WORKDIR /usr/src/app

COPY go.* ./

RUN go mod download && go mod verify

COPY . . 

RUN go build -v -o main /usr/local/bin/app ./...

EXPOSE 8000

CMD ["app"]
