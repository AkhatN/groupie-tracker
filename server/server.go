package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"tracker/model"
)

var Mod model.Mode
var Welcome model.Welcome
var errtmpl *template.Template

var err error

func init() {
	errtmpl, err = template.ParseFiles("templates/error.html")
	if err != nil {
		log.Fatal(err)
	}
}

//Handle Main Page
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, Mod)
}

//Handle Artist Page
func artist(w http.ResponseWriter, r *http.Request) {
	ind, err := strconv.Atoi(r.RequestURI[9:])
	//Check url
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusBadRequest)
		return
	}
	//Check ID is existing
	if !GetUrlRightNumber(ind) {
		w.WriteHeader(http.StatusNotFound)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusNotFound)
		return
	}
	//Check Method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("templates/infoArtist.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, Mod.Art[ind-1])
}

func search(w http.ResponseWriter, r *http.Request) {
	//Checking url
	if r.URL.RequestURI() == "/search" {
		w.WriteHeader(http.StatusNotFound)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusMethodNotAllowed)
		return
	}

	param := r.URL.Query()
	word := param.Get("search")
	option := param.Get("option")

	//Check if there is right option and word is not empty
	if !CheckOptionWord(word, option) {
		w.WriteHeader(http.StatusBadRequest)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusBadRequest)
		return
	}

	var SearchBarArtist model.Mode

	//Init Searched Data to SearchBar struct
	SearchBar(word, option, &SearchBarArtist)

	if SearchBarArtist.Art == nil {
		w.WriteHeader(http.StatusOK)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusOK)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, SearchBarArtist)
}

func filter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusMethodNotAllowed)
		return
	}
	//init data in struct
	var Filter model.Filter
	var Filtered *model.Mode
	if er := Filter.Constructor(r); er != nil {
		w.WriteHeader(http.StatusOK)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusOK)
		return
	}
	Filtered = Filter.FilterArtists(&Mod)
	if Filtered.Art == nil {
		w.WriteHeader(http.StatusOK)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusOK)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errtmpl.ExecuteTemplate(w, "error.html", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, Filtered)
}

//Set Main SerVer
func Server() {
	go GetInfoJson()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//suggestion...
	http.HandleFunc("/", home)
	http.HandleFunc("/filter", filter)
	http.HandleFunc("/search", search)
	http.HandleFunc("/artists/", artist)
	fmt.Println("Port: 8080 is running!")
	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
		return
	}
}
