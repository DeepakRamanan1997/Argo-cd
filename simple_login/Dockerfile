FROM golang:alpine
WORKDIR /app
COPY go.mod main.go index.html /app/
EXPOSE 8080
ENTRYPOINT ["go", "run", "main.go"]
