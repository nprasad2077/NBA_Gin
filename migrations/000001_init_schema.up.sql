-- 1. Create the Schema
CREATE SCHEMA IF NOT EXISTS analytics_dev_intermediate;

-- 2. Create Games Table
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
    total_points BIGINT,
    is_overtime BOOLEAN,
    home_offensive_rating NUMERIC,
    home_defensive_rating NUMERIC,
    home_net_rating NUMERIC,
    home_pace NUMERIC,
    home_effective_fg_pct NUMERIC,
    home_turnover_rate NUMERIC,
    home_offensive_tier TEXT,
    home_defensive_tier TEXT,
    visitor_offensive_rating NUMERIC,
    visitor_defensive_rating NUMERIC,
    visitor_net_rating NUMERIC,
    visitor_pace NUMERIC,
    visitor_effective_fg_pct NUMERIC,
    visitor_turnover_rate NUMERIC,
    visitor_offensive_tier TEXT,
    visitor_defensive_tier TEXT,
    matchup_pace NUMERIC
);

-- 3. Create Player Performance Table
CREATE TABLE IF NOT EXISTS analytics_dev_intermediate.int_player_performance (
    game_id TEXT,
    player_id TEXT,
    player_name TEXT,
    team TEXT,
    game_date TIMESTAMPTZ,
    season_start_year NUMERIC,
    is_playoff BOOLEAN,
    game_result TEXT,
    team_location TEXT,
    minutes_played NUMERIC,
    points BIGINT,
    assists BIGINT,
    total_rebounds BIGINT,
    steals BIGINT,
    blocks BIGINT,
    turnovers BIGINT,
    plus_minus BIGINT,
    net_rating BIGINT,
    box_plus_minus NUMERIC,
    field_goals_made BIGINT,
    field_goals_attempted BIGINT,
    field_goal_pct NUMERIC,
    three_pointers_made BIGINT,
    three_pointers_attempted BIGINT,
    three_point_pct NUMERIC,
    true_shooting_pct NUMERIC,
    effective_fg_pct NUMERIC,
    usage_pct NUMERIC,
    offensive_rating BIGINT,
    defensive_rating BIGINT,
    usage_tier TEXT,
    impact_tier TEXT,
    shooting_efficiency_tier TEXT,
    minutes_based_role TEXT,
    is_double_double BOOLEAN,
    is_triple_double BOOLEAN,
    is_versatile BOOLEAN,
    is_defensive_specialist BOOLEAN,
    is_three_and_d BOOLEAN,
    -- Composite Primary Key to ensure uniqueness
    PRIMARY KEY (game_id, player_id)
);

-- 4. Create Team Performance Table
CREATE TABLE IF NOT EXISTS analytics_dev_intermediate.int_team_performance (
    game_id TEXT,
    team TEXT,
    game_date TIMESTAMPTZ,
    season_start_year NUMERIC,
    is_playoff BOOLEAN,
    game_result TEXT,
    opponent_team TEXT,
    points BIGINT,
    offensive_rating NUMERIC,
    defensive_rating NUMERIC,
    net_rating NUMERIC,
    pace NUMERIC,
    effective_fg_pct NUMERIC,
    turnover_rate NUMERIC,
    offensive_rebound_rate NUMERIC,
    free_throw_rate NUMERIC,
    offensive_tier TEXT,
    defensive_tier TEXT,
    shot_selection_style TEXT,
    ball_movement_style TEXT,
    ball_security_tier TEXT,
    defensive_activity TEXT,
    -- Composite Primary Key
    PRIMARY KEY (game_id, team)
);

-- 5. Performance Indexes
-- Games
CREATE INDEX IF NOT EXISTS idx_games_date ON analytics_dev_intermediate.int_games_enriched(game_date);
CREATE INDEX IF NOT EXISTS idx_games_season ON analytics_dev_intermediate.int_games_enriched(season_start_year);
CREATE INDEX IF NOT EXISTS idx_games_hometeam ON analytics_dev_intermediate.int_games_enriched(home_team);
CREATE INDEX IF NOT EXISTS idx_games_visitorteam ON analytics_dev_intermediate.int_games_enriched(visitor_team);

-- Player Performance
CREATE INDEX IF NOT EXISTS idx_player_lookup ON analytics_dev_intermediate.int_player_performance(player_id);
CREATE INDEX IF NOT EXISTS idx_player_game_date ON analytics_dev_intermediate.int_player_performance(game_date);
CREATE INDEX IF NOT EXISTS idx_player_team ON analytics_dev_intermediate.int_player_performance(team);

-- Team Performance
CREATE INDEX IF NOT EXISTS idx_team_perf_lookup ON analytics_dev_intermediate.int_team_performance(team);
CREATE INDEX IF NOT EXISTS idx_team_perf_date ON analytics_dev_intermediate.int_team_performance(game_date);