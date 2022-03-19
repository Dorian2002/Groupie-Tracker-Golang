package main

// import des packages
import (
	"html/template"
	"net/http"
	"./controllers"
)

// fonction main, lancement du serveur sur le port : 1111, éxécution des  fonctions ArtistPage et MoreContent
func main() {
	fs := http.FileServer(http.Dir("./template/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handler)
	http.HandleFunc("/artist", controllers.ArtistPage)
	http.HandleFunc("/morecontent", controllers.MoreContent)
	http.ListenAndServe("localhost:1111", nil)
}

// fonction handler appelé dans la fonction main, éxécution de la template html
func handler(w http.ResponseWriter, r *http.Request) {
	tM, _ := template.ParseFiles("template/index.html")
	tM.Execute(w, nil)
}
