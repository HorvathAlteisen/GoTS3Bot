package router

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

type staticFiles struct {
	directories []string
}

func Static(directories ...string) *staticFiles {
	s := staticFiles{}
	s.directories = append(s.directories, directories...)
	if len(s.directories) <= 0 {
		s.directories = append(s.directories, "public")
	}
	return &s
}

func (s *staticFiles) serverStaticFiles(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	if r.Method != "GET" && r.Method != "HEAD" {
		return true
	}

	for i, dir := range s.directories {
		file := dir + r.URL.Path
		if strings.HasSuffix(r.URL.Path, "/") {
			file = file + "index.html"
		}

		if _, err := os.Stat(file); err != nil {
			if i+1 == len(s.directories) {
				return true
			}
			continue
		}
		http.ServeFile(w, r, file)
		return false
	}
	return true
}
