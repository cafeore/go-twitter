package template

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)
type Template struct{
	Templates *template.Template
}
func NewTemplates(fileNames string)*Template{
	return &Template{
		template.Must(template.ParseGlob(fileNames)),		
	}
}
func (t *Template) Render(w io.Writer,name string,data interface{},c echo.Context)error{
	return t.Templates.ExecuteTemplate(w,name,data)
}