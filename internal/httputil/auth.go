// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package httputil

import (
	"errors"
	"net/http"
	"strings"
)

// FormatAuthString reads the value of request Authorization and removes the Bearer substring or returns error if it does not exist
func FormatAuthString(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if !strings.Contains(token, "Bearer ") {
		return token, errors.New("Token not containing a Bearer.")
	}

	return strings.ReplaceAll(token, "Bearer ", ""), nil
}
