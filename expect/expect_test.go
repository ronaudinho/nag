package expect_test

import (
	"fmt"
)

const (
	testGameID   = "0012100038" // using existing game id
	testPlayerID = "201937"     // rubio
	testTeamID   = "1610612748" // heat
)

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
