// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package httputil

import (
	"errors"
	"net/http"
	"strings"
)

const (
	userTokenPrefix = "Bearer "
)

var (
	errAuthSchema = "authentication scheme not starting with a bearer"
)

// ExtractAuthToken reads the value of request Authorization and removes the Bearer substring or returns error if it does not exist
func ExtractAuthToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")

	if !strings.HasPrefix(token, userTokenPrefix) {
		return token, errors.New(errAuthSchema)
	}

	return strings.TrimPrefix(token, userTokenPrefix), nil
}
