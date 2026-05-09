package network

import (
	"fmt"
)

type ProgressTracker struct {
	Total       int64
	Transferred int64
	LastPercent int
}

func (pt *ProgressTracker) Write(p []byte) (int, error) {
	n := len(p)

	pt.Transferred += int64(n)

	percent := int(float64(pt.Transferred) / float64(pt.Total) * 100)

	if percent != pt.LastPercent {
		fmt.Printf("\rUploading... %d%%", percent)
		pt.LastPercent = percent
	}

	return n, nil
}
