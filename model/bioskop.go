package model

type Bioskop struct {
	ID     *string  `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Nama   *string  `json:"nama"`
	Lokasi *string  `json:"lokasi"`
	Rating *float64 `json:"rating"`
}
