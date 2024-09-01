package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/sohailshaukat/rssagg/internal/database"
)


type User struct {
  ID        uuid.UUID `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
  APIKey    string `json:"api_key"`
}

func databaseUserToUser(dbUSer database.User) User {
  return User{
    ID: dbUSer.ID,
    CreatedAt: dbUSer.CreatedAt,
    UpdatedAt: dbUSer.UpdatedAt,
    Name: dbUSer.Name,
    APIKey: dbUSer.ApiKey,
  }
}
