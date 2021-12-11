package expect_test

import (
	"testing"

	"github.com/ronaudinho/nag"
	"github.com/ronaudinho/nag/expect"
)

func TestTeamGameLog_Get(t *testing.T) {
	sbv2 := nag.NewTeamGameLog(testTeamID)
	err := sbv2.Get()
	if err != nil {
		t.Error(err)
	}
	if sbv2.Response == nil {
		t.Error("nil response")
	}

	m := nag.Map(*sbv2.Response)
	if !isExpected(m, expect.TeamGameLog) {
		t.Error("unexpected datasets")
	}
}
