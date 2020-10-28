
FROM golang:1.15.3-alpine3.12
WORKDIR "/app"
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o main main.go

CMD ["./main"]

#docker run -it --rm --name gotest gotest
