package router

import (
	"context"
	"fmt"
	"gebes.io/sticker_backend/pkg/database"
	"gebes.io/sticker_backend/pkg/router/response"
	"gebes.io/sticker_backend/pkg/router/ressources"
	. "github.com/Gebes/there/v2"
	"github.com/pkg/errors"
)

func StickerGet(request HttpRequest) HttpResponse {
	stickers, err := database.Stickers()
	if err != nil {
		return Error(StatusInternalServerError, fmt.Errorf("unable to query stickers: %v", err))
	}
	return response.Gzip(Json(StatusOK, stickers))
}

func StickerPost(request HttpRequest) HttpResponse {
	var body ressources.Sticker
	err := request.Body.BindJson(&body)
	if err != nil {
		return Error(StatusBadRequest, err)
	}
	err = postValidator.Struct(body)
	if err != nil {
		return Error(StatusBadRequest, err)
	}

	stickerCreated, err := database.StickerCreate().
		SetLocationDescription(*body.LocationDescription).
		SetLatitude(*body.Latitude).
		SetLongitude(*body.Longitude).
		Save(context.Background())

	if err != nil {
		return Error(StatusInternalServerError, errors.Wrap(err, "could not create user in the database"))
	}
	return Json(StatusCreated, stickerCreated)
}
