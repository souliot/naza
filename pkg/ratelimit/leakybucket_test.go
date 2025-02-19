// Copyright 2020, Chef.  All rights reserved.
// https://github.com/souliot/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package ratelimit_test

import (
	"testing"
	"time"

	"github.com/souliot/naza/pkg/log"

	"github.com/souliot/naza/pkg/assert"
	"github.com/souliot/naza/pkg/ratelimit"
)

func TestNewLeakyBucket(t *testing.T) {
	lb := ratelimit.NewLeakyBucket(10)
	log.DefaultBeeLogger.Debug("MaybeAvailableIntervalMSec=%d", lb.MaybeAvailableIntervalMSec())
}

func TestLeakyBucket_TryAquire(t *testing.T) {
	var (
		lb  *ratelimit.LeakyBucket
		err error
	)

	lb = ratelimit.NewLeakyBucket(1)
	time.Sleep(10 * time.Millisecond)
	err = lb.TryAquire()
	log.DefaultBeeLogger.Debug("MaybeAvailableIntervalMSec=%d", lb.MaybeAvailableIntervalMSec())
	assert.Equal(t, nil, err)
	time.Sleep(10 * time.Millisecond)
	err = lb.TryAquire()
	assert.Equal(t, nil, err)
	log.DefaultBeeLogger.Debug("MaybeAvailableIntervalMSec=%d", lb.MaybeAvailableIntervalMSec())

	lb = ratelimit.NewLeakyBucket(100)
	err = lb.TryAquire()
	assert.Equal(t, ratelimit.ErrResourceNotAvailable, err)
	log.DefaultBeeLogger.Debug("MaybeAvailableIntervalMSec=%d", lb.MaybeAvailableIntervalMSec())
	err = lb.TryAquire()
	assert.Equal(t, ratelimit.ErrResourceNotAvailable, err)
	log.DefaultBeeLogger.Debug("MaybeAvailableIntervalMSec=%d", lb.MaybeAvailableIntervalMSec())
}

func TestLeakyBucket_WaitUntilAquire(t *testing.T) {
	var lb *ratelimit.LeakyBucket

	lb = ratelimit.NewLeakyBucket(1)
	lb.WaitUntilAquire()
	log.DefaultBeeLogger.Debug("MaybeAvailableIntervalMSec=%d", lb.MaybeAvailableIntervalMSec())
	time.Sleep(100 * time.Millisecond)
	lb.WaitUntilAquire()
	log.DefaultBeeLogger.Debug("MaybeAvailableIntervalMSec=%d", lb.MaybeAvailableIntervalMSec())

	lb = ratelimit.NewLeakyBucket(200)
	lb.WaitUntilAquire()
	log.DefaultBeeLogger.Debug("MaybeAvailableIntervalMSec=%d", lb.MaybeAvailableIntervalMSec())
	lb.WaitUntilAquire()
	log.DefaultBeeLogger.Debug("MaybeAvailableIntervalMSec=%d", lb.MaybeAvailableIntervalMSec())
}
