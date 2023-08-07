FROM golang:1.20-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o main .
RUN chmod +x ./wait.sh

EXPOSE 8888

CMD ./wait.sh database:5432 -- ./main