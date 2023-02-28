# Changelog

## [0.0.4] - 2023-02-28

### Change

- WebSocket 重连默认 10 次. 可通过 `WageMaxRetry` 设置

## [0.0.3] - 2023-02-28

### Improve

- 可直接在 node 中使用 GoFetch 了

## [0.0.2] - 2023-02-27

### Fix

- `DisableKeepAlives` 每次链接都打开一个新的 `smux stream`
