FROM golang:1.19-alpine

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o main .

ENV PORT=2424

EXPOSE 2424

CMD ["./main"]