CREATE SCHEMA IF NOT EXISTS analytics_dev_intermediate;

CREATE TABLE IF NOT EXISTS analytics_dev_intermediate.int_games_enriched (
    game_id TEXT PRIMARY KEY,
    game_date TIMESTAMPTZ,
    season_start_year NUMERIC,
    is_playoff BOOLEAN,
    arena TEXT,
    arena_city TEXT,
    home_team TEXT,
    visitor_team TEXT,
    winning_team TEXT,
    home_points BIGINT,
    visitor_points BIGINT,
    point_differential BIGINT,
    home_offensive_rating NUMERIC,
    home_net_rating NUMERIC
);
-- Add indexes for performance
CREATE INDEX idx_games_date ON analytics_dev_intermediate.int_games_enriched(game_date);
CREATE INDEX idx_games_home_team ON analytics_dev_intermediate.int_games_enriched(home_team);