package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// @author valor.

func run(port int64) {
	http.HandleFunc("/", gofs)

	addr := "[::]:" + strconv.FormatInt(port, 10)
	_ = http.ListenAndServe(addr, nil)
}

type Meta struct {
	Path string
	Tags []template.HTML
}

func gofs(w http.ResponseWriter, r *http.Request) {
	reqPath := r.URL.Path

	filePath := BaseDir + reqPath
	logger.Printf("Request from [%s] was looking for: [%s]", r.RemoteAddr, filePath)

	// true  - existed
	// false - not existed
	f, exist := isExisted(filePath)

	// short HTTP
	w.Header().Set("Connection", "close")

	if !exist {
		// parse `404.gohtml`
		w.WriteHeader(http.StatusNotFound)
		_ = tmpl.ExecuteTemplate(w, _404_, nil)
	} else {
		if f.IsDir() {
			tags := readDir(filePath)

			// parse `gofs.gohtml`
			meta := Meta{
				Path: reqPath,
				Tags: tags,
			}
			w.WriteHeader(http.StatusOK)
			_ = tmpl.ExecuteTemplate(w, _gofs_, meta)
		} else {
			file, _ := os.Open(filePath)

			// download file
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Header().Set("Content-Length", strconv.FormatInt(f.Size(), 10))
			w.Header().Set("Accept-Ranges", "none")
			w.WriteHeader(http.StatusOK)

			_, _ = io.Copy(w, file)
			_ = file.Close()
		}
	}
}

func isExisted(filePath string) (os.FileInfo, bool) {
	// true  - existed
	// false - not existed
	exist := false

	file, err := os.Stat(filePath)
	if err == nil {
		exist = true
	} else {
		if os.IsExist(err) {
			exist = true
		}
	}
	return file, exist
}

func readDir(dirname string) []template.HTML {
	list, _ := ioutil.ReadDir(dirname)
	length := len(list)

	if length == 0 {
		// no files
		return []template.HTML{}
	}

	var dir = make([]template.HTML, 0, length)
	var file = make([]template.HTML, 0, length)

	var b strings.Builder

	for _, f := range list {
		var link = f.Name()
		var name = f.Name()
		var time = f.ModTime().Format(TimeFmt)
		var size = strconv.FormatInt(f.Size(), 10)

		if f.IsDir() {
			b.Reset()
			b.WriteString(link)
			b.WriteString("/")
			link = b.String()

			size = "-"

			if len(name)+1 > MaxFilenameLength {
				b.Reset()
				b.WriteString(name[:MaxFilenameLength-3])
				b.WriteString("../")
				name = b.String()
			} else {
				b.Reset()
				b.WriteString(name)
				b.WriteString("/")
				name = b.String()
			}
		} else {
			if len(name) > MaxFilenameLength {
				b.Reset()
				b.WriteString(name[:MaxFilenameLength-3])
				b.WriteString("..&gt;")
				name = b.String()
			}
		}

		if count := MaxFilenameLength - len(name); count > 0 {
			b.Reset()
			// repeat
			b.WriteString(" ")
			for b.Len() < count {
				if b.Len() <= count/2 {
					b.WriteString(b.String())
				} else {
					b.WriteString(b.String()[:count-b.Len()])
					break
				}
			}
			// :p
			b.WriteString(time)
			time = b.String()
		}

		// <a href="${link}">${name}</a> ${time} ${size}
		tag := template.HTML(
			fmt.Sprintf("<a href=\"%s\">%s</a> %s %19s", link, name, time, size))

		if f.IsDir() {
			dir = append(dir, tag)
		} else {
			file = append(file, tag)
		}
	}
	return append(dir, file...)
}
