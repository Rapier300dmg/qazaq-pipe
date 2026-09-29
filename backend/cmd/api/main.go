package main

import (
	"log"
	"net/http"
	"pipe-company/backend/internal/database"
	"pipe-company/backend/internal/handlers"
	"time"
)

// enableCORS разрешает браузерам с любого домена получать данные с сервера
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Подключение к БД PostgreSQL (укажите ваш пароль)
	database.InitDB("postgres", "123", "qazaq_pipe")

	mux := http.NewServeMux()

	// Маршруты API
	mux.HandleFunc("/api/v1/products", handlers.GetProductsHandler)
	mux.HandleFunc("/api/v1/products/", handlers.GetProductsHandler)
	mux.HandleFunc("/api/v1/contacts", handlers.GetContactsHandler)
	mux.HandleFunc("/api/v1/contacts/", handlers.GetContactsHandler)

	// Главный тестовый эндпоинт
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Бэкенд QAZAQ PIPE работает в режиме каталога!"))
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      enableCORS(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Сервер запущен на http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}