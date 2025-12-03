package domain

import (
	"context"
	"time"
)

type Game struct {
	// Core Info
	GameID          string     `json:"game_id" db:"game_id"`
	GameDate        *time.Time `json:"game_date" db:"game_date"`
	SeasonStartYear *float64   `json:"season_start_year" db:"season_start_year"`
	IsPlayoff       *bool      `json:"is_playoff" db:"is_playoff"`
	Arena           *string    `json:"arena" db:"arena"`
	ArenaCity       *string    `json:"arena_city" db:"arena_city"`

	// Teams & Result
	HomeTeam          *string `json:"home_team" db:"home_team"`
	VisitorTeam       *string `json:"visitor_team" db:"visitor_team"`
	WinningTeam       *string `json:"winning_team" db:"winning_team"`
	HomePoints        *int64  `json:"home_points" db:"home_points"`
	VisitorPoints     *int64  `json:"visitor_points" db:"visitor_points"`
	PointDifferential *int64  `json:"point_differential" db:"point_differential"`
	TotalPoints       *int64  `json:"total_points" db:"total_points"`
	IsOvertime        *bool   `json:"is_overtime" db:"is_overtime"`

	// Home Advanced Stats
	HomeOffensiveRating *float64 `json:"home_offensive_rating" db:"home_offensive_rating"`
	HomeDefensiveRating *float64 `json:"home_defensive_rating" db:"home_defensive_rating"`
	HomeNetRating       *float64 `json:"home_net_rating" db:"home_net_rating"`
	HomePace            *float64 `json:"home_pace" db:"home_pace"`
	HomeEffectiveFGPct  *float64 `json:"home_effective_fg_pct" db:"home_effective_fg_pct"`
	HomeTurnoverRate    *float64 `json:"home_turnover_rate" db:"home_turnover_rate"`
	HomeOffensiveTier   *string  `json:"home_offensive_tier" db:"home_offensive_tier"`
	HomeDefensiveTier   *string  `json:"home_defensive_tier" db:"home_defensive_tier"`

	// Visitor Advanced Stats
	VisitorOffensiveRating *float64 `json:"visitor_offensive_rating" db:"visitor_offensive_rating"`
	VisitorDefensiveRating *float64 `json:"visitor_defensive_rating" db:"visitor_defensive_rating"`
	VisitorNetRating       *float64 `json:"visitor_net_rating" db:"visitor_net_rating"`
	VisitorPace            *float64 `json:"visitor_pace" db:"visitor_pace"`
	VisitorEffectiveFGPct  *float64 `json:"visitor_effective_fg_pct" db:"visitor_effective_fg_pct"`
	VisitorTurnoverRate    *float64 `json:"visitor_turnover_rate" db:"visitor_turnover_rate"`
	VisitorOffensiveTier   *string  `json:"visitor_offensive_tier" db:"visitor_offensive_tier"`
	VisitorDefensiveTier   *string  `json:"visitor_defensive_tier" db:"visitor_defensive_tier"`

	// Game Flow
	MatchupPace *float64 `json:"matchup_pace" db:"matchup_pace"`
}

// Repository Interface
type GameRepository interface {
	GetGameByID(ctx context.Context, id string) (*Game, error)
	// Updated Signature
	ListGames(ctx context.Context, filter GameFilter) (*GamesResponse, error)
}