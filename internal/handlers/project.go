package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/with-insomnia/profile-frontend/internal/model"
)

func (h *Handlers) ProjectGet(w http.ResponseWriter, r *http.Request) {
	req, err := http.Get("http://localhost:8080/project")
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()
	var project []model.Project
	err = json.NewDecoder(req.Body).Decode(&project)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "index.html", project)
}

func (h *Handlers) ProjectPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	project := model.Project{
		Title:       r.FormValue("title"),
		Description: r.FormValue("desc"),
		GithubLink:  r.FormValue("git"),
		Image:       r.FormValue("image"),
	}
	body, err := json.Marshal(project)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	c, err := r.Cookie("jwt_token")
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusUnauthorized)
		return
	}
	fmt.Println(c.Value)
	req, err := http.NewRequest("POST", "http://localhost:8080/project", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(c)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handlers) CreateProject(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "create-project-form.html", nil)
}

func renderTemplate(w http.ResponseWriter, html string, data any) {
	temp, err := template.ParseFiles("./ui/html/" + html)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = temp.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
