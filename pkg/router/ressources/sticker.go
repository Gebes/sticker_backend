package ressources

type Sticker struct {
	LocationDescription *string  `json:"location_description" post:"required"`
	Latitude            *float64 `json:"latitude" post:"required"`
	Longitude           *float64 `json:"longitude" post:"required"`
}
