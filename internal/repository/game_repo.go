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
		SELECT *
		FROM analytics_dev_intermediate.int_games_enriched
		WHERE game_id = @id
	`
	args := pgx.NamedArgs{"id": id}

	row, _ := r.db.Pool.Query(ctx, query, args)
	
	// FIX: Use RowToAddrOfStructByName to return a pointer (*domain.Game)
	return pgx.CollectOneRow(row, pgx.RowToAddrOfStructByName[domain.Game])
}

func (r *gameRepository) ListGames(ctx context.Context, limit, offset int) ([]domain.Game, error) {
	query := `
		SELECT *
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
	defer rows.Close()

	// This remains RowToStructByName because ListGames returns []domain.Game (slice of values)
	return pgx.CollectRows(rows, pgx.RowToStructByName[domain.Game])
}