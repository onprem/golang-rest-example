package store

import (
	"context"

	"github.com/onprem/go-db-example/ent"
)

func (s *Store) GetUsers(ctx context.Context) ([]*ent.User, error) {
	return s.db.User.Query().All(ctx)
}
