FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8000

RUN go get github.com/gorilla/mux;
RUN go get -u github.com/go-playground/validator/v10;
RUN go get -u github.com/google/wire;
RUN go get -u github.com/hashicorp/go-memdb

RUN go build

CMD ["./payloadrest"]