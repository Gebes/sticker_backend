package router

import . "github.com/Gebes/there/v2"

func AuthMiddleware(request HttpRequest, next HttpResponse) HttpResponse {
	_, errResponse := oauthUserInfo(request, discordConfig)
	if errResponse != nil {
		return Error(StatusUnauthorized, "not logged in to discord")
	}
	return next
}
