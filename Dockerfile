FROM golang:alpine
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY ./ ./

# Соберём приложение 
RUN go build ./cmd/webserver
EXPOSE 8090
# Запустим приложение
CMD ["./webserver"]

#команды для сборки и запуска 
#docker build -t testgo1 .
#docker run -p 8090:8090 --name testgo1 testgo1