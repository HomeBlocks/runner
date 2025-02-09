package runner

type Storage interface {
	Cancel()
	IsCancelled() bool
}

type DefaultStorage struct {
	cancelled bool
}

func (d *DefaultStorage) Cancel() {
	d.cancelled = true
}

func (d *DefaultStorage) IsCancelled() bool {
	return d.cancelled
}
