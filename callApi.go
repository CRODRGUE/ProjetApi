package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Permette d'effectuer un appelle à une api avec une methode "get"
func callApi(urlApi string, data interface{}) {
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
}
