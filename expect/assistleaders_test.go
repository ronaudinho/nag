package expect_test

import (
	"testing"

	"github.com/ronaudinho/nag"
	"github.com/ronaudinho/nag/expect"
)

func TestAssistLeaders_Get(t *testing.T) {
	sbv2 := nag.NewAssistLeaders()
	err := sbv2.Get()
	if err != nil {
		t.Error(err)
	}
	if sbv2.Response == nil {
		t.Error("nil response")
	}

	m := nag.GetMap(*sbv2.Response)
	if !isExpected(m, expect.AssistLeaders) {
		t.Error("unexpected datasets")
	}
}
