package dto

type (
	MovieListResponseElement struct {
		ID          int      `json:"id"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Duration    string   `json:"duration"`
		WatchUrl    string   `json:"watch_url"`
		Views       int      `json:"views"`
		Artists     []string `json:"artists"`
		Genres      []string `json:"genres"`
	}

	MovieListResponse []MovieListResponseElement

	CreateMoviePayload struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Duration    string   `json:"duration"`
		Artists     []string `json:"artists"`
		Genres      []string `json:"genres"`
		WatchUrl    string   `json:"watch_url"`
	}
)
