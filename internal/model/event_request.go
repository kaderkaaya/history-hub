package model

type GetEventsRequest struct {
	Month int    `form:"month" binding:"required,gte=1,lte=31"`
	Day   int    `form:"day" binding:"required,gte=1,lte=31"`
	Typ   string `form:"type" binding:"omitempty,oneof=events births deaths holidays selected all"`
	Lang  string `form:"language" binding:"omitempty,oneof=tr en"`
}

type GetTodayRequest struct {
	Typ  string `form:"type" binding:"omitempty,oneof=events births deaths holidays selected all"`
	Lang string `form:"language" binding:"omitempty,oneof=tr en"`
}
