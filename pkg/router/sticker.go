package router

import (
	"context"
	"fmt"
	"gebes.io/sticker_backend/pkg/database"
	"gebes.io/sticker_backend/pkg/ent"
	"gebes.io/sticker_backend/pkg/router/response"
	"gebes.io/sticker_backend/pkg/router/ressources"
	. "github.com/Gebes/there/v2"
	"github.com/pkg/errors"
	"strconv"
)

func StickerGet(request HttpRequest) HttpResponse {
	stickers, err := database.Stickers()
	if err != nil {
		return Error(StatusInternalServerError, fmt.Errorf("unable to query stickers: %v", err))
	}
	return response.Gzip(Json(StatusOK, stickers))
}

func StickerPost(request HttpRequest) HttpResponse {
	requester, ok := request.Context().Value("user").(*ent.User)
	if !ok {
		return Error(StatusUnprocessableEntity, "could not get user from context")
	}
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
		SetEdition(*body.Edition).
		SetOwner(requester).
		Save(context.Background())

	if err != nil {
		return Error(StatusInternalServerError, errors.Wrap(err, "could not create user in the database"))
	}
	return Json(StatusCreated, stickerCreated)
}

func StickerDelete(request HttpRequest) HttpResponse {
	requester, ok := request.Context().Value("user").(*ent.User)
	if !ok {
		return Error(StatusUnprocessableEntity, "could not get user from context")
	}
	idString := request.Params.GetDefault("id", "")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return Error(StatusBadRequest, fmt.Errorf("invalid id: %v", err))
	}

	sticker, err := database.StickerById(id)
	if ent.IsNotFound(err) {
		return Error(StatusNotFound, "sticker not found")
	} else if err != nil {
		return Error(StatusInternalServerError, fmt.Errorf("unable to get sticker: %v", err))
	}
	if sticker.Edges.Owner.ID != requester.ID {
		return Error(StatusForbidden, "not your sticker")
	}

	err = database.StickerDelete(sticker.ID)
	if err != nil {
		return Error(StatusInternalServerError, errors.Wrap(err, "could not delete sticker"))
	}
	return Json(StatusCreated, sticker)
}
