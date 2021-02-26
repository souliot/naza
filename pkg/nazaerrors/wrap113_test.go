// Copyright 2021, Chef.  All rights reserved.
// https://github.com/souliot/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

// +build go1.13

package nazaerrors

import (
	"errors"
	"io"
	"testing"

	"github.com/souliot/naza/pkg/assert"

	"github.com/souliot/naza/pkg/log"
)

func TestWrap(t *testing.T) {
	err := Wrap(io.EOF)
	log.DefaultBeeLogger.Debug("%+v", err)
	assert.Equal(t, true, errors.Is(err, io.EOF))
	err = Wrap(err)
	log.DefaultBeeLogger.Debug("%+v", err)
	assert.Equal(t, true, errors.Is(err, io.EOF))
}
