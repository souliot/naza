// Copyright 2019, Chef.  All rights reserved.
// https://github.com/souliot/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package fake_test

import (
	"testing"

	"github.com/souliot/naza/pkg/assert"
	"github.com/souliot/naza/pkg/fake"
)

func TestWithFakeExit(t *testing.T) {
	var er fake.ExitResult
	er = fake.WithFakeOSExit(func() {
		fake.OS_Exit(1)
	})
	assert.Equal(t, true, er.HasExit)
	assert.Equal(t, 1, er.ExitCode)

	er = fake.WithFakeOSExit(func() {
	})
	assert.Equal(t, false, er.HasExit)

	er = fake.WithFakeOSExit(func() {
		fake.OS_Exit(2)
	})
	assert.Equal(t, true, er.HasExit)
	assert.Equal(t, 2, er.ExitCode)
}
