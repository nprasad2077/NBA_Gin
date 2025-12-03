package repository

import (
	"context"
	"fmt"
	"github.com/nprasad2077/NBA_Gin/internal/domain"
	
	"github.com/jackc/pgx/v5"
)

type gameRepository struct {
	db *DB
}

func NewGameRepository(db *DB) domain.GameRepository {
	return &gameRepository{db: db}
}

func (r *gameRepository) GetGameByID(ctx context.Context, id string) (*domain.Game, error) {
	query := `
		SELECT game_id, game_date, home_team, visitor_team, home_points, visitor_points
		FROM analytics_dev_intermediate.int_games_enriched
		WHERE game_id = @id
	`
	args := pgx.NamedArgs{"id": id}

	// 1. Execute Query and handle error
	rows, err := r.db.Pool.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	// 2. Collect the row into a value
	game, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Game])
	if err != nil {
		return nil, err
	}

	// 3. Return the pointer to the value
	return &game, nil
}

func (r *gameRepository) ListGames(ctx context.Context, limit, offset int) ([]domain.Game, error) {
	query := `
		SELECT game_id, game_date, home_team, visitor_team, home_points, visitor_points
		FROM analytics_dev_intermediate.int_games_enriched
		ORDER BY game_date DESC
		LIMIT @limit OFFSET @offset
	`
	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}

	rows, err := r.db.Pool.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	// No defer rows.Close() needed here because CollectRows closes them, 
	// but keeping it doesn't hurt. pgx.CollectRows handles the closing usually.
	
	return pgx.CollectRows(rows, pgx.RowToStructByName[domain.Game])
}