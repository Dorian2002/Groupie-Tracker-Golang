package controllers

// import des packages
import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"reflect"
)

// structure 1 dédiée à l'API
type Artist2 struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relation     Relation
	Concert		 []string
}

// structure 2 dédiée à l'API
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// fonction MoreContent, appel des parties de l'API puis unmarshall et stockage de l'API. Vérification des erreurs puis éxécution de la template avec le contenu récupéré.
func MoreContent(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	resp, err1 := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	resp2, err2 := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	var read  Artist2
	var read2 Relation

	body, err3 := ioutil.ReadAll(resp.Body)
	body2, err4 := ioutil.ReadAll(resp2.Body)

	err6 := json.Unmarshal(body2, &read2)
	err5 := json.Unmarshal(body, &read)

	keys := reflect.ValueOf(read2.DatesLocations).MapKeys()

	read = Artist2.structRelation(read, read2)
	read = Artist2.structConcert(read, keys)

	if err1 != nil {
		fmt.Println(err1)
		fmt.Println("error 1")
	} else if err2 != nil {
		fmt.Println(err2)
		fmt.Println("error 2")
	} else if err3 != nil {
		fmt.Println(err3)
		fmt.Println("error 3")
	} else if err4 != nil {
		fmt.Println(err4)
		fmt.Println("error 4")
	} else if err5 != nil {
		fmt.Println(err5)
		fmt.Println("error 5")
	} else if err6 != nil {
		fmt.Println(err6)
		fmt.Println("error 6")
	}

	tMc, _ := template.ParseFiles("template/morecontent.html")
	tMc.Execute(w, read)
}

// fonction pour mettre toutes les informations récupérées dans une seule structure
func (A Artist2) structRelation(read2 Relation) Artist2 {
	A.Relation = read2
	return A
}

// fonction pour mettre toutes les informations récupérées dans une seule structures
func (A Artist2) structConcert(keys []reflect.Value) Artist2 {
	lenK := 0
	for _,key := range keys{
		A.Concert = append(A.Concert, key.String())
		lenK += 1
	}
	return A
	
}