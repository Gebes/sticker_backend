package router

import (
	. "github.com/Gebes/there/v2"
)

func DiscordLogin(request HttpRequest) HttpResponse {
	return loginFlow(request, discordConfig)
}

func DiscordAuthCallback(request HttpRequest) HttpResponse {
	return callbackFlow(request, discordConfig)
}
