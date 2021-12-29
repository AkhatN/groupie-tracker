package data

import (
	"errors"
	"math"
	"net/http"
	"strconv"
	"strings"
)

//Filter Method Func ...
func (f *Filter) Constructor(r *http.Request) error {
	if r.FormValue("isCreationDate") == "true" {
		f.isCreationDate = true
		if r.FormValue("creationFrom") == "" {
			f.FromDate = 0
		} else {
			fromDate, err := strconv.Atoi(r.FormValue("creationFrom"))
			if err != nil {
				return err
			}
			f.FromDate = fromDate
		}

		if r.FormValue("creationTo") == "" {
			f.ToDate = math.MaxInt32
		} else {
			toDate, err := strconv.Atoi(r.FormValue("creationTo"))
			if err != nil {
				return err
			}
			f.ToDate = toDate
		}
		if f.FromDate > f.ToDate {
			return errors.New("from is greater than to")
		}
	}
	if r.FormValue("isAlbum") == "true" {
		f.isAlbum = true
		if r.FormValue("albumFrom") == "" {
			f.AlbumFrom = 0
		} else {
			fromDate, err := strconv.Atoi(r.FormValue("albumFrom"))
			if err != nil {
				return err
			}
			f.AlbumFrom = fromDate
		}

		if r.FormValue("albumTo") == "" {
			f.AlbumTo = math.MaxInt32
		} else {
			toDate, err := strconv.Atoi(r.FormValue("albumTo"))
			if err != nil {
				return err
			}
			f.AlbumTo = toDate
		}
		if f.AlbumFrom > f.AlbumTo {
			return errors.New("from is greater than to")
		}
	}
	if r.FormValue("isMembers") == "true" {
		f.isMembers = true
		if r.FormValue("membersFrom") == "" {
			f.MembersFrom = 0
		} else {
			from, err := strconv.Atoi(r.FormValue("membersFrom"))
			if err != nil {
				return err
			}
			f.MembersFrom = from
		}

		if r.FormValue("membersTo") == "" {
			f.MembersTo = math.MaxInt32
		} else {
			to, err := strconv.Atoi(r.FormValue("membersTo"))
			if err != nil {
				return err
			}
			f.MembersTo = to
		}
		if f.MembersFrom > f.MembersTo {
			return errors.New("from is greater than to")
		}
	}
	if r.FormValue("isLocation") == "true" {
		if r.FormValue("locationFilter") != "" {
			f.isLocation = true
			temp := r.FormValue("locationFilter")
			f.Location = strings.Replace(temp, ", ", "-", -1)
		} else {
			f.isLocation = false
		}
	}
	return nil
}

//Filtered Data
func (f *Filter) GetArtists(Mod *Mode) *Mode {
	result := &Mode{}
	for i := range Mod.Art {
		if f.isCreationDate && (Mod.Art[i].CreationDate < f.FromDate || Mod.Art[i].CreationDate > f.ToDate) {
			continue
		}
		if albumDate, _ := strconv.Atoi(Mod.Art[i].FirstAlbum[6:]); f.isAlbum && (albumDate < f.AlbumFrom || albumDate > f.AlbumTo) {
			continue
		}
		if f.isMembers && (len(Mod.Art[i].Members) < f.MembersFrom || len(Mod.Art[i].Members) > f.MembersTo) {
			continue
		}
		if f.isLocation {
		loop:
			for j := range Mod.Art[i].Relations.DatesLocations {
				if j == strings.ToLower(f.Location) || strings.Contains(j, strings.ToLower(f.Location)) {
					result.Art = append(result.Art, Mod.Art[i])
					break loop
				}
			}
			continue
		}
		result.Art = append(result.Art, Mod.Art[i])
	}
	return result
}
