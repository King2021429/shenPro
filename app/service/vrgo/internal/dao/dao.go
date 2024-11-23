package dao

import (
	"context"
)

// dao dao.
type Dao struct {
}

func New() *Dao {
	d := &Dao{}
	return d
}

// Close close the resource.
func (d *Dao) Close() {
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return nil
}
