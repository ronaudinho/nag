package expect_test

import (
	"testing"

	"github.com/ronaudinho/nag"
	"github.com/ronaudinho/nag/expect"
)

func TestPlayByPlayV2_Get(t *testing.T) {
	sbv2 := nag.NewPlayByPlayV2("0012100038") // using existing game id
	err := sbv2.Get()
	if err != nil {
		t.Error(err)
	}
	if sbv2.Response == nil {
		t.Error("nil response")
	}

	m := nag.GetMap(*sbv2.Response)
	if !isExpected(m, expect.PlayByPlayV2) {
		t.Error("unexpected datasets")
	}
}
