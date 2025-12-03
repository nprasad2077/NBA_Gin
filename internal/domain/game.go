package domain

import (
    "context"
    "time"
)

type Game struct {
	GameID              string     `json:"game_id" db:"game_id"`
	GameDate            *time.Time `json:"game_date" db:"game_date"`
	SeasonStartYear     *float64   `json:"season_start_year" db:"season_start_year"`
	IsPlayoff           *bool      `json:"is_playoff" db:"is_playoff"`
	Arena               *string    `json:"arena" db:"arena"`
	ArenaCity           *string    `json:"arena_city" db:"arena_city"`
	HomeTeam            *string    `json:"home_team" db:"home_team"`
	VisitorTeam         *string    `json:"visitor_team" db:"visitor_team"`
	WinningTeam         *string    `json:"winning_team" db:"winning_team"`
	HomePoints          *int64     `json:"home_points" db:"home_points"`
	VisitorPoints       *int64     `json:"visitor_points" db:"visitor_points"`
	HomeNetRating       *float64   `json:"home_net_rating" db:"home_net_rating"`
}

// Repository Interface
type GameRepository interface {
	GetGameByID(ctx context.Context, id string) (*Game, error)
	ListGames(ctx context.Context, limit, offset int) ([]Game, error)
}