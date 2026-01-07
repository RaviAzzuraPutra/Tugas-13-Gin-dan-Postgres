package request

type Bioskop_Request struct {
	Nama   *string  `json:"nama" binding:"required"`
	Lokasi *string  `json:"lokasi" binding:"required"`
	Rating *float64 `json:"rating"`
}
