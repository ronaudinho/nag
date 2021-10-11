package expect_test

import (
	"testing"

	"github.com/ronaudinho/nag"
	"github.com/ronaudinho/nag/expect"
)

func TestAllTimeLeadersGrids_Get(t *testing.T) {
	sbv2 := nag.NewAllTimeLeadersGrids()
	err := sbv2.Get()
	if err != nil {
		t.Error(err)
	}
	if sbv2.Response == nil {
		t.Error("nil response")
	}

	m := nag.GetMap(*sbv2.Response)
	if !isExpected(m, expect.AllTimeLeadersGrids) {
		t.Error("unexpected datasets")
	}
}
