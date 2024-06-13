FROM golang:1.22.4
WORKDIR /app
RUN go install github.com/air-verse/air@v1.52.2
CMD ["air"]