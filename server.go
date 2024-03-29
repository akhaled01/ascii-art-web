package webart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

/*
This file is what could be considered the router. it redirects all methods sent by the client, and
returns appropriate responses and response codes
*/

// * the code below initializes all HTML templates in the ascii-art-web/templates directory, including css
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("../templates/*"))
}

// * we create 2 structs. One Unmarshal JSON data sent by the client to the golang server as golang objects
type ASCII_ART struct {
	Text     string `json:"Text"`
	Banner   string `json:"Banner"`
	Newcolor string `json:"Newcolor"`
}

// * and one represents the result string
type RESULT_ASCII_ART struct {
	Result     string
	ApplyColor string
}

//* this function is responsible for processing the GET request for the main page in the case it is requested.
//* is returns a Bad request error in the case something else rather than "/" is typed

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//* if the HTTP method is GET, serve the HTML files in the template directory, otherwise, serve the
	//* custom HTML for bad requests
	case "GET":
		r.ParseForm()
		path := "../templates" + r.URL.Path
		http.ServeFile(w, r, path)
	default:
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "../templates/badrequest.html")
	}
}

//* this function basically appends the ascii art to the result div in HTML

func Gen_ASCII(w http.ResponseWriter, r *http.Request) {
	//
	var ascii_art ASCII_ART
	var ascii_result RESULT_ASCII_ART

	if r.Method != http.MethodPost {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&ascii_art)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(ascii_art.Text) > 255 {
		http.Error(w, "Too much letters خوك", http.StatusBadRequest)
		return

	}
	if MapFont(ascii_art.Banner) == "" {
		http.Error(w, "Invalid Banner Type", http.StatusNotFound)
		http.ServeFile(w, r, "../templates/badrequest.html")
		return
	}
	ascii_result.Result = PrintART(ascii_art.Text, ascii_art.Banner)
	ascii_result.ApplyColor = ascii_art.Newcolor
	jsonENC := json.NewEncoder(w)
	jsonENC.Encode(ascii_result)
	fmt.Println((ascii_art.Text))
	fmt.Println((ascii_art.Banner))
	fmt.Println((ascii_art.Newcolor))
}
