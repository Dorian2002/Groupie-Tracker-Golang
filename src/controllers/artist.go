package controllers

// import des packages
import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// structure dédiée à l'API
type Artist []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `son:"concertDates"`
	Relations    string   `json:"relations"`
}

// fonction ArtistPage, appel de l'API puis unmarshall et stockage de l'API pour l'éxécution de la template avec le contenu récupéré.
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	resp, err1 := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	var read Artist

	body, err3 := ioutil.ReadAll(resp.Body)

	err2 := json.Unmarshal(body, &read)

	if err1 != nil {
		fmt.Println(err1)
	} else if err2 != nil {
		fmt.Println(err2)
	} else if err3 != nil {
		fmt.Println(err3)
	}

	tA, _ := template.ParseFiles("template/artist.html")
	tA.Execute(w, read)
}
