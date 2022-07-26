package database

import (
	"context"
	"gebes.io/sticker_backend/pkg/ent"
)

func Stickers() ([]*ent.Sticker, error) {
	return Client.Sticker.Query().All(context.Background())
}

func StickerCreate() *ent.StickerCreate {
	return Client.Sticker.Create()
}
