package store

import (
	"context"

	"github.com/onprem/go-db-example/ent"
)

func (s *Store) GetUsers(ctx context.Context) ([]*ent.User, error) {
	return s.db.User.Query().All(ctx)
}

func (s *Store) CreateUser(ctx context.Context, name string, age int) (*ent.User, error) {
	return s.db.User.Create().
		SetName(name).
		SetAge(age).
		Save(ctx)
}
