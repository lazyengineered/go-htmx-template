package main

import (
    "embed"
    "html/template"
)

var (
    //go:embed templates/*.html
    templatesFS embed.FS
)

func NewTemplate(names ...string) *template.Template {
    return template.Must(template.ParseFS(templatesFS, templatePaths(names)...))
}

func templatePaths(names []string) []string {
    var paths []string

    for _, p := range names {
        paths = append(paths, "templates/"+p+".html")
    }

    return paths
}
