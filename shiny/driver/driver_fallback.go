// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !darwin && (!linux || android) && !windows && !dragonfly && !openbsd

package driver

import (
	"errors"

	"github.com/nathan-joslin/exp/shiny/driver/internal/errscreen"
	"github.com/nathan-joslin/exp/shiny/screen"
)

func main(f func(screen.Screen)) {
	f(errscreen.Stub(errors.New("no driver for accessing a screen")))
}
