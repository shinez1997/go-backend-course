package handlers

import (
	"net/http"
	"server/models"
)

func SignUpHanlder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	models.IsExisted()
}
