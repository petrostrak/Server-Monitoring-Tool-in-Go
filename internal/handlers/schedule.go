package handlers

import (
	"net/http"

	"github.com/petrostrak/Server-Monitoring-Tool-in-Go/internal/helpers"
)

// ListEntries lists schedule entries
func (repo *DBRepo) ListEntries(w http.ResponseWriter, r *http.Request) {
	err := helpers.RenderPage(w, r, "schedule", nil, nil)
	if err != nil {
		printTemplateError(w, err)
	}
}
