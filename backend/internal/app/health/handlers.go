package health

import (
	"net/http"

	"github.com/go-chi/render"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"status": "ok"})
}
