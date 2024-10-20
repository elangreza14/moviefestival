package dto

type (
	MostViewedGenreListResponseElement struct {
		Name  string `json:"name"`
		Views int    `json:"views"`
	}

	GenreListResponse []MostViewedGenreListResponseElement
)
