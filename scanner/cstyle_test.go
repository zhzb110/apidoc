// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package scanner

import (
	"testing"

	"github.com/issue9/assert"
)

var _ scanFunc = cstyle

var code1 = `
int x = 5;
/* line1
line2
line3*/
`
var comment1 = []byte(` line1
line2
line3`)

func TestCStyle(t *testing.T) {
	a := assert.New(t)

	s := &scanner{
		data: []byte(code1),
	}
	block, err := cstyle(s)
	a.NotError(err).NotNil(block)
	a.Equal(block, comment1)
}