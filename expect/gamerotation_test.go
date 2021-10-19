package expect_test

import (
	"testing"

	"github.com/ronaudinho/nag"
	"github.com/ronaudinho/nag/expect"
)

func TestGameRotation_Get(t *testing.T) {
	sbv2 := nag.NewGameRotation(testGameID)
	err := sbv2.Get()
	if err != nil {
		t.Error(err)
	}
	if sbv2.Response == nil {
		t.Error("nil response")
	}

	m := nag.Map(*sbv2.Response)
	if !isExpected(m, expect.GameRotation) {
		t.Error("unexpected datasets")
	}
}
