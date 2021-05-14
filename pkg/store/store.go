package store

import "github.com/onprem/go-db-example/ent"

type Store struct {
	db *ent.Client
}

func New(client *ent.Client) *Store {
	return &Store{
		db: client,
	}
}
