package expect

var BoxScoreSummaryV2 = map[string][]string{
	"AvailableVideo":  []string{"GAME_ID", "VIDEO_AVAILABLE_FLAG", "PT_AVAILABLE", "PT_XYZ_AVAILABLE", "WH_STATUS", "HUSTLE_STATUS", "HISTORICAL_STATUS"},
	"GameInfo":        []string{"GAME_DATE", "ATTENDANCE", "GAME_TIME"},
	"GameSummary":     []string{"GAME_DATE_EST", "GAME_SEQUENCE", "GAME_ID", "GAME_STATUS_ID", "GAME_STATUS_TEXT", "GAMECODE", "HOME_TEAM_ID", "VISITOR_TEAM_ID", "SEASON", "LIVE_PERIOD", "LIVE_PC_TIME", "NATL_TV_BROADCASTER_ABBREVIATION", "LIVE_PERIOD_TIME_BCAST", "WH_STATUS"},
	"InactivePlayers": []string{"PLAYER_ID", "FIRST_NAME", "LAST_NAME", "JERSEY_NUM", "TEAM_ID", "TEAM_CITY", "TEAM_NAME", "TEAM_ABBREVIATION"},
	"LastMeeting":     []string{"GAME_ID", "LAST_GAME_ID", "LAST_GAME_DATE_EST", "LAST_GAME_HOME_TEAM_ID", "LAST_GAME_HOME_TEAM_CITY", "LAST_GAME_HOME_TEAM_NAME", "LAST_GAME_HOME_TEAM_ABBREVIATION", "LAST_GAME_HOME_TEAM_POINTS", "LAST_GAME_VISITOR_TEAM_ID", "LAST_GAME_VISITOR_TEAM_CITY", "LAST_GAME_VISITOR_TEAM_NAME", "LAST_GAME_VISITOR_TEAM_CITY1", "LAST_GAME_VISITOR_TEAM_POINTS"},
	"LineScore":       []string{"GAME_DATE_EST", "GAME_SEQUENCE", "GAME_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_CITY_NAME", "TEAM_NICKNAME", "TEAM_WINS_LOSSES", "PTS_QTR1", "PTS_QTR2", "PTS_QTR3", "PTS_QTR4", "PTS_OT1", "PTS_OT2", "PTS_OT3", "PTS_OT4", "PTS_OT5", "PTS_OT6", "PTS_OT7", "PTS_OT8", "PTS_OT9", "PTS_OT10", "PTS"},
	"Officials":       []string{"OFFICIAL_ID", "FIRST_NAME", "LAST_NAME", "JERSEY_NUM"},
	"OtherStats":      []string{"LEAGUE_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_CITY", "PTS_PAINT", "PTS_2ND_CHANCE", "PTS_FB", "LARGEST_LEAD", "LEAD_CHANGES", "TIMES_TIED", "TEAM_TURNOVERS", "TOTAL_TURNOVERS", "TEAM_REBOUNDS", "PTS_OFF_TO"},
	"SeasonSeries":    []string{"GAME_ID", "HOME_TEAM_ID", "VISITOR_TEAM_ID", "GAME_DATE_EST", "HOME_TEAM_WINS", "HOME_TEAM_LOSSES", "SERIES_LEADER"},
}
