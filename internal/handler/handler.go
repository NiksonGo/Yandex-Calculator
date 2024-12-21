package handler

import (
 "encoding/json"
 "net/http"

 "github.com/NiksonGo/Yandex-Calculator/internal/calculator"
)

type Request struct {
 Expression string `json:"expression"`
}

type Response struct {
 Result *float64 `json:"result,omitempty"`
 Error  string   `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
 if r.Method != http.MethodPost {
  http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
  return
 }

 var req Request
 if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
  w.WriteHeader(http.StatusBadRequest)
  json.NewEncoder(w).Encode(Response{Error: "Invalid request payload"})
  return
 }

 result, err := calculator.Calc(req.Expression)
 if err != nil {
  if err.Error() == "invalid character in expression" {
   w.WriteHeader(http.StatusUnprocessableEntity)
   json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
   return
  }
  w.WriteHeader(http.StatusInternalServerError)
  json.NewEncoder(w).Encode(Response{Error: "Internal server error"})
  return
 }

 w.WriteHeader(http.StatusOK)
 json.NewEncoder(w).Encode(Response{Result: &result})
}
