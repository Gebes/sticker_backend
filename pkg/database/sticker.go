package database

import (
	"context"
	"gebes.io/sticker_backend/pkg/ent"
	"gebes.io/sticker_backend/pkg/ent/sticker"
)

func Stickers() ([]*ent.Sticker, error) {
	return Client.Sticker.Query().WithOwner().All(context.Background())
}

func StickerById(id int) (*ent.Sticker, error) {
	return Client.Sticker.Query().Where(sticker.ID(id)).WithOwner().Only(context.Background())
}

func StickerCreate() *ent.StickerCreate {
	return Client.Sticker.Create()
}

func StickerDelete(id int) error {
	return Client.Sticker.DeleteOneID(id).Exec(context.Background())
}
