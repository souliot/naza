// Copyright 2020, Chef.  All rights reserved.
// https://github.com/souliot/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazahttp_test

import (
	"bufio"
	"fmt"
	"net"
	"testing"

	"github.com/souliot/naza/pkg/assert"

	"github.com/souliot/naza/pkg/log"

	"github.com/souliot/naza/pkg/nazahttp"
)

func TestHeader(t *testing.T) {
	for port := 8080; port != 8090; port++ {
		addr := fmt.Sprintf(":%d", port)
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			continue
		}

		go func() {
			_, _ = nazahttp.GetHTTPFile(fmt.Sprintf("http://%s/test", addr), 100)
		}()

		conn, err := ln.Accept()
		r := bufio.NewReader(conn)
		fl, hs, err := nazahttp.ReadHTTPHeader(r)
		assert.Equal(t, nil, err)
		assert.Equal(t, true, len(hs) > 0)
		log.DefaultBeeLogger.Debug("first line:%s", fl)
		log.DefaultBeeLogger.Debug("header fields:%+v", hs)

		m, u, v, err := nazahttp.ParseHTTPRequestLine(fl)
		assert.Equal(t, nil, err)
		log.DefaultBeeLogger.Debug("method:%s, uri:%s, version:%s", m, u, v)
		assert.Equal(t, "GET", m)
		assert.Equal(t, "/test", u)
		assert.Equal(t, "HTTP/1.1", v)
		break
	}
}

func TestParseHTTPStatusLine(t *testing.T) {
	v, c, r, e := nazahttp.ParseHTTPStatusLine("HTTP/1.0 200 OK")
	assert.Equal(t, nil, e)
	assert.Equal(t, "HTTP/1.0", v)
	assert.Equal(t, "200", c)
	assert.Equal(t, "OK", r)

	v, c, r, e = nazahttp.ParseHTTPStatusLine("HTTP/1.1 400 Bad Request")
	assert.Equal(t, nil, e)
	assert.Equal(t, "HTTP/1.1", v)
	assert.Equal(t, "400", c)
	assert.Equal(t, "Bad Request", r)

	statusLine := "HTTP/1.1 400 "
	for i := 0; i <= len(statusLine); i++ {
		sl := statusLine[0:i]
		_, _, _, e = nazahttp.ParseHTTPStatusLine(sl)
		assert.IsNotNil(t, e, sl)
	}
}
