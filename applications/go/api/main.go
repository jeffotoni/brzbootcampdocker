package main

import (
	"encoding/json"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// Estrutura do Usuário
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Manter em memória
// para simular o banco
// de forma segura
var (
	users = make(map[int64]User)
	mu    sync.Mutex
)

// Configurando o logger global
var logger *slog.Logger

func init() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	logger = slog.New(handler)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/user", getAllUsers)
	mux.HandleFunc("GET /v1/user/{id}", getUser)
	mux.HandleFunc("POST /v1/user", createUser)
	mux.HandleFunc("PUT /v1/user/{id}", updateUser)
	mux.HandleFunc("DELETE /v1/user/{id}", deleteUser)

	logger.Info("Run Server", "em http://localhost:", 8080)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))
}

// Handlers

// createUser
func createUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("Recebida requisição para /v1/user", "method", r.Method, "path", r.URL.Path)
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Error("Erro ao decodificar JSON", "error", err)
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[user.ID]; exists {
		logger.Error("Usuário já existe", "error", errors.New("usuario já existe"))
		http.Error(w, "Usuário já existe", http.StatusConflict)
		return
	}

	logger.Info("Usuário criado com sucesso")

	users[user.ID] = user
	w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(map[string]string{"message": "Usuário criado com sucesso"})}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Usuário criado com sucesso",
		"user": map[string]interface{}{
			"id":   user.ID,
			"name": user.Name,
		},
	})
}

// getAllUsers
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	logger.Info("Recebida requisição para /v1/user", "method", r.Method, "path", r.URL.Path)
	mu.Lock()
	defer mu.Unlock()

	var allUsers []User
	for _, user := range users {
		allUsers = append(allUsers, user)
	}

	logger.Info("Resposta enviada com sucesso", "count", len(allUsers))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allUsers)
}

// getUser
func getUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("Recebida requisição para /v1/user", "method", r.Method, "path", r.URL.Path)
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || idStr == "" {
		if err != nil {
			logger.Error("ID inválido ou não fornecido", "error", err)
		} else {
			logger.Error("ID inválido ou não fornecido", "error", errors.New("error id invalido"))
		}
		http.Error(w, "ID inválido ou não fornecido", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	user, exists := users[id]
	if !exists {
		logger.Error("Usuário não encontrado", "error", nil)
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	logger.Info("Resposta enviada com sucesso")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// updateUser
func updateUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("Recebida requisição para /v1/user", "method", r.Method, "path", r.URL.Path)
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || idStr == "" {
		if err != nil {
			logger.Error("Erro ao decodificar JSON", "error", err)
		} else {
			logger.Error("Erro ao decodificar JSON", "error", errors.New("error id vazio"))
		}
		http.Error(w, "ID inválido ou não fornecido", http.StatusBadRequest)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Error("Erro ao decodificar JSON", "error", err)
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[id]; !exists {
		logger.Error("Usuário não encontrado", "error", errors.New("usuario nao encontrado"))
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	logger.Info("Usuário atualizado com sucesso", "id", id)

	user.ID = id
	users[id] = user
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário atualizado com sucesso"})
}

// deleteUser
func deleteUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("Recebida requisição para /v1/user", "method", r.Method, "path", r.URL.Path)
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || idStr == "" {
		if err != nil {
			logger.Error("Erro de ID", "error", err)
		} else {
			logger.Error("Erro de ID", "error", errors.New("error id invalido/vazio"))
		}

		http.Error(w, "ID inválido ou não fornecido", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[id]; !exists {
		logger.Error("Usuário não encontrado", "error", errors.New("error id não encontrado"))
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	delete(users, id)
	logger.Info("Usuário deletado com sucesso")
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário deletado com sucesso"})
}
