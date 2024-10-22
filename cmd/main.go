package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sritejachilakapati/movietix/internal/config"
	"github.com/sritejachilakapati/movietix/internal/database"
	"github.com/sritejachilakapati/movietix/internal/repository"
)

func contentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

type contextKey string

const paramsKey contextKey = "params"

func mapQueryParamsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := make(map[string][]string)
		for key, value := range r.URL.Query() {
			params[key] = value
		}

		ctx := context.WithValue(r.Context(), paramsKey, params)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type jsonResponseWriter struct {
	http.ResponseWriter
}

func (w jsonResponseWriter) WriteJson(statusCode int, data interface{}) (int, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	w.WriteHeader(statusCode)
	return w.Write(jsonBytes)
}

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := mux.NewRouter()
	r.Use(contentTypeApplicationJsonMiddleware, mapQueryParamsMiddleware)
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		jw := jsonResponseWriter{w}

		ctx := r.Context()
		conn := database.Connect(ctx)
		defer conn.Close(ctx)

		params := r.Context().Value(paramsKey).(map[string][]string)
		fmt.Printf("Params: %v\n", params)

		userRepo := repository.New(conn)
		users, err := userRepo.GetAllUsers(ctx)
		if err != nil {
			log.Fatalf("Error getting movies: %v", err)
		}

		jw.WriteJson(http.StatusOK, users)
	}).Methods(http.MethodGet)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
