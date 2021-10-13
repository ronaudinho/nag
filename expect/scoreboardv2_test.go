package expect_test

import (
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
	m := nag.Map(*sbv2.Response)
	if !isExpected(m, expect.ScoreBoardV2) {
		t.Error("unexpected datasets")
	}
}
