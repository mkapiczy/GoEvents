package main

import (
	"net/http"
	"html/template"
	"os"
	"fmt"
	"path/filepath"
)

var (
	mainPath, _   = os.Getwd()
	templatesPath = filepath.Join(mainPath, "templates")
)

func mainView(responseWriter http.ResponseWriter, request *http.Request) {
	render(responseWriter, "main", ViewData{})
}

func getPlacesActionHandler(responseWriter http.ResponseWriter, request *http.Request) {
	latitude := request.FormValue("latitude")
	longitude := request.FormValue("longitude")
	distance := request.FormValue("distance")
	query := request.FormValue("query")

	places := getPlacesByLocation(latitude, longitude, distance, query)

	render(responseWriter, "main", ViewData{Places: places})
}

func render(writer http.ResponseWriter, name string, data interface{}) {
	tmpl, err := template.ParseGlob(templatesPath + "\\*.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(writer, name, data)
}

func main() {
	fmt.Println("Start")
	fmt.Print(mainPath)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", mainView)
	http.HandleFunc("/getPlaces", getPlacesActionHandler)

	http.ListenAndServe(":8000", nil)

}
