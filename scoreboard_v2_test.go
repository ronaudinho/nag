package nag_test

import (
	"fmt"
	"testing"

	"github.com/ronaudinho/nag"
	"github.com/ronaudinho/nag/expect"
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
	m := nag.GetMap(*sbv2.Response)
	if !isExpected(m, expect.ScoreBoardV2) {
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
