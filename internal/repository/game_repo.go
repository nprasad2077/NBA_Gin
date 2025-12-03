package repository

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/nprasad2077/NBA_Gin/internal/domain"

	"github.com/jackc/pgx/v5"
)

// Update Interface
func NewGameRepository(db *DB) domain.GameRepository {
	return &gameRepository{db: db}
}

type gameRepository struct {
	db *DB
}

// Update Interface definition in domain/game.go first, or just match this signature:
// ListGames(ctx context.Context, filter domain.GameFilter) (*domain.GamesResponse, error)

func (r *gameRepository) ListGames(ctx context.Context, filter domain.GameFilter) (*domain.GamesResponse, error) {
	// 1. Base Query
	baseQuery := `SELECT * FROM analytics_dev_intermediate.int_games_enriched WHERE 1=1`
	countQuery := `SELECT COUNT(*) FROM analytics_dev_intermediate.int_games_enriched WHERE 1=1`
	
	// 2. Dynamic Filters
	args := pgx.NamedArgs{}
	var conditions []string

	if filter.Team != "" {
		// Checks if team is Home OR Visitor
		conditions = append(conditions, "(home_team = @team OR visitor_team = @team)")
		args["team"] = filter.Team
	}

	if filter.Season != nil {
		conditions = append(conditions, "season_start_year = @season")
		args["season"] = *filter.Season
	}

	if filter.IsPlayoff != nil {
		conditions = append(conditions, "is_playoff = @is_playoff")
		args["is_playoff"] = *filter.IsPlayoff
	}

	if filter.GameDateStart != nil {
		conditions = append(conditions, "game_date >= @game_date_start")
		args["game_date_start"] = *filter.GameDateStart
	}
    
	if filter.GameDateEnd != nil {
		// Postgres compares timestamps efficiently. 
        // passing "2025-11-30" defaults to midnight, so we typically might want 
        // to handle inclusive end dates by adding 24h, but for now simple <= is fine.
		conditions = append(conditions, "game_date <= @game_date_end")
		args["game_date_end"] = *filter.GameDateEnd
	}

	// 3. Assemble Query
	if len(conditions) > 0 {
		joinCond := " AND " + strings.Join(conditions, " AND ")
		baseQuery += joinCond
		countQuery += joinCond
	}

	// 4. Get Total Count (For Pagination Meta)
	var totalItems int64
	row := r.db.Pool.QueryRow(ctx, countQuery, args)
	if err := row.Scan(&totalItems); err != nil {
		return nil, fmt.Errorf("counting rows: %w", err)
	}

	// 5. Apply Sorting & Pagination
	// Sanitize SortBy to prevent SQL injection
	sortColumn := "game_date"
	sortDir := "DESC"
	
	if filter.SortBy != "" {
		if strings.HasPrefix(filter.SortBy, "-") {
			sortDir = "DESC"
			sortColumn = strings.TrimPrefix(filter.SortBy, "-")
		} else {
			sortDir = "ASC"
			sortColumn = filter.SortBy
		}
        // TODO: Validate sortColumn against a whitelist of allowed fields here
	}
    
    // Add Ordering and Limit
	baseQuery += fmt.Sprintf(" ORDER BY %s %s LIMIT @limit OFFSET @offset", sortColumn, sortDir)
	
	offset := (filter.Page - 1) * filter.PageSize
	args["limit"] = filter.PageSize
	args["offset"] = offset

	// 6. Execute Final Query
	rows, err := r.db.Pool.Query(ctx, baseQuery, args)
	if err != nil {
		return nil, fmt.Errorf("querying games: %w", err)
	}
	defer rows.Close()

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Game])
	if err != nil {
		return nil, err
	}

	// 7. Calculate Meta
	totalPages := int(totalItems) / filter.PageSize
	if int(totalItems)%filter.PageSize > 0 {
		totalPages++
	}

	return &domain.GamesResponse{
		Data: data,
		Meta: domain.Meta{
			CurrentPage: filter.Page,
			PageSize:    filter.PageSize,
			TotalItems:  totalItems,
			TotalPages:  totalPages,
		},
	}, nil
}

func (r *gameRepository) GetGameByID(ctx context.Context, id string) (*domain.Game, error) {
	// SELECT * works now because the domain.Game struct has fields for ALL DB columns
	query := `
		SELECT * FROM analytics_dev_intermediate.int_games_enriched 
		WHERE game_id = @id
	`
	args := pgx.NamedArgs{"id": id}

	row, _ := r.db.Pool.Query(ctx, query, args)
	
	// Use RowToAddrOfStructByName to return a pointer (*Game)
	return pgx.CollectOneRow(row, pgx.RowToAddrOfStructByName[domain.Game])
}