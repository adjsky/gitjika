package handlers

import (
	"net/http"

	"github.com/adjsky/gitjika/ui/pages"
)

func Index(w http.ResponseWriter, r *http.Request) {
	pages.Index().Render(r.Context(), w)
}
