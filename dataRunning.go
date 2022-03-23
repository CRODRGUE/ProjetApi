package main

import "strings"

// Permet de separer les atristes solistes et les groupe en deux parties
// qui sont envoye par l'api principale
func DataGestionP(main []Artist) ([]Artist, []Artist) {
	var artists []Artist
	var groupes []Artist
	for indexMain, contentsMain := range main {
		if len(contentsMain.Members) > 1 {
			groupes = append(groupes, main[indexMain])
		} else {
			artists = append(artists, main[indexMain])
		}
	}
	return artists, groupes
}

// Permet de faire le lien entre les apis secondaires et principale
// la fonction met en lien les dates et lieux de concerts avec les artistes/groupes
func DataGestionS(date Date, location Location) []DateLocations {
	var dateLocations []DateLocations
	for indexLocation := range location.Index {
		var contenu []Contenu
		joinTab := strings.Join(date.Index[indexLocation].Dates, " ")
		splitStr := strings.Split(joinTab, "*")
		splitStr = splitStr[1:]
		for i, cnt := range location.Index[indexLocation].Location {
			//println(cnt, splitStr[i], i)
			contenu = append(contenu, Contenu{cnt, splitStr[i]})
		}
		temp := DateLocations{location.Index[indexLocation].Id, contenu}
		dateLocations = append(dateLocations, temp)
	}
	return dateLocations
}
