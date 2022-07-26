package ressources

import "gebes.io/sticker_backend/pkg/ent/sticker"

type Sticker struct {
	LocationDescription *string          `json:"location_description" post:"required"`
	Latitude            *float64         `json:"latitude" post:"required"`
	Longitude           *float64         `json:"longitude" post:"required"`
	Edition             *sticker.Edition `json:"edition" post:"required"`
}
