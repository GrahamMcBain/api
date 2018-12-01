// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"net/http"

	moov "github.com/moov-io/go-client/client"

	"github.com/antihax/optional"
)

// setMoovCookie adds authentication onto our Moov API client for all requests
func setMoovCookie(conf *moov.Configuration, cookie *http.Cookie) {
	if cookie.Value != "" {
		conf.AddDefaultHeader("Cookie", fmt.Sprintf("moov_auth=%s", cookie.Value))
	}
}

// verifyUserIsLoggedIn takes the given moov.APIClient and checks if it's is logged in. A non-nil error signals
// the client doens't have valid authentication.
func verifyUserIsLoggedIn(ctx context.Context, api *moov.APIClient, user *user, requestId string) error {
	resp, err := api.UserApi.CheckUserLogin(ctx, &moov.CheckUserLoginOpts{
		XRequestId: optional.NewString(requestId),
	})
	if err != nil {
		return fmt.Errorf("problem checking user (id=%s) login: %v", user.ID, err)
	}
	if resp != nil {
		if resp.StatusCode != 200 {
			return fmt.Errorf("on cookie check, got %s status", resp.Status)
		}
		return resp.Body.Close()
	}
	return nil
}