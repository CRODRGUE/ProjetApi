package main

// !!! il faut que les variables de la sturct commence par une MAJ
// sinon elle ne sont pas exporter vers template !!!

//Structure qui repond à l'api principale !
type Artist struct {
	Name      string   `json:"name"`
	Id        int      `json:"id"`
	Image     string   `json:"image"`
	Members   []string `json:"members"`
	DateCrea  int      `json:"creationDate"`
	DateAlbum string   `json:"firstAlbum"`
}

// /Structure qui repondent aux l'apis secondaires !
type Date struct {
	Index []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Location struct {
	Index []struct {
		Id       int      `json:"id"`
		Location []string `json:"locations"`
	} `json:"index"`
}

// Structures qui permette le traitement des données des apis secondaires !
type Contenu struct {
	Location string
	Dates    string
}

type DateLocations struct {
	Id      int
	Contenu []Contenu
}

// Declaration de la structure pour passer les données aux templates :
// Index et Genre (les routes "/", "/artistes" puis "/groupes")
type IndexGenre struct {
	Cat   string
	Liste []Artist
}

// Declaration de la structure pour passer les données au template :
// Cards qui se trouve sur la route "/id" (avec id egale au numero de l'artiste ou du groupe)
type Cards struct {
	Type    string
	Artist  Artist
	Contenu DateLocations
}
