package router

import (
	"context"
	"gebes.io/sticker_backend/pkg/auth"
	"gebes.io/sticker_backend/pkg/env"
	"github.com/Gebes/there/v2"
	"github.com/Gebes/there/v2/middlewares"
	"github.com/go-playground/validator/v10"
	"golang.org/x/oauth2"
)

var (
	postValidator  = validator.New()
	patchValidator = validator.New()
)
var discordConfig = &CallbackFlowConfiguration{
	OAuth2Config: &oauth2.Config{
		RedirectURL:  env.DiscordClientRedirectUrl,
		ClientID:     env.DiscordClientId,
		ClientSecret: env.DiscordClientSecret,
		Scopes:       []string{auth.DiscordScopeIdentify},
		Endpoint:     auth.DiscordEndpoint,
	},
	CurrentUserUrl:          "https://discord.com/api/users/@me",
	UserInfoAuthCodeOptions: []oauth2.AuthCodeOption{},
	RedirectAuthCodeOptions: []oauth2.AuthCodeOption{},
	StateCookie:             StateCookieDiscord,
	AccessTokenCookie:       AccessTokenCookieDiscord,
	DestinationUrl:          env.DestinationUrl,
}

var router = there.NewRouter()

func Listen() error {
	discordConfig.LoginRoute = DiscordLogin

	postValidator.SetTagName("post")
	patchValidator.SetTagName("patch")

	router.Use(middlewares.Cors(middlewares.AllowAllConfiguration()))

	router.Group("/auth/discord").
		Get("/login", DiscordLogin).
		Get("/callback", DiscordAuthCallback)

	router.Group("/sticker").
		Get("/", StickerGet).With(AuthMiddleware).
		Post("/", StickerPost).With(AuthMiddleware)

	return router.Listen(8080)
}

func RequesterIp(request there.HttpRequest) string {
	return request.Headers.GetDefault("X-Real-Ip", "127.0.0.1")
}

func Shutdown() error {
	return router.Server.Shutdown(context.Background())
}
