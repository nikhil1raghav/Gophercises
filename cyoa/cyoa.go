package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"strings"
)

var StoryTmpl=`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Choose your own adventure</title>
</head>
<body>
<section class="page">
<h1>{{.Title}}</h1>
{{range .Paragraphs}}
<p>{{.}}</p>
{{end}}
<ul>
{{range .Options}}
    <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
{{end}}
</ul>
</section>
<style>
    body{
        font-family: helvetica, arial;
    }
    h1{
        text-align: center;
        position: relative;
    }
    .page{
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
    }
</style>
</body>
</html>
`

func NewHandler(s Story) *handler{
	return &handler{s, tpl}
}
func (h *handler) WithTemplate(tpl *template.Template) *handler{
	h.template=tpl
	return h
}
var tpl *template.Template
func init() {
	tpl=template.Must(template.New("").Parse(StoryTmpl))
}
type handler struct {
	story Story
	template *template.Template
}
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	path:=strings.TrimSpace(r.URL.Path)
	if path=="" || path=="/"{
		path="/intro"
	}
	path=path[1:]
	if chapter,ok:=h.story[path];ok{
		err:=h.template.Execute(w, chapter)
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found", http.StatusNotFound)


}

func GetStory(r io.Reader)(Story, error){
	d:=json.NewDecoder(r)
	var story Story
	if err:=d.Decode(&story);err!=nil{
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter
type Chapter struct {
	Title   string   `json:"title"`
	Paragraphs   []string `json:"story"`
	Options []Option`json:"options"`
}
type Option struct {
	Text string `json:"text"`
	Chapter string `json:"arc"`
}
