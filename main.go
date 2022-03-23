package main

import (
	"net/http"
	"os"
	"strconv"
	"text/template"
)

/* func callApi(urlApi string, data interface{}) {
	// Initialisation du client qui va emmettre/demander les requetes !
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	//Creation de la requete HTTP vers l'api
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		log.Fatalln(errReq)

	}
	temps := time.Now().String()
	println(temps)
	//Premet d'eviter la limit d'appelle à l'api ! Il faut le changer de temps en temps en cas de coupure !
	req.Header.Set("User_Agent", temps)

	//Envoie de la requete HTTP vers l'api !
	res, errRes := httpClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		log.Printf(errRes.Error())
		//log.Fatalln(errRes)
	}

	// Lecture et recuperation du corps de la requete HTTP !
	body, _ := ioutil.ReadAll(res.Body)

	//Decodage de l'api et envoie des données à la structure !
	json.Unmarshal(body, data)
	println("----------------------------")
} */

func main() {
	// ================= Intialisation des APIs =================

	// Intialisation et appelle à l'API principale
	var main []Artist
	urlAPI := "https://groupietrackers.herokuapp.com/api/artists"
	callApi(urlAPI, &main)

	// Gestion et traitement des données de l'API pricipale
	var artists []Artist
	var groupes []Artist
	artists, groupes = DataGestionP(main)

	// Intialisation et appelle des l'APIs secondaires (des dates et des localisations des concerts)
	var date Date
	urlDate := "https://groupietrackers.herokuapp.com/api/dates"
	callApi(urlDate, &date)

	var location Location
	urlLocation := "https://groupietrackers.herokuapp.com/api/locations"
	callApi(urlLocation, &location)

	// Gestion et traitement des données recupere par le biais des l'APIs secondaires
	var dateLocations []DateLocations
	dateLocations = DataGestionS(date, location)

	// ================= Serveur Golang =================

	// Exploitation des templates en parsent le fichier "temp"
	templ, _ := template.ParseGlob("./temp/*.html")

	// Mise en place des diffrentes routes pour le site web !
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		println("Route Principale")
		dataTemp := IndexGenre{"", main}
		if main == nil {
			println("Rappele de l'api principale pour remplire main !")
			callApi(urlAPI, &main)
			artists, groupes = DataGestionP(main)
			dataTemp = IndexGenre{"", main}
		}
		templ.ExecuteTemplate(rw, "Index", dataTemp)
	})

	http.HandleFunc("/groupes", func(rw http.ResponseWriter, r *http.Request) {
		println("Route Groupes")
		dataTemp := IndexGenre{"groupe", groupes}
		if groupes == nil {
			println("Rappele de l'api principale pour remplire groupe !")
			callApi(urlAPI, &main)
			artists, groupes = DataGestionP(main)
			dataTemp = IndexGenre{"groupe", groupes}
		}
		templ.ExecuteTemplate(rw, "Genre", dataTemp)
	})

	http.HandleFunc("/artistes", func(rw http.ResponseWriter, r *http.Request) {
		println("Route Artistes")
		dataTemp := IndexGenre{"artiste soliste", artists}
		if artists == nil {
			println("Rappele de l'api principale pour remplire artists !")
			callApi(urlAPI, &main)
			artists, groupes = DataGestionP(main)
			dataTemp = IndexGenre{"artiste soliste", artists}
		}
		templ.ExecuteTemplate(rw, "Genre", dataTemp)
	})

	// Créeation des routes pour chaque Artiste/Groupe selon le nombre !
	for _, slecArtist := range main {
		link := "/" + strconv.Itoa(slecArtist.Id)
		//println(slecArtist.Name, slecArtist.Id, link)
		http.HandleFunc(link, func(rw http.ResponseWriter, r *http.Request) {
			indexRep, _ := strconv.Atoi(link[1:])
			var typeCard string
			println(indexRep)
			for indexArtistes := range artists {
				if artists[indexArtistes].Id == main[indexRep-1].Id {
					typeCard = " de l'artiste soliste"
					break
				} else {
					typeCard = " du groupe"
				}
			}
			Cards := Cards{typeCard, main[indexRep-1], dateLocations[indexRep-1]}
			templ.ExecuteTemplate(rw, "Cards", Cards)
		})
	}

	println("fin de chargement du serveur !")
	println("lien du serveur : localhost:8080 ")

	// Mise en place du fichier static (qui permet de parteger des fichiers avec le front --')
	root, _ := os.Getwd()
	fileServe := http.FileServer(http.Dir(root + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static", fileServe))

	// Lien du serveur (local !)
	http.ListenAndServe("localhost:8080", nil)
}
