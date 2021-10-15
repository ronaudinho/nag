package expect_test

import (
	"testing"

	"github.com/ronaudinho/nag"
	"github.com/ronaudinho/nag/expect"
)

func TestCommonTeamRoster_Get(t *testing.T) {
	sbv2 := nag.NewCommonTeamRoster(testTeamID)
	err := sbv2.Get()
	if err != nil {
		t.Error(err)
	}
	if sbv2.Response == nil {
		t.Error("nil response")
	}

	m := nag.Map(*sbv2.Response)
	if !isExpected(m, expect.CommonTeamRoster) {
		t.Error("unexpected datasets")
	}
}
