package tasks

import "time"

func Task25(d time.Duration) {
	<-time.After(d)
}
