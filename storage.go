package runner

// Storage defines methods for managing storage state.
type Storage interface {
	Close()
	IsClosed() bool
}

// DefaultStorage is a simple implementation of Storage using a cancelled flag.
type DefaultStorage struct {
	cancelled bool
}

// Close sets the cancelled flag to true, indicating the storage is closed.
func (d *DefaultStorage) Close() {
	d.cancelled = true
}

// IsClosed returns true if the storage is closed (cancelled).
func (d *DefaultStorage) IsClosed() bool {
	return d.cancelled
}
