package domain

import "time"

type GameFilter struct {
	Team      string   `form:"team"`
	Season    *float64 `form:"season"`
	IsPlayoff *bool    `form:"is_playoff"`

	// CHANGED: Renamed to match the DB column 'game_date' explicitly
	GameDateStart *time.Time `form:"game_date_start" time_format:"2006-01-02"`
	GameDateEnd   *time.Time `form:"game_date_end" time_format:"2006-01-02"`

	// Pagination
	Page     int    `form:"page,default=1"`
	PageSize int    `form:"page_size,default=20"`
	SortBy   string `form:"sort_by"`
}

type Meta struct {
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
}

type GamesResponse struct {
	Data []Game `json:"data"`
	Meta Meta   `json:"meta"`
}