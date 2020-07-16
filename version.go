// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package api

import (
	"time"
)

func Version() string {
	// This format mirrors what is in the makefile
	return time.Now().In(time.UTC).Format("2006.01.02")
}
