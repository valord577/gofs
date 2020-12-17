package main

import (
	"html/template"
)

// @author valor.

const (
	_404_  = "404"
	_gofs_ = "gofs"
)

// cache templates
var tmpl = cacheTemplates()

func cacheTemplates() *template.Template {
	tmpl, _ := template.New(_404_).Parse(page404())
	_, _ = tmpl.New(_gofs_).Parse(pageGofs())

	return tmpl
}

func page404() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <link rel="icon" href="data:image/ico;base64,aWNv">
    <title>404 Not Found</title>
</head>

<body>
<div style="text-align: center;">
    <h1>404 Not Found</h1>
</div>
<hr>
<div style="text-align: center;">
    <a href="https://github.com/valord577/gofs" target="_blank">gofs</a>
</div>
</body>
</html>
`
}

func pageGofs() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <link rel="icon" href="data:image/ico;base64,aWNv">
    <title>Index of {{ .Path }}</title>
</head>

<body>
<h1>Index of {{ .Path }}</h1>
<p>Powered by <a href="https://github.com/valord577/gofs" target="_blank">gofs</a></p>
<hr><pre><a href="../">../</a>
{{ range $index, $tag := .Tags }}{{ $tag }}
{{ end }}</pre><hr>
</body>
</html>
`
}
