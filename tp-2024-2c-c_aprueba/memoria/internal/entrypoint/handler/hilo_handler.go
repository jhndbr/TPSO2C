package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sisoputnfrba/tp-golang/memoria/internal/core/usecase"
)

type Hilo struct {
	PID      uint32 `json:"pid"`
	FilePath string `json:"file_path"`
}

func CrearHiloHandler(w http.ResponseWriter, r *http.Request) {

	var request Hilo
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	err = usecase.CrearHilo(request.PID, request.FilePath)
	if err != nil {
		http.Error(w, "Error creando el hilo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func FinalizarHiloHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		PID int `json:"pid"`
		TID int `json:"tid"`
	}

	// Decodificar el cuerpo de la solicitud en JSON
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Cuerpo de la solicitud inválido", http.StatusBadRequest)
		return
	}
	pid := requestBody.PID
	tid := requestBody.TID

	err := usecase.EliminarHilo(uint32(pid), uint32(tid))
	if err != nil {
		http.Error(w, "Error finalizando el hilo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"OK"}`))

}
