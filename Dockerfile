FROM golang:1.19.3-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go env -w GO111MODULE=auto
RUN go build -o main .
CMD ["/app/main"]