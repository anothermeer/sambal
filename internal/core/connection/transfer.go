package connection

import (
	"context"
	"time"
)

type Transfer struct {
	ID string

	FileName string
	Size     int64

	DeviceID   string
	DeviceName string

	State State

	Progress int64

	CreatedAt time.Time

	ctx    context.Context
	cancel context.CancelFunc
}
