package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type ClientData struct {
	ClientId string `json:"client_id" xml:"client_id"`
	Secret   string `json:"secret" xml:"secret"`
}

type Response struct {
	Message string      `json:"message"`
	Data    *ClientData `json:"data"`
}

func GetClientCredentials(c echo.Context) error {

	// configuration := config.GetConfig()

	// u := &ClientData{
	// 	ClientId: configuration.CLIENT_ID,
	// 	Secret:   configuration.CLIENT_SECRET,
	// }

	// t := &Response{
	// 	Message: "Client data extracted",
	// 	Data:    u,
	// }

	return c.String(http.StatusOK, "tes")

}

var oauth = oauth2.Config{
	ClientID:     "coba-login",
	ClientSecret: "a9645997-4e1e-47fa-aafb-bde08c3dfa81",
	RedirectURL:  "localhost:5000",
	Scopes:       []string{},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "http://iam.siasn.info/auth/realms/hello-world-authz/protocol/openid-connect/auth",
		TokenURL: "http://iam.siasn.info/auth/realms/hello-world-authz/protocol/openid-connect/token",
	},
}

func GetToken(c echo.Context) error {

	token, _ := oauth.Exchange(oauth2.NoContext, c.FormValue("code"))

	// fmt.Println(token.AccessToken)

	return c.JSON(http.StatusOK, token.AccessToken)
}
