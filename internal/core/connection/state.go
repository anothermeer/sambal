package connection

type State int

const (
	StateWaiting State = iota
	StateConnecting
	StateNegotiating
	StateUploading
	StateVerifying
	StateCompleted
	StateCancelled
	StateTimedOut
	StateFailed
)
