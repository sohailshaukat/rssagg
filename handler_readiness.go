package main

import (
  "net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request){
  type successfulResponse struct {
    StatusCode int `json:"statusCode"`
    Message string `json:"message"`
  }

  respondWithJson(w, 200, successfulResponse{
    StatusCode: 200,
    Message: "Looking good, as it should.",
  })

}
