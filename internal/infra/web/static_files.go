package web

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static/index.html static/styles.css
var staticFiles embed.FS

func StaticFileSystem() http.FileSystem {
	static, err := fs.Sub(staticFiles, "static")
	if err != nil {
		panic(err)
	}

	return http.FS(static)
}
