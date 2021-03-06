package form

import (
	"time"
)

// PhotoSearch represents search form fields for "/api/v1/photos".
type PhotoSearch struct {
	Query     string    `form:"q"`
	ID        string    `form:"id"`
	Title     string    `form:"title"`
	Hash      string    `form:"hash"`
	Duplicate bool      `form:"duplicate"`
	Archived  bool      `form:"archived"`
	Error     bool      `form:"error"`
	Lat       float32   `form:"lat"`
	Lng       float32   `form:"lng"`
	Dist      uint      `form:"dist"`
	Fmin      float32   `form:"fmin"`
	Fmax      float32   `form:"fmax"`
	Chroma    uint8     `form:"chroma"`
	Diff      uint32    `form:"diff"`
	Mono      bool      `form:"mono"`
	Portrait  bool      `form:"portrait"`
	Location  bool      `form:"location"`
	Album     string    `form:"album"`
	Label     string    `form:"label"`
	Country   string    `form:"country"`
	Year      uint      `form:"year"`
	Month     uint      `form:"month"`
	Color     string    `form:"color"`
	Quality   int       `form:"quality"`
	Review    bool      `form:"review"`
	Camera    int       `form:"camera"`
	Lens      int       `form:"lens"`
	Before    time.Time `form:"before" time_format:"2006-01-02"`
	After     time.Time `form:"after" time_format:"2006-01-02"`
	Favorites bool      `form:"favorites"`
	Public    bool      `form:"public"`
	Private   bool      `form:"private"`
	Story     bool      `form:"story"`
	Safe      bool      `form:"safe"`
	Nsfw      bool      `form:"nsfw"`
	Count     int       `form:"count" binding:"required"`
	Offset    int       `form:"offset"`
	Order     string    `form:"order"`
	Merged    bool      `form:"merged"`
}

func (f *PhotoSearch) GetQuery() string {
	return f.Query
}

func (f *PhotoSearch) SetQuery(q string) {
	f.Query = q
}

func (f *PhotoSearch) ParseQueryString() error {
	return ParseQueryString(f)
}

func NewPhotoSearch(query string) PhotoSearch {
	return PhotoSearch{Query: query}
}
