package runner

type Storage interface {
	Close()
	IsClosed() bool
}

type DefaultStorage struct {
	cancelled bool
}

func (d *DefaultStorage) Close() {
	d.cancelled = true
}

func (d *DefaultStorage) IsClosed() bool {
	return d.cancelled
}
