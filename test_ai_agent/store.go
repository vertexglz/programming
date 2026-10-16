package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// Task представляет одну задачу в трекере.
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// TaskStore хранит задачи в памяти и синхронизирует доступ с помощью mutex.
// Данные сохраняются в файл tasks.json после каждого изменения.
type TaskStore struct {
	mu       sync.Mutex
	tasks    []Task
	nextID   int
	filePath string
}

// NewTaskStore создаёт новое хранилище.
// Считывает задачи из указанного файла, если он существует.
// Если файл отсутствует или повреждён, начинает с пустого списка.
func NewTaskStore(filePath string) *TaskStore {
	store := &TaskStore{
		filePath: filePath,
		nextID:   1,
	}
	store.load()
	return store
}

// load загружает задачи из JSON-файла при старте приложения.
// Если файл не найден, это нормально — хранилище остаётся пустым.
func (s *TaskStore) load() {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		// Файл может не существовать — начинаем с пустого списка.
		return
	}
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		// При повреждении файла логируем предупреждение и начинаем с нуля.
		log.Printf("Предупреждение: не удалось разобрать файл %s: %v", s.filePath, err)
		return
	}
	s.tasks = tasks
	// Определяем следующий ID как максимальный из существующих + 1.
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	s.nextID = maxID + 1
}

// save сохраняет текущие задачи в JSON-файл.
// Вызывается после каждого изменения списка задач.
func (s *TaskStore) save() {
	data, err := json.MarshalIndent(s.tasks, "", "  ")
	if err != nil {
		log.Printf("Ошибка при кодировании задач в JSON: %v", err)
		return
	}
	if err := os.WriteFile(s.filePath, data, 0644); err != nil {
		log.Printf("Ошибка при сохранении файла %s: %v", s.filePath, err)
	}
}

// List возвращает копию всех задач.
// Блокирует доступ на время чтения, чтобы избежать гонок данных.
func (s *TaskStore) List() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Возвращаем копию, чтобы внешний код не мог изменять внутренний срез напрямую.
	result := make([]Task, len(s.tasks))
	copy(result, s.tasks)
	return result
}

// Add добавляет новую задачу с указанным заголовком.
// Присваивает уникальный ID и текущее время создания.
// Пустые заголовки отклоняются.
// После успешного добавления сохраняет изменения в файл.
// Возвращает добавленную задачу или ошибку.
func (s *TaskStore) Add(title string) (Task, error) {
	if title == "" {
		return Task{}, fmt.Errorf("заголовок задачи не может быть пустым")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	task := Task{
		ID:        s.nextID,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	s.tasks = append(s.tasks, task)
	s.nextID++
	// Сохраняем изменения в файл.
	s.save()
	return task, nil
}