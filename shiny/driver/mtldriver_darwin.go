// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin && metal

package driver

import (
	"github.com/nathan-joslin/exp/shiny/driver/mtldriver"
	"github.com/nathan-joslin/exp/shiny/screen"
)

func main(f func(screen.Screen)) {
	mtldriver.Main(f)
}
