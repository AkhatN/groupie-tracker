package model

import (
	"errors"
	"math"
	"net/http"
	"strconv"
	"strings"
)

//Filter Method Func ...
func (f *Filter) Constructor(r *http.Request) error {

	// Reading CreationDate from client
	if r.FormValue("isCreationDate") == "true" {
		from, to := r.FormValue("creationFrom"), r.FormValue("creationTo")
		if err := f.CreationDate.InitFilterParam(from, to); err != nil {
			return err
		}

	}

	// Reading AlbumDate from client
	if r.FormValue("isAlbum") == "true" {
		from, to := r.FormValue("albumFrom"), r.FormValue("albumTo")
		if err := f.CreationDate.InitFilterParam(from, to); err != nil {
			return err
		}
	}

	// Reading AmountMembers from client
	if r.FormValue("isMembers") == "true" {
		from, to := r.FormValue("membersFrom"), r.FormValue("membersTo")
		if err := f.CreationDate.InitFilterParam(from, to); err != nil {
			return err
		}
	}

	// Reading Location from client
	if r.FormValue("isLocation") == "true" {
		if r.FormValue("locationFilter") != "" {
			f.Location.isClicked = true
			temp := r.FormValue("locationFilter")
			temp = strings.Replace(temp, ", ", "-", -1)
			f.Location.Location = strings.Replace(temp, " ", "_", -1)
		} else {
			f.Location.isClicked = false
		}
	}
	return nil
}

//Filtered Data
func (f *Filter) FilterArtists(Mod *Mode) *Mode {
	result := &Mode{}
	for i := range Mod.Art {
		if f.CreationDate.isClicked && (Mod.Art[i].CreationDate < f.CreationDate.From || Mod.Art[i].CreationDate > f.CreationDate.To) {
			continue
		}
		if albumDate, _ := strconv.Atoi(Mod.Art[i].FirstAlbum[6:]); f.Album.isClicked && (albumDate < f.Album.From || albumDate > f.Album.To) {
			continue
		}
		if f.Member.isClicked && (len(Mod.Art[i].Members) < f.Member.From || len(Mod.Art[i].Members) > f.Member.To) {
			continue
		}
		if f.Location.isClicked {
		loop:
			for j := range Mod.Art[i].Relations.DatesLocations {
				if j == strings.ToLower(f.Location.Location) || strings.Contains(j, strings.ToLower(f.Location.Location)) {
					result.Art = append(result.Art, Mod.Art[i])
					break loop // new-york, usa = strings.Replace new-york-usa
				}
			}
			continue
		}
		result.Art = append(result.Art, Mod.Art[i])
	}
	return result
}

func (f *FilterParam) InitFilterParam(from, to string) error {
	f.isClicked = true

	// Handle from ...
	if from == "" {
		f.From = 0
	} else {
		fromDate, err := strconv.Atoi(from)
		if err != nil {
			return err
		}
		f.From = fromDate
	}

	// Handle to ...
	if to == "" {
		f.To = math.MaxInt32
	} else {
		toDate, err := strconv.Atoi(to)
		if err != nil {
			return err
		}
		f.To = toDate
	}
	// Check if from is bigger than to ... then return error
	if f.From > f.To {
		return errors.New("from is greater than to")
	}
	return nil
}
