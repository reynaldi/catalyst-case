package request

import (
	"net/http"
	"strings"
)

func GetPathValue(r *http.Request, trim string) string {
	trimmed := strings.TrimPrefix(r.URL.Path, trim)
	return trimmed
}
