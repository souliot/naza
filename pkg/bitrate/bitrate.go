// Copyright 2019, Chef.  All rights reserved.
// https://github.com/souliot/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package bitrate

import (
	"sync"
	"time"
)

type Bitrate interface {
	// @param nowUnixMSec: 变参，可选择从外部传入当前 unix 时间戳，单位毫秒
	Add(bytes int, nowUnixMSec ...int64)

	Rate(nowUnixMSec ...int64) float32
}

type Unit uint8

const (
	UnitBitPerSec Unit = iota + 1
	UnitBytePerSec
	UnitKBitPerSec
	UnitKBytePerSec
)

// TODO chef: 考虑支持配置是否在内部使用锁
type Option struct {
	WindowMS int
	Unit     Unit
}

var defaultOption = Option{
	WindowMS: 1000,
	Unit:     UnitKBitPerSec,
}

type ModOption func(option *Option)

func New(modOptions ...ModOption) Bitrate {
	option := defaultOption
	for _, fn := range modOptions {
		fn(&option)
	}
	return &bitrate{
		option: option,
	}
}

type bitrate struct {
	option Option

	mu          sync.Mutex
	bucketSlice []bucket
}

type bucket struct {
	n int
	t int64 // unix 时间戳，单位毫秒
}

func (b *bitrate) Add(bytes int, nowUnixMSec ...int64) {
	var now int64
	if len(nowUnixMSec) == 0 {
		now = time.Now().UnixNano() / 1e6
	} else {
		now = nowUnixMSec[0]
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	b.sweepStale(now)
	b.bucketSlice = append(b.bucketSlice, bucket{
		n: bytes,
		t: now,
	})
}

func (b *bitrate) Rate(nowUnixMSec ...int64) float32 {
	var now int64
	if len(nowUnixMSec) == 0 {
		now = time.Now().UnixNano() / 1e6
	} else {
		now = nowUnixMSec[0]
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	b.sweepStale(now)
	var total int
	for i := range b.bucketSlice {
		total += b.bucketSlice[i].n
	}

	var ret float32
	switch b.option.Unit {
	case UnitBitPerSec:
		ret = float32(total*8*1000) / float32(b.option.WindowMS)
	case UnitBytePerSec:
		ret = float32(total*1000) / float32(b.option.WindowMS)
	case UnitKBitPerSec:
		ret = float32(total*8) / float32(b.option.WindowMS)
	case UnitKBytePerSec:
		ret = float32(total) / float32(b.option.WindowMS)
	}
	return ret
}

func (b *bitrate) sweepStale(now int64) {
	i := 0
	l := len(b.bucketSlice)
	for ; i < l; i++ {
		if now-b.bucketSlice[i].t <= int64(b.option.WindowMS) {
			break
		}
	}
	if i == l {
		b.bucketSlice = nil
	} else {
		b.bucketSlice = b.bucketSlice[i:]
	}
}
