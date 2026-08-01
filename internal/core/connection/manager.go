package connection

import (
	"errors"
	"time"

	"golang.org/x/net/context"
)

type Manager struct {
	Transfers map[string]*Transfer

	Timeout time.Duration
	Retry   int
}

func NewManager() *Manager {
	return &Manager{
		Transfers: make(map[string]*Transfer),
		Timeout:   30 * time.Second,
		Retry:     3,
	}
}

func (m *Manager) NewTransfer(id string) *Transfer {
	ctx, cancel := context.WithCancel(context.Background())

	t := &Transfer{
		ID:        id,
		State:     StateWaiting,
		CreatedAt: time.Now(),
		ctx:       ctx,
		cancel:    cancel,
	}

	m.Transfers[id] = t

	return t
}

func (t *Transfer) Cancel() {
	t.State = StateCancelled
	t.Cancel()
}

func (m *Manager) Run(t *Transfer, fn func(context.Context) error) error {
	var err error

	for i := 0; i <= m.Retry; i++ {
		err = fn(t.ctx)

		if err == nil {
			t.State = StateCompleted
			return nil
		}

		if errors.Is(err, context.Canceled) {
			t.State = StateCancelled
			return err
		}

		if errors.Is(err, context.DeadlineExceeded) {
			t.State = StateTimedOut
		}

		t.State = StateFailed
	}

	return err
}
