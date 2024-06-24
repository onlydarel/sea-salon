package models

// Review is a model for the user review and comment
type Review struct {
	ReviewID int    `json:"review_id"`
	Name     string `json:"name"`
	Rating   string `json:"rating"`
	Comment  string `json:"comment"`
}

type Reservation struct {
	ReservationID int    `json:"reservation_id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	TypeOfService string `json:"type_of_service"`
	DateAndTime   string `json:"date"`
}
