package main

import (
	"context"
	"net"
	"net/http"
	"sync"
	"syscall/js"
	"time"

	promise "github.com/nlepage/go-js-promise"
	"github.com/shynome/wage/pkg/wshttp"
	"github.com/shynome/wahttp"
	"github.com/xtaci/smux"
	"nhooyr.io/websocket"
)

func main() {
	var ep = js.Global().Get("WageEndpoint").String()
	js.Global().Set("GoFetch", GoFetch(ep))
	<-make(chan any)
}

func GoFetch(endpoint string) js.Func {

	if endpoint == "" {
		panic("env Endpoint is not set")
	}

	var session *smux.Session
	var locker = &sync.RWMutex{}
	var connect = func() (err error) {
		ctx := context.Background()
		conn, _, err := websocket.Dial(ctx, endpoint, nil)
		if err != nil {
			return
		}
		rwc := wshttp.NewWSConn(conn)
		config := smux.DefaultConfig()
		config.Version = 2
		session, err = smux.Client(rwc, config)
		if err != nil {
			return
		}
		return
	}
	var mustConnect = func() {
		locker.Lock()
		defer locker.Unlock()
		for {
			if err := connect(); err == nil {
				return
			}
			time.Sleep(time.Second)
		}
	}
	go mustConnect()

	var client *http.Client = &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				locker.RLock()
				defer locker.RUnlock()
				conn, err := session.OpenStream()
				if err != nil {
					locker.RUnlock()
					mustConnect()
					locker.RLock()
					conn, err = session.OpenStream()
				}
				return conn, err
			},
		},
	}

	return js.FuncOf(func(this js.Value, args []js.Value) any {
		p, resolve, reject := promise.New()
		go func() {
			var err error
			defer func() {
				if err != nil {
					reject(err.Error())
				}
			}()
			req, err := wahttp.JsRequest(args[0])
			if err != nil {
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				return
			}
			resolve(wahttp.GoResponse(resp))
		}()
		return p
	})
}
