package router

import (
	"gebes.io/sticker_backend/pkg/ent"
	. "github.com/Gebes/there/v2"
)

func UserGet(request HttpRequest) HttpResponse {
	requester, ok := request.Context().Value("user").(*ent.User)
	if !ok {
		return Error(StatusUnprocessableEntity, "could not get user from context")
	}
	return Json(StatusOK, requester)
}
