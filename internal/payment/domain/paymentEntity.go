package domain

type Payment struct {
	ID int64 `json:"id"`
	Username string `json:"user"`
	PlaceID int64 `json:"place_id"`
}
