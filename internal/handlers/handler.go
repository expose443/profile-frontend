package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/with-insomnia/profile-frontend/internal/model"
)

func NewHandler() *Handlers {
	return &Handlers{}
}

type Handlers struct{}

func (h *Handlers) LoginGet(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./ui/html/login.html")
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = template.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) LoginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	credintails := model.Credintails{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	body, err := json.Marshal(credintails)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	req, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	fmt.Println(string(b))
}
