package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates/*
var embedFS embed.FS

// Глобальный шаблон страницы, встраиваемый в бинарник через пакет embed.
// Это делает приложение независимым от рабочей папки при запуске.
var tmpl = template.Must(
	template.ParseFS(embedFS, "templates/index.html"),
)

// handleIndex возвращает обработчик для GET-запроса на главной странице.
// Отображает список задач и форму добавления.
func handleIndex(store *TaskStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks := store.List()
		data := struct {
			Tasks []Task
		}{
			Tasks: tasks,
		}
		// Устанавливаем тип контента HTML.
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		// Рендерим шаблон. Пользовательский ввод в шаблоне экранируется автоматически.
		if err := tmpl.Execute(w, data); err != nil {
			// Ошибка рендеринга шаблона — логируем и возвращаем 500.
			log.Printf("Ошибка при рендеринге шаблона: %v", err)
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		}
	}
}

// handleAdd обрабатывает POST-запрос на добавление новой задачи.
// После успешного добавления перенаправляет пользователя на главную страницу (303 See Other).
func handleAdd(store *TaskStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Обрабатываем только POST-запросы; остальные методы отклоняем.
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
			return
		}
		// Парсим форму с одним полем "title".
		if err := r.ParseForm(); err != nil {
			log.Printf("Ошибка при парсинге формы: %v", err)
			http.Error(w, "Неверный запрос", http.StatusBadRequest)
			return
		}
		title := r.FormValue("title")
		// Добавляем задачу в хранилище.
		_, err := store.Add(title)
		if err != nil {
			// Ошибка добавления (например, пустое название) — возвращаем 400.
			log.Printf("Ошибка при добавлении задачи: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Успешное добавление — редирект на главную страницу.
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusSeeOther)
	}
}

// handleHealthz отвечает на проверку живости сервера.
// Используется хостингом (например, Render) для мониторинга доступности.
func handleHealthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}