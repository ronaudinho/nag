package nag_test

import (
	"fmt"
	"testing"

	"github.com/ronaudinho/nag"
)

func TestScoreBoardV2_Get(t *testing.T) {
	sbv2 := nag.NewScoreBoardV2()
	err := sbv2.Get()
	if err != nil {
		t.Error(err)
	}
	if sbv2.Response == nil {
		t.Error("nil response")
	}

	// this is just to check if response mapping is as expected
	// or if we need to update the SDK
	// probably can use this `go generate` relevant response struct
	expect := map[string][]string{
		"Available":              []string{"GAME_ID", "PT_AVAILABLE"},
		"EastConfStandingsByDay": []string{"TEAM_ID", "LEAGUE_ID", "SEASON_ID", "STANDINGSDATE", "CONFERENCE", "TEAM", "G", "W", "L", "W_PCT", "HOME_RECORD", "ROAD_RECORD"},
		"GameHeader":             []string{"GAME_DATE_EST", "GAME_SEQUENCE", "GAME_ID", "GAME_STATUS_ID", "GAME_STATUS_TEXT", "GAMECODE", "HOME_TEAM_ID", "VISITOR_TEAM_ID", "SEASON", "LIVE_PERIOD", "LIVE_PC_TIME", "NATL_TV_BROADCASTER_ABBREVIATION", "HOME_TV_BROADCASTER_ABBREVIATION", "AWAY_TV_BROADCASTER_ABBREVIATION", "LIVE_PERIOD_TIME_BCAST", "ARENA_NAME", "WH_STATUS"},
		"LastMeeting":            []string{"GAME_ID", "LAST_GAME_ID", "LAST_GAME_DATE_EST", "LAST_GAME_HOME_TEAM_ID", "LAST_GAME_HOME_TEAM_CITY", "LAST_GAME_HOME_TEAM_NAME", "LAST_GAME_HOME_TEAM_ABBREVIATION", "LAST_GAME_HOME_TEAM_POINTS", "LAST_GAME_VISITOR_TEAM_ID", "LAST_GAME_VISITOR_TEAM_CITY", "LAST_GAME_VISITOR_TEAM_NAME", "LAST_GAME_VISITOR_TEAM_CITY1", "LAST_GAME_VISITOR_TEAM_POINTS"},
		"LineScore":              []string{"GAME_DATE_EST", "GAME_SEQUENCE", "GAME_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_CITY_NAME", "TEAM_NAME", "TEAM_WINS_LOSSES", "PTS_QTR1", "PTS_QTR2", "PTS_QTR3", "PTS_QTR4", "PTS_OT1", "PTS_OT2", "PTS_OT3", "PTS_OT4", "PTS_OT5", "PTS_OT6", "PTS_OT7", "PTS_OT8", "PTS_OT9", "PTS_OT10", "PTS", "FG_PCT", "FT_PCT", "FG3_PCT", "AST", "REB", "TOV"},
		"SeriesStandings":        []string{"GAME_ID", "HOME_TEAM_ID", "VISITOR_TEAM_ID", "GAME_DATE_EST", "HOME_TEAM_WINS", "HOME_TEAM_LOSSES", "SERIES_LEADER"},
		"TeamLeaders":            []string{"GAME_ID", "TEAM_ID", "TEAM_CITY", "TEAM_NICKNAME", "TEAM_ABBREVIATION", "PTS_PLAYER_ID", "PTS_PLAYER_NAME", "PTS", "REB_PLAYER_ID", "REB_PLAYER_NAME", "REB", "AST_PLAYER_ID", "AST_PLAYER_NAME", "AST"},
		"TicketLinks":            []string{"GAME_ID", "LEAG_TIX"},
		"WestConfStandingsByDay": []string{"TEAM_ID", "LEAGUE_ID", "SEASON_ID", "STANDINGSDATE", "CONFERENCE", "TEAM", "G", "W", "L", "W_PCT", "HOME_RECORD", "ROAD_RECORD"},
		"WinProbability":         []string{},
	}
	m := nag.GetMap(*sbv2.Response)
	if !isExpected(m, expect) {
		t.Error("unexpected datasets")
	}
}

func isExpected(m map[string]interface{}, expect map[string][]string) bool {
	for k, headers := range expect {
		vv, ok := m[k]
		if !ok {
			fmt.Printf("%s does not exist\n", k)
			return false
		}

		rows, _ := vv.([]map[string]interface{})
		if len(rows) == 0 {
			continue
		}

		row := rows[0]
		for _, h := range headers {
			if _, ok := row[h]; !ok {
				fmt.Printf("%s.%s does not exist\n", k, h)
				return false
			}
		}
	}
	return true
}
