package main

import "encoding/json"
import "fmt"
import "io/ioutil"
import "net/http"
import "net/url"
import "strings"


type OAuth2Provider struct {
    auth_endpoint, client_id, client_secret, redirect_uri string
}

type OAuth2Credentials struct {
    Access_token, Token_type, Scope string
    Expires_in int
}

func tokenHasScope(token OAuth2Credentials, required_scope string) bool {
    return strings.Index(token.Scope, required_scope) != -1
}

func ValidateUserToken(authorization_code string) (bool, string) {
    
    if authorization_code == "" {
        return false, "No authorization_code provided"
    }
        
    application_information := OAuth2Provider{
        auth_endpoint: "http://127.0.0.1:8000/oauth2/token/",
        redirect_uri: "http://127.0.0.1:8080/oauth2/callback",
        client_id: "",
    }
    
    grant_type := "authorization_code"
    
    stuff := url.Values{
        "client_id": {application_information.client_id}, 
        "grant_type": {grant_type}, 
        "code": {authorization_code},
        "redirect_uri": {application_information.redirect_uri},
    }

    // Getting a token
    resp, err := http.PostForm(application_information.auth_endpoint,
        stuff,
    )
    
    if err != nil {
        // @ToDo: Log instead of print
        fmt.Printf("err: %v\n", err)
    } else {
        // Connection was OK, let's process it
        defer resp.Body.Close()
        body, body_err := ioutil.ReadAll(resp.Body)
        
        if body_err != nil || strings.Index(string(body), "error") != -1 {
            // @ToDo: Log instead of print
            fmt.Printf("body_err: %v\n", body_err)
            fmt.Printf("body: %v\n", string(body))
            return false, string(body)
        } else {
            var tokeninfo OAuth2Credentials
            json_err := json.Unmarshal([]uint8(body), &tokeninfo)
            
            if json_err != nil {
                // @ToDo: Log instead of print
                fmt.Printf("Json unmarshal err: %v\n", json_err)
                return false, fmt.Sprintf("JSON parse error: %v", json_err)
            } else {
                // Parsing is OK, verify token scope and return
                has_scope := tokenHasScope(tokeninfo, "read")
                return has_scope, ""
            }
        }
    }
    return false, "Something went wrong"
}
