package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Точка входа в приложение.
// Запускает HTTP-сервер с поддержкой graceful shutdown и настройками через переменные окружения.
func main() {
	// Определяем порт из переменной окружения PORT, по умолчанию 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Определяем путь к файлу с задачами из TASKS_FILE, по умолчанию "tasks.json".
	tasksFile := os.Getenv("TASKS_FILE")
	if tasksFile == "" {
		tasksFile = "tasks.json"
	}

	// Создаём хранилище задач. При инициализации оно читает tasks.json,
	// а если файл отсутствует, начинает с пустого списка.
	store := NewTaskStore(tasksFile)

	// Регистрируем обработчики маршрутов.
	// Используем стандартный http.NewServeMux, который в Go 1.22+
	// поддерживает шаблоны вида "GET /path" и method-кондиции.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", handleIndex(store))
	mux.HandleFunc("POST /add", handleAdd(store))
	mux.HandleFunc("GET /healthz", handleHealthz())

	// Создаём HTTP-сервер с таймаутами для предотвращения утечек ресурсов.
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Запускаем сервер в отдельной горoutine, чтобы не блокировать обработку сигналов.
	go func() {
		log.Printf("Сервер запускается на http://localhost:%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Если сервер не смог запуститься (например, порт занят), завершаем программу с ошибкой.
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	// Ожидаем сигнал завершения (SIGINT, SIGTERM) от ОС или хостинга.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("Получен сигнал завершения, останавливаем сервер...")

	// Выполняем graceful shutdown с таймаутом 5 секунд.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		// Если graceful shutdown не успел завершиться за 5 сек, принудительно закрываем.
		log.Printf("Ошибка при плавном завершении: %v", err)
	} else {
		log.Println("Сервер успешно остановлен.")
	}
}