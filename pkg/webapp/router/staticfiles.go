package router

import (
	"net/http"
	"os"
	"strings"
)

func NewStatic(directories ...string) Middleware {
	if len(directories) <= 0 {
		directories = append(directories, "public")
	}
	return Middleware(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r.Method != "GET" && r.Method != "HEAD" {
			next(w, r)
			return // bail out.
		}

		for i, dir := range directories {
			file := dir + r.URL.Path
			if strings.HasSuffix(r.URL.Path, "/") {
				file = file + "index.html"
			}

			if _, err := os.Stat(file); err != nil {
				if i+1 == len(directories) {
					next(w, r)
					return
				}
				continue
			}
			http.ServeFile(w, r, file)
			return
		}
		next(w, r)
	})
}
