package main

import (
	"net/http"
	"html/template"
	"os"
	"fmt"
	"path/filepath"
	"strings"
	"strconv"

	json "github.com/bitly/go-simplejson"
)

var (
	config        = loadConfiguration()
	mainPath, _   = os.Getwd()
	templatesPath = filepath.Join(mainPath, "templates")
)

func mainView(responseWriter http.ResponseWriter, request *http.Request) {
	render(responseWriter, "main", ViewData{})
}

func getPlacesActionHandler(responseWriter http.ResponseWriter, request *http.Request) {
	city := request.FormValue("city")
	query := request.FormValue("query")
	searchDistance := "20000" // TODO: Hardcoded for now. Find better solution.

	city = city[:strings.Index(request.FormValue("city"), ",")]

	googleApiUrlWithCity := fmt.Sprintf(config.GoogleApiUrl, city)

	if resp, err := http.Get(googleApiUrlWithCity); err == nil {
		responseJson, err := json.NewFromReader(resp.Body)

		if err != nil {
			fmt.Println("Error parsing response body:", err.Error())
			return
		}

		latitude, _ := responseJson.Get("results").GetIndex(0).Get("geometry").Get("location").Get("lat").Float64()
		longitude, _ := responseJson.Get("results").GetIndex(0).Get("geometry").Get("location").Get("lng").Float64()

		// TODO: Remove for testing purposes
		fmt.Println(latitude)
		fmt.Println(longitude)
		fmt.Println(query)

		places := getPlacesByLocation(strconv.FormatFloat(latitude, 'f', 6, 64), strconv.FormatFloat(longitude, 'f', 6, 64),
			searchDistance, query)

		render(responseWriter, "main", ViewData{Places: places})
		return

	}
	render(responseWriter, "main", ViewData{})
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
