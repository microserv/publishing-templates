package main

import "encoding/json"
import "fmt"
import "github.com/gin-gonic/gin"
import "io/ioutil"
import "net/http"
import "net/url"
import "strings"

type OAuth2Provider struct {
	auth_endpoint, client_id, client_secret, redirect_uri string
}

type OAuth2Credentials struct {
	Access_token, Token_type, Scope string
	Expires_in                      int
}

func tokenHasScope(token OAuth2Credentials, required_scope string) bool {
	return strings.Index(token.Scope, required_scope) != -1
}

var valid_tokens []OAuth2Credentials

func printValidTokens() {
	fmt.Println("Valid tokens:")

	for key, val := range valid_tokens {
		fmt.Printf("%v: %v\n", key, val)
	}
}

func IsTokenValid(token, scope string) bool {
	for _, val := range valid_tokens {
		if val.Access_token == token {
			return tokenHasScope(val, scope)
		}
	}
	return false
}

func ValidateUserToken(authorization_code string) (bool, string) {
	// Intentionally kept for easier debugging
	// printValidTokens()
	if authorization_code == "" {
		return false, "Missing authorization_code"
	}

	application_information := OAuth2Provider{
		auth_endpoint: "http://127.0.0.1:8000/oauth2/token/",
		redirect_uri:  "http://127.0.0.1:8080/oauth2/callback",
		client_id:     "49UCr6bEdK4cOIWWCESP3mviXD0onxJpvwoSg9Ao",
	}

	grant_type := "authorization_code"

	stuff := url.Values{
		"client_id":    {application_information.client_id},
		"grant_type":   {grant_type},
		"code":         {authorization_code},
		"redirect_uri": {application_information.redirect_uri},
	}

	// Getting a token
	resp, err := http.PostForm(application_information.auth_endpoint,
		stuff,
	)

	if err != nil {
		// @ToDo: Log instead of print
		fmt.Printf("err: %v\n", err)
		return false, fmt.Sprintf("Error: %v", err)
	} else {
		// Connection was OK, let's process it
		defer resp.Body.Close()
		body, body_err := ioutil.ReadAll(resp.Body)

		if body_err != nil || strings.Index(string(body), "error") != -1 {
			// @ToDo: Log instead of print
			fmt.Printf("body_err: %v\n", body_err)
			fmt.Printf("body: %v\n", string(body))
			return false, fmt.Sprintf("Error: %v\n%v", body_err, string(body))
		} else {
			var tokeninfo OAuth2Credentials
			json_err := json.Unmarshal([]uint8(body), &tokeninfo)
			valid_tokens = append(valid_tokens, tokeninfo)

			// Intentionally kept for easier debug
			// fmt.Printf("body: %v\n", string(body))

			if json_err != nil {
				// @ToDo: Log instead of print
				fmt.Printf("Json unmarshal err: %v\n", json_err)
				return false, fmt.Sprintf("JSON Error: %v", json_err)
			} else {
				// Parsing is OK, verify token scope and return
				return true, ""
			}
		}
	}
}

func getTokenFromHeader(r *http.Request) string {
	authorization_header := r.Header.Get("Authorization")
	if authorization_header == "" {
		return ""
	}

	header_split := strings.Split(authorization_header, " ")

	if len(header_split) != 2 {
		fmt.Println("Malformed Authorization header")
		return ""
	}

	token := header_split[1] // Header is 2nd item in Authorization header split by " "
	return token
}

func getToken(c *gin.Context) string {
	token_from_header := getTokenFromHeader(c.Request)
	if token_from_header != "" {
		return token_from_header
	}

	token_from_urlparams := c.DefaultQuery("token", "")
	if token_from_urlparams != "" {
		return token_from_urlparams
	}

	return ""
}

func ValidateRequest(scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getToken(c)

		if token == "" {
			c.AbortWithStatus(401)
			c.JSON(401, generateJSONErr(401, "Missing authentication credentials"))
			return
		}

		access_granted := IsTokenValid(token, scope)

		if !access_granted {
			c.AbortWithStatus(403)
			c.JSON(403, generateJSONErr(403, "Permission denied"))
			return
		}
		c.Header("WWW-Authenticate", "No authorization_code provided")
	}
}
