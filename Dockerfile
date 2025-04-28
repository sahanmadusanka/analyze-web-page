FROM golang:1.24-bookworm

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app

EXPOSE 3000

CMD ["/build/app"]