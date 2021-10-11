package nag

import (
	"fmt"
	"strconv"
	"time"
)

func GetSeason(t time.Time) string {
	cur := t.Year()
	if t.Month() <= 9 {
		cur = cur - 1
	}
	nxt := strconv.Itoa(cur + 1)[2:]
	return fmt.Sprintf("%d-%s", cur, nxt)
}
