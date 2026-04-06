package model

type GetEventsRequest struct {
	Month int    `json:"month" binding:"required,gte=1,lte=31"`
	Day   int    `json:"day" binding:"required,gte=1,lte=31"`
	Typ   string `json:"type" binding:"omitempty,oneof=events births deaths holidays selected all"`
	Lang  string `json:"language" binding:"omitempty,oneof=tr en"`
}
