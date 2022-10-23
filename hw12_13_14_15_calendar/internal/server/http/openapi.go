package internalhttp

import (
	"github.com/AZhur771/otus-go-homework/hw12_13_14_15_calendar/third_party"
	"io/fs"
	"mime"
	"net/http"
)

// GetOpenAPIHandler serves an OpenAPI UI.
func GetOpenAPIHandler() (http.Handler, error) {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPIV2, "openapiv2")
	if err != nil {
		return nil, err
	}
	return http.FileServer(http.FS(subFS)), nil
}
