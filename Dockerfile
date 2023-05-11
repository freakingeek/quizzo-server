FROM golang:latest

ENV PORT 8080

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /quizzo-server

EXPOSE 8080

ENTRYPOINT [ "/quizzo-server" ]