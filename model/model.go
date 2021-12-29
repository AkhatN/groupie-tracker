package model

type Artist struct {
	ID           int       `json:"id"`
	Image        string    `json:"image"`
	Name         string    `json:"name"`
	Members      []string  `json:"members"`
	CreationDate int       `json:"creationDate"`
	FirstAlbum   string    `json:"firstAlbum"`
	Relations    Relations `json:"-"`
}

type Mode struct {
	Art          []Artist
	SugestionLoc []string
}

type Welcome struct {
	Relations []Relations `json:"index"`
}

type Relations struct {
	ID             int64               `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Filter struct {
	CreationDate FilterParam
	Album        FilterParam
	Member       FilterParam
	Location     FilterParamLoc
	// isCreationDate, isAlbum, isMembers, isLocation               bool
	// FromDate, ToDate, AlbumFrom, AlbumTo, MembersFrom, MembersTo int
	// Location                                                     string
}

type FilterParam struct {
	isClicked bool
	From      int
	To        int
}

type FilterParamLoc struct {
	isClicked bool
	Location  string
}
