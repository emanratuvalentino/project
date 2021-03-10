package api

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// GetJWTPayload is an example function for get JWT Payload. Uncomment for use.
func GetJWTPayload(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	/*--------- Get the name payload from JWT, which is a string. Level 1 dimension ---------*/
	// name := claims["name"].(string)
	// return c.String(http.StatusOK, name)

	/*--------- Get the resource_access payload from JWT, which is an array. Level 1 dimension ---------*/
	// resourceAccess := claims["resource_access"].(map[string]interface{})
	// return c.JSON(http.StatusOK, resourceAccess)

	/*--------- Check the user roles payload from JWT, which is an array. Level 2 dimension ---------*/
	roles := claims["realm_access"].(map[string]interface{})["roles"]

	chkRoles := roles.([]interface{})
	for _, role := range chkRoles {
		if role == "Admin" {
			fmt.Println("Check Roles passed")
			break
		}
	}
	return c.JSON(http.StatusOK, roles)
}
