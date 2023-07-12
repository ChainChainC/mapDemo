FROM golang:1.18-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY .. ./
RUN go build -o /mapDemo
EXPOSE 30022
CMD [ "/mapDemo" ]

# docker build --tag map-demo:latest .