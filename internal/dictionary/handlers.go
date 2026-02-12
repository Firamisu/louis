package dictionary

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/firamisu/louis/internal/views"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Word(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	word := r.PathValue("word")
	if len(word) == 0 {
		w.WriteHeader(400)
		slog.Error("cannot get path value")
		return
	}

	entry, err := h.service.GetWord(ctx, word)
	if err != nil {
		if errors.Is(err, ERR_NOT_FOUND) {
			w.WriteHeader(404)
			comp := views.NotFound(ERR_NOT_FOUND.Error())
			comp.Render(ctx, w)
			return
		}

		w.WriteHeader(500)
		return
	}

	comp := views.Entry(entry)
	comp.Render(ctx, w)
}
