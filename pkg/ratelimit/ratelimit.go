// Copyright 2020, Chef.  All rights reserved.
// https://github.com/souliot/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package ratelimit

type RateLimiter interface {
	TryAquire() error
	WaitUntilAquire()
}
