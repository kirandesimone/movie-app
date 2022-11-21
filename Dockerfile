FROM golang:1.19-alpine

RUN mkdir /app

WORKDIR /app

COPY go.* ./

RUN go mod download && go mod verify

COPY . . 

RUN go build -v -o main .

EXPOSE 8000

CMD ["app/main"]
