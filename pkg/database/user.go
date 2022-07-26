package database

import (
	"context"
	"gebes.io/sticker_backend/pkg/ent"
	"gebes.io/sticker_backend/pkg/ent/user"
)

func UserById(id string) (*ent.User, error) {
	return Client.User.Query().Where(user.ID(id)).Only(context.Background())
}

func UserCreate() *ent.UserCreate {
	return Client.User.Create()
}
