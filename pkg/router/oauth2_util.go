package router

import (
	"context"
	"encoding/base64"
	"encoding/json"
	. "github.com/Gebes/there/v2"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type AccessTokenCookie string

const (
	AccessTokenCookieDiscord AccessTokenCookie = "access_token_discord"
)

type StateCookie string

const (
	StateCookieDiscord StateCookie = "oauthstatediscord"
)

// Generate random cookie to prevent CSRF attacks	res, err := s.twitterConfig.Client(context.Background(), token).Get("https://api.twitter.com/2/users/me")
func generateStateCookie(name StateCookie) (*http.Cookie, error) {
	var expires = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{
		Name:    string(name),
		Value:   state,
		Expires: expires,
		Secure:  true,
		Path:    "/",
	}
	return &cookie, nil
}

func tokenFromCookie(accessTokenCookie *http.Cookie) *oauth2.Token {
	return &oauth2.Token{
		AccessToken: accessTokenCookie.Value,
		Expiry:      accessTokenCookie.Expires,
	}
}

type CallbackFlowConfiguration struct {
	LoginRoute              func(request HttpRequest) HttpResponse
	OAuth2Config            *oauth2.Config
	CurrentUserUrl          string
	RedirectAuthCodeOptions []oauth2.AuthCodeOption
	UserInfoAuthCodeOptions []oauth2.AuthCodeOption
	DestinationUrl          string

	StateCookie       StateCookie
	AccessTokenCookie AccessTokenCookie
}

func loginFlow(request HttpRequest, config *CallbackFlowConfiguration) HttpResponse {
	cookie, err := generateStateCookie(config.StateCookie)
	if err != nil {
		return Error(StatusInternalServerError, "Could not generate auth cookie: "+err.Error())
	}
	http.SetCookie(request.ResponseWriter, cookie)
	return Redirect(StatusTemporaryRedirect, config.OAuth2Config.AuthCodeURL(cookie.Value, config.RedirectAuthCodeOptions...))
}

func callbackFlow(request HttpRequest, config *CallbackFlowConfiguration) HttpResponse {
	// check if the user stored an access token cookie
	var token *oauth2.Token
	accessTokenCookie, err := request.Request.Cookie(string(config.AccessTokenCookie))

	if err != nil { // seems like no, so lets get one
		oauthState, err := request.Request.Cookie(string(config.StateCookie))
		if err != nil {
			return Error(StatusBadRequest, err)
		}

		if request.Request.FormValue("state") != oauthState.Value {
			// If the state is invalid, the user should be redirected to the login screen
			// return Redirect(StatusTemporaryRedirect, "https://brokkr.finance")
			// currently there is no login screen, so get a new refresh token
			return config.LoginRoute(request)
		}

		token, err = config.OAuth2Config.Exchange(context.Background(), request.Request.FormValue("code"), config.UserInfoAuthCodeOptions...)

		if err != nil {
			return Error(StatusBadRequest, err)
		}
	}

	if token == nil {
		token = tokenFromCookie(accessTokenCookie)
	}

	_, errResponse := oauthUserInfo(request, config, token)
	if errResponse != nil {
		return errResponse
	}

	http.SetCookie(request.ResponseWriter, &http.Cookie{
		Name:    string(config.AccessTokenCookie),
		Value:   token.AccessToken,
		Expires: token.Expiry,
		Secure:  true,
		Path:    "/",
	})
	return Redirect(StatusTemporaryRedirect, config.DestinationUrl+"/")
}

// oauthUserInfo returns the user info as a map
// if the user is not logged in (no access token cookie provided)
// then the user will be redirected to the according login url
func oauthUserInfo(request HttpRequest, config *CallbackFlowConfiguration, optionalToken ...*oauth2.Token) (Map, HttpResponse) {

	var token *oauth2.Token
	if len(optionalToken) == 0 {

		accessTokenCookie, err := request.Request.Cookie(string(config.AccessTokenCookie))

		if err != nil {
			r := config.LoginRoute(request)
			return nil, r
		}

		token = tokenFromCookie(accessTokenCookie)
	} else {
		token = optionalToken[0]
	}

	res, err := config.OAuth2Config.Client(context.Background(), token).Get(config.CurrentUserUrl)

	if err != nil {
		return nil, Error(StatusBadRequest, err)
	}

	// let us check if the token is valid
	if res.StatusCode == StatusUnauthorized {
		// the token expired
		// at this point, the token should only be expired,
		// if the user deleted the oauth2 app from his account
		http.SetCookie(request.ResponseWriter, &http.Cookie{
			Name:    string(config.AccessTokenCookie),
			Expires: time.Unix(0, 0),
			Secure:  true,
			Path:    "/",
		})
		// reduce response time, by calling the other route,
		// instead of redirecting the user
		return nil, config.LoginRoute(request)
	} else if res.StatusCode != StatusOK {
		return nil, Error(res.StatusCode, res.Status)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, Error(StatusBadRequest, errors.Wrap(err, "could not read body from "+config.CurrentUserUrl))
	}
	var data Map
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, Error(StatusBadRequest, errors.Wrap(err, "could not read body from "+config.CurrentUserUrl))
	}
	return data, nil
}
