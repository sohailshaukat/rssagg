package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sohailshaukat/rssagg/internal/auth"
	"github.com/sohailshaukat/rssagg/internal/database"
)

func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request){
  type parameters struct {
    Name string `json:"name"`
  }

  decoder := json.NewDecoder(r.Body)
  params := parameters{}
  err := decoder.Decode(&params)
  if err != nil {
    respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
    return
  }

  user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
    ID: uuid.New(),
    CreatedAt: time.Now().UTC(),
    UpdatedAt: time.Now().UTC(),
    Name: params.Name,
  })
  if err != nil {
    respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
    return
  }

  type successfulResponse struct {
    StatusCode int `json:"statusCode"`
    Message string `json:"message"`
    User User `json:"user"`
  }

  respondWithJson(w, 201, successfulResponse{
    StatusCode: 201,
    Message: "User created.",
    User: databaseUserToUser(user),
  })

}

func (apiCfg *apiConfig)handlerGetUser(w http.ResponseWriter, r *http.Request){
  apiKey, err := auth.GetAPIKey(r.Header)
  
  if err != nil {
    respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
    return
  }

  user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
  if err != nil {
    respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
    return
  }

  respondWithJson(w, 200, databaseUserToUser(user))
}
