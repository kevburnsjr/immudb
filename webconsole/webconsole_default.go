// +build !webconsole
//go:generate go run github.com/rakyll/statik -f -src=./default

package webconsole

import (
	"github.com/codenotary/immudb/pkg/logger"
	"net/http"
	// embedded static files
	_ "github.com/codenotary/immudb/webconsole/statik"
	"github.com/rakyll/statik/fs"
)

func SetupWebconsole(mux *http.ServeMux, l logger.Logger, addr string) error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}
	l.Infof("Webconsole not built-in")
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/missing/", http.StatusMovedPermanently)
	})
	mux.Handle("/missing/", http.FileServer(statikFS))
	return nil
}
