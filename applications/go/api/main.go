package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Estrutura do Usuário
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Manter em memória
// para simular o banco
// de forma segura
var (
	users = make(map[int]User)
	mu    sync.Mutex
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/user", getAllUsers)
	mux.HandleFunc("GET /v1/user/{id}", getUser)
	mux.HandleFunc("POST /v1/user", createUser)
	mux.HandleFunc("PUT /v1/user/{id}", updateUser)
	mux.HandleFunc("DELETE /v1/user/{id}", deleteUser)

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))
}

// Handlers
// getAllUsers
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var allUsers []User
	for _, user := range users {
		allUsers = append(allUsers, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allUsers)
}

// getUser
func getUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		http.Error(w, "ID inválido ou não fornecido", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	user, exists := users[id]
	if !exists {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// createUser
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[user.ID]; exists {
		http.Error(w, "Usuário já existe", http.StatusConflict)
		return
	}

	users[user.ID] = user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário criado com sucesso"})
}

// updateUser
func updateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		http.Error(w, "ID inválido ou não fornecido", http.StatusBadRequest)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[id]; !exists {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	users[id] = user
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário atualizado com sucesso"})
}

// deleteUser
func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		http.Error(w, "ID inválido ou não fornecido", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[id]; !exists {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	delete(users, id)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário deletado com sucesso"})
}
