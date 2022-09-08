FROM golang:1.18.6

WORKDIR /ebiznes-consoleshop-backend

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

RUN go build -o /consoleshop-backend
EXPOSE 8000

CMD ["/consoleshop-backend"]
