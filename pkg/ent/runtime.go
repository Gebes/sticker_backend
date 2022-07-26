// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"gebes.io/sticker_backend/pkg/ent/schema"
	"gebes.io/sticker_backend/pkg/ent/sticker"
	"gebes.io/sticker_backend/pkg/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	stickerFields := schema.Sticker{}.Fields()
	_ = stickerFields
	// stickerDescLocationDescription is the schema descriptor for location_description field.
	stickerDescLocationDescription := stickerFields[0].Descriptor()
	// sticker.LocationDescriptionValidator is a validator for the "location_description" field. It is called by the builders before save.
	sticker.LocationDescriptionValidator = stickerDescLocationDescription.Validators[0].(func(string) error)
	// stickerDescCreatedAt is the schema descriptor for created_at field.
	stickerDescCreatedAt := stickerFields[3].Descriptor()
	// sticker.DefaultCreatedAt holds the default value on creation for the created_at field.
	sticker.DefaultCreatedAt = stickerDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[2].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[3].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}