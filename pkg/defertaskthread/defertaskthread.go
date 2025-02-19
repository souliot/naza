// Copyright 2020, Chef.  All rights reserved.
// https://github.com/souliot/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package defertaskthread

import "time"

type deferTaskThread struct {
}

func (d *deferTaskThread) Go(deferMS int, task TaskFn, param ...interface{}) {
	go func() {
		time.Sleep(time.Duration(deferMS) * time.Millisecond)
		task(param...)
	}()
}
