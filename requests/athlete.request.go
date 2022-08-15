package requests

type GetAthletesRequest struct {
	limit int `json:"limit" validate:"min=1,max=50"`
	page  int `json:"page" validate:"min=1"`
}
