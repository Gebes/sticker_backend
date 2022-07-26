package router

import (
	"context"
	"fmt"
	"gebes.io/sticker_backend/pkg/database"
	"gebes.io/sticker_backend/pkg/ent"
	. "github.com/Gebes/there/v2"
)

func AuthMiddleware(request HttpRequest, next HttpResponse) HttpResponse {
	discordUser, errResponse := oauthUserInfo(request, discordConfig)
	if errResponse != nil {
		return Error(StatusUnauthorized, "not logged in to discord")
	}

	id, ok := discordUser["id"].(string)
	if !ok {
		return Error(StatusInternalServerError, "unable to get id from discord user")
	}

	user, err := database.UserById(id)
	if ent.IsNotFound(err) {
		name, ok := discordUser["username"].(string)
		if !ok {
			return Error(StatusInternalServerError, "unable to get name from discord user")
		}

		user, err = database.UserCreate().
			SetID(id).
			SetName(name).
			Save(context.Background())
	} else if err != nil {
		return Error(StatusInternalServerError, fmt.Errorf("unable to get user: %v", err))
	}

	request.WithContext(context.WithValue(request.Context(), "user", user))

	return next
}
