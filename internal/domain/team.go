package domain

import (
	"context"
	"time"
)

type TeamPerformance struct {
	GameID          string     `json:"game_id" db:"game_id"`
	Team            string     `json:"team" db:"team"`
	GameDate        *time.Time `json:"game_date" db:"game_date"`
	SeasonStartYear *float64   `json:"season_start_year" db:"season_start_year"`
	IsPlayoff       *bool      `json:"is_playoff" db:"is_playoff"`
	GameResult      *string    `json:"game_result" db:"game_result"`
	OpponentTeam    *string    `json:"opponent_team" db:"opponent_team"`
	Points          *int64     `json:"points" db:"points"`

	// Ratings
	OffensiveRating *float64 `json:"offensive_rating" db:"offensive_rating"`
	DefensiveRating *float64 `json:"defensive_rating" db:"defensive_rating"`
	NetRating       *float64 `json:"net_rating" db:"net_rating"`
	Pace            *float64 `json:"pace" db:"pace"`

	// Four Factors & Style
	EffectiveFGPct       *float64 `json:"effective_fg_pct" db:"effective_fg_pct"`
	TurnoverRate         *float64 `json:"turnover_rate" db:"turnover_rate"`
	OffensiveReboundRate *float64 `json:"offensive_rebound_rate" db:"offensive_rebound_rate"`
	FreeThrowRate        *float64 `json:"free_throw_rate" db:"free_throw_rate"`

	// Tiers & Style
	OffensiveTier       *string `json:"offensive_tier" db:"offensive_tier"`
	DefensiveTier       *string `json:"defensive_tier" db:"defensive_tier"`
	ShotSelectionStyle  *string `json:"shot_selection_style" db:"shot_selection_style"`
	BallMovementStyle   *string `json:"ball_movement_style" db:"ball_movement_style"`
	BallSecurityTier    *string `json:"ball_security_tier" db:"ball_security_tier"`
	DefensiveActivity   *string `json:"defensive_activity" db:"defensive_activity"`
}

type TeamRepository interface {
	GetTeamPerformance(ctx context.Context, gameID, team string) (*TeamPerformance, error)
}