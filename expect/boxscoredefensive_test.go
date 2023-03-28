package expect_test

import (
	"testing"

	"github.com/ronaudinho/nag"
	"github.com/ronaudinho/nag/expect"
)

func TestBoxScoreDefensive_Get(t *testing.T) {
	t.Skip("TODO needs update")
	sbv2 := nag.NewBoxScoreDefensive(testGameID)
	err := sbv2.Get()
	if err != nil {
		t.Error(err)
	}
	if sbv2.Response == nil {
		t.Error("nil response")
	}

	m := nag.Map(*sbv2.Response)
	if !isExpected(m, expect.BoxScoreDefensive) {
		t.Error("unexpected datasets")
	}
}
