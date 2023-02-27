module github.com/shynome/wage/pkg/wshttp/client

go 1.20

replace github.com/shynome/wage => ../../../

require (
	github.com/nlepage/go-js-promise v1.1.0
	github.com/shynome/wage v0.0.0-00010101000000-000000000000
	github.com/shynome/wahttp v0.0.2
	github.com/xtaci/smux v1.5.24
	nhooyr.io/websocket v1.8.7
)

require github.com/klauspost/compress v1.10.3 // indirect
