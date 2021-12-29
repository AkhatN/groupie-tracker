package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"tracker/model"
	"tracker/url"
)

//Handle Data ...
// var Mod data.Mode
// var Welcome data.Welcome
// var errtmpl *template.Template
// var SearchBarArtist data.Mode
// var Filter data.Filter
// var err error

//Init locations: add data in suggestion from relation
func initSuggestion() {
	deleteDubl := make(map[string]bool) //[japan]false

	for i := range Mod.Art {
		for key := range Mod.Art[i].Relations.DatesLocations { //dateLoc[key string][]string
			if !deleteDubl[key] { //[saitamo-japan] = false
				Mod.SugestionLoc = append(Mod.SugestionLoc, key)
				deleteDubl[key] = true
			}
		}
	}
}

//Get information from urlArtist
//Handle Mod Artist ...
func GetInfoJson() {
	for {
		r, err := http.Get(url.UrlArtist)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(data, &Mod.Art)
		if nil != err {
			log.Fatal(err)
		}
		//Loading json data-relation in welcome struct
		GetInfoRelation()
		//Adding data to relations
		for i := range Welcome.Relations {
			Mod.Art[i].Relations = Welcome.Relations[i] //
		}
		initSuggestion()

		time.Sleep(time.Minute * 10)
	}
}

//Get Information from urlRelation ...
//Handle Data Welcome
//GetInfoRelation ...
func GetInfoRelation() {
	r, err := http.Get(url.UrlRelation)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &Welcome)
	if err != nil {
		log.Fatal(err)
	}
}

//GetUrlRightNumber ...
func GetUrlRightNumber(index int) bool {
	for i := range Mod.Art {
		if Mod.Art[i].ID == index {
			return true
		}
	}
	return false
}

// SearchBar Func ...
func SearchBar(words, option string, SearchBarArtist *model.Mode) {
	for i := range Mod.Art {
		if option == "name" && Mod.Art[i].Name == words {
			SearchBarArtist.Art = append(SearchBarArtist.Art, Mod.Art[i])
		} else if option == "locations" {
			for j := range Mod.Art[i].Relations.DatesLocations {
				if strings.ToLower(j) == strings.ToLower(words) {
					SearchBarArtist.Art = append(SearchBarArtist.Art, Mod.Art[i])
				}
			}
		} else if option == "albumdate" && Mod.Art[i].FirstAlbum == words {
			SearchBarArtist.Art = append(SearchBarArtist.Art, Mod.Art[i])
		} else if option == "creationdate" && strconv.Itoa(Mod.Art[i].CreationDate) == words {
			SearchBarArtist.Art = append(SearchBarArtist.Art, Mod.Art[i])
		} else if option == "members" {
			for j := range Mod.Art[i].Members {
				if strings.ToLower(Mod.Art[i].Members[j]) == strings.ToLower(words) {
					SearchBarArtist.Art = append(SearchBarArtist.Art, Mod.Art[i])
				}
			}
		}
	}
}

//CheckOptionWord ...
func CheckOptionWord(word, option string) bool {
	if option != "name" && option != "locations" && option != "albumdate" && option != "creationdate" && option != "members" {
		return false
	}

	if word == "" {
		return false
	}
	return true
}
