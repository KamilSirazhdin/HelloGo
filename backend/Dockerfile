# backend/Dockerfile
FROM golang:1.24-alpine

# Установим необходимые зависимости
WORKDIR /go/src/app

# Копируем исходный код
COPY . .

# Загружаем зависимости
RUN go mod tidy

# Скомпилируем приложение
RUN go build -o backend .

# Устанавливаем переменную окружения для MongoDB
ENV MONGO_URI mongodb://mongo:27017

# Запускаем приложение
CMD ["./backend"]
