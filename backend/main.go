// backend/main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Структура для ответа
type Message struct {
	Message string `json:"message"`
}

// Функция для подключения к MongoDB
func connectToMongo() {
	var err error
	mongoURI := os.Getenv("MONGO_URI") // URI MongoDB
	if mongoURI == "" {
		mongoURI = "mongodb://mongo:27017" // По умолчанию подключение к MongoDB контейнеру (если используется Docker)
	}

	// Подключение к MongoDB
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	// Проверка подключения
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB!")
}

// Обработчик для запроса
func getMessage(w http.ResponseWriter, r *http.Request) {
    collection := client.Database("mydatabase").Collection("messages")

    var result Message
    err := collection.FindOne(context.TODO(), nil).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            result = Message{Message: "No messages found."} // Возвращаем дефолтное сообщение
        } else {
            http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
            return
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}


func main() {
	// Подключаемся к MongoDB
	connectToMongo()

	// Настройка CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Разрешенные источники
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Обработчик маршрутов
	mux := http.NewServeMux()
	mux.HandleFunc("/api", getMessage)

	// Применяем CORS к маршрутам
	handler := c.Handler(mux)

	// Запуск сервера на порту 8080
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
