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

func (s *Store) DeleteUser(ctx context.Context, id int) error {
	return s.db.User.DeleteOneID(id).Exec(ctx)
}

func (s *Store) UpdateUser(ctx context.Context, id int, name *string, age *int) (*ent.User, error) {
	upd := s.db.User.UpdateOneID(id)

	if name != nil {
		upd = upd.SetName(*name)
	}
	if age != nil {
		upd = upd.SetAge(*age)
	}

	return upd.Save(ctx)
}
