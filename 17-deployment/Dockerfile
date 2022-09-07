FROM golang:1.18

# membuat direktori app
RUN mkdir /app

# set working directory /app
WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o be11-api

CMD ["./be11-api"]