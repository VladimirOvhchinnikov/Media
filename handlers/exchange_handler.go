package handlers

import (
	"Media/models"
	"Media/usecase"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Handlers struct {
	Logger  *logrus.Logger
	cash    models.CashInfo
	usecase *usecase.Case
}

func NewHandlers(logger *logrus.Logger, usecase *usecase.Case) *Handlers {
	return &Handlers{
		Logger:  logger,
		cash:    models.CashInfo{},
		usecase: usecase,
	}
}

func (h *Handlers) CulcHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&h.cash)
	if err != nil {
		h.Logger.Error("Error in decoding ", err)
		h.sendResponse(w, "Error in decoding: "+err.Error(), http.StatusBadRequest)
		return
	}

	h.usecase.Cash = h.cash
	responseData := h.usecase.Culc()

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		h.Logger.Error("Error in marshaling response", err)
		h.sendResponse(w, "Error in forming response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	h.sendJsonResponse(w, jsonData, http.StatusOK)
}

type ApiResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func (h *Handlers) sendResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ApiResponse{
		Message: message,
		Success: statusCode >= 200 && statusCode < 300,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		h.Logger.Error("Не получилось создать ответ", err)
	}
}

func (h *Handlers) sendJsonResponse(w http.ResponseWriter, jsonData []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}
