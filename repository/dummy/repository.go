package dummy

import (
	honeypot "github.com/0n1shi/redis-honeypot"
)

type DummyRepository struct {
}

var _ honeypot.Repository = (*DummyRepository)(nil)

func NewDummyRepository() honeypot.Repository {
	return &DummyRepository{}
}

func (r *DummyRepository) Save(cmd *honeypot.Command) error {
	return nil
}
