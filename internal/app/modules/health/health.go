package health

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"simple-leave-tracker/internal/app"
)

const moduleName = "health"

type Health struct {
	log *slog.Logger
}

func New(log *slog.Logger) *Health {
	return &Health{
		log: log.With(
			"module",
			moduleName,
		),
	}
}

func (*Health) Name() string {
	return moduleName
}

func (h *Health) Register() app.Routes {
	return app.Routes{
		http.MethodGet: map[app.Path]http.HandlerFunc{
			"/health": h.HealthGet(),
		},
		http.MethodPost: map[app.Path]http.HandlerFunc{
			"/health": h.HealthPost(),
		},
	}
}

func (h *Health) HealthPost() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *Health) HealthGet() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		msg := map[string]string{
			"health": "healthy",
		}

		data, err := json.Marshal(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			h.log.Error("error marshaling", "message", msg, "error", err)

			return
		}

		if _, err := w.Write(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			h.log.Error("error writing response", "error", err)

			return
		}
	}
}
