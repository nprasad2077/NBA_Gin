package domain

import (
	"context"
	"time"
)

type PlayerPerformance struct {
	// Identifiers
	GameID          string     `json:"game_id" db:"game_id"`
	PlayerID        string     `json:"player_id" db:"player_id"`
	PlayerName      *string    `json:"player_name" db:"player_name"`
	Team            *string    `json:"team" db:"team"`
	GameDate        *time.Time `json:"game_date" db:"game_date"`
	SeasonStartYear *float64   `json:"season_start_year" db:"season_start_year"`
	IsPlayoff       *bool      `json:"is_playoff" db:"is_playoff"`
	GameResult      *string    `json:"game_result" db:"game_result"`
	TeamLocation    *string    `json:"team_location" db:"team_location"`

	// Basic Stats
	MinutesPlayed *float64 `json:"minutes_played" db:"minutes_played"`
	Points        *int64   `json:"points" db:"points"`
	Assists       *int64   `json:"assists" db:"assists"`
	TotalRebounds *int64   `json:"total_rebounds" db:"total_rebounds"`
	Steals        *int64   `json:"steals" db:"steals"`
	Blocks        *int64   `json:"blocks" db:"blocks"`
	Turnovers     *int64   `json:"turnovers" db:"turnovers"`

	// Advanced & Impact
	PlusMinus    *int64   `json:"plus_minus" db:"plus_minus"`
	NetRating    *int64   `json:"net_rating" db:"net_rating"`
	BoxPlusMinus *float64 `json:"box_plus_minus" db:"box_plus_minus"`

	// Shooting
	FieldGoalsMade          *int64   `json:"field_goals_made" db:"field_goals_made"`
	FieldGoalsAttempted     *int64   `json:"field_goals_attempted" db:"field_goals_attempted"`
	FieldGoalPct            *float64 `json:"field_goal_pct" db:"field_goal_pct"`
	ThreePointersMade       *int64   `json:"three_pointers_made" db:"three_pointers_made"`
	ThreePointersAttempted  *int64   `json:"three_pointers_attempted" db:"three_pointers_attempted"`
	ThreePointPct           *float64 `json:"three_point_pct" db:"three_point_pct"`
	TrueShootingPct         *float64 `json:"true_shooting_pct" db:"true_shooting_pct"`
	EffectiveFGPct          *float64 `json:"effective_fg_pct" db:"effective_fg_pct"`

	// Usage & Ratings
	UsagePct        *float64 `json:"usage_pct" db:"usage_pct"`
	OffensiveRating *int64   `json:"offensive_rating" db:"offensive_rating"`
	DefensiveRating *int64   `json:"defensive_rating" db:"defensive_rating"`

	// Tiers & Roles
	UsageTier              *string `json:"usage_tier" db:"usage_tier"`
	ImpactTier             *string `json:"impact_tier" db:"impact_tier"`
	ShootingEfficiencyTier *string `json:"shooting_efficiency_tier" db:"shooting_efficiency_tier"`
	MinutesBasedRole       *string `json:"minutes_based_role" db:"minutes_based_role"`

	// Flags
	IsDoubleDouble        *bool `json:"is_double_double" db:"is_double_double"`
	IsTripleDouble        *bool `json:"is_triple_double" db:"is_triple_double"`
	IsVersatile           *bool `json:"is_versatile" db:"is_versatile"`
	IsDefensiveSpecialist *bool `json:"is_defensive_specialist" db:"is_defensive_specialist"`
	IsThreeAndD           *bool `json:"is_three_and_d" db:"is_three_and_d"`
}

type PlayerRepository interface {
	GetPlayerPerformance(ctx context.Context, gameID, playerID string) (*PlayerPerformance, error)
	ListPlayerPerformances(ctx context.Context, gameID string) ([]PlayerPerformance, error)
}