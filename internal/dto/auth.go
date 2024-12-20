package dto

type (
	RegisterPayload struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		Token       string `json:"token"`
		Permissions any    `json:"permissions"`
	}
)
