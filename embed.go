package frontend

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed ui/dist
var ui embed.FS

func DistFS() http.FileSystem {
	fsys, err := fs.Sub(ui, "ui/dist")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
