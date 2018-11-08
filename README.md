# go-coinsuper
[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hexoul/go-coinsuper/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/hexoul/go-coinsuper)](https://goreportcard.com/report/github.com/hexoul/go-coinsuper)
[![GoDoc](https://godoc.org/github.com/hexoul/go-coinsuper?status.svg)](https://godoc.org/github.com/hexoul/go-coinsuper)

> Coinsuper API Client written in Golang

## Usage
- As library, start from `coinsuper.GetInstanceWithKey('YOUR_ACCESS_KEY', 'YOUR_SECRET_KEY')`
- As program, start from `coinsuper.GetInstance()` after executing `go run -accesskey=[YOUR_ACCESS_KEY] -secret=[YOUR_SECRET_KEY]`

## Features
| Type | Endpoint                               | Done |
|------|----------------------------------------|------|
| Auth | /v1/asset/userAssetInfo                | ✔ |
| Auth | /v1/order/history                      | - |
| Auth | /v1/order/tradeHistory                 | - |

## Reference
[Coinsuper API docs](https://github.com/coinsuperapi/API_docs_en)
