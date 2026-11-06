package connection

import (
	"context"
	"errors"
	"time"
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
	t.cancel()
	t.State = StateCancelled
}

func (m *Manager) Run(t *Transfer, fn func(context.Context) error) error {
	var err error

	for i := 0; i <= m.Retry; i++ {
		if t.ctx.Err() != nil {
			t.State = StateCancelled
			return t.ctx.Err()
		}

		attemptCtx, cancel := context.WithTimeout(t.ctx, m.Timeout)
		err = fn(attemptCtx)
		cancel()

		if err == nil {
			t.State = StateCompleted
			t.cancel()
			return nil
		}

		if errors.Is(err, context.Canceled) {
			t.State = StateCancelled
			t.cancel()
			return err
		}

		if errors.Is(err, context.DeadlineExceeded) {
			t.State = StateTimedOut

			if i < m.Retry {
				continue
			}
			t.cancel()
			return err
		}

		if i < m.Retry {
			continue
		}

		t.State = StateFailed
		t.cancel()
		return err
	}

	return err
}
