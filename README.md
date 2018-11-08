# go-coinsuper
[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hexoul/go-coinsuper/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/hexoul/go-coinsuper)](https://goreportcard.com/report/github.com/hexoul/go-coinsuper)
[![GoDoc](https://godoc.org/github.com/hexoul/go-coinsuper?status.svg)](https://godoc.org/github.com/hexoul/go-coinsuper)

> Coinsuper API Client written in Golang

## Usage
- As library, start from `coinsuper.GetInstanceWithKey('YOUR_ACCESS_KEY', 'YOUR_SECRET_KEY')`
- As program, start from `coinsuper.GetInstance()` after executing `go run -coinsuperAccesskey=[YOUR_ACCESS_KEY] -coinsuperSecret=[YOUR_SECRET_KEY]`

## Features
| Type        | Endpoint                        | Done |
|-------------|---------------------------------|------|
| User        | /v1/asset/userAssetInfo         | ✔ |
| Transaction | /v1/order/buy                   | - |
| Transaction | /v1/order/sell                  | - |
| Transaction | /v1/order/cancel                | - |
| Transaction | /v1/order/batchCancel           | - |
| Transaction | /v1/order/list                  | - |
| Transaction | /v1/order/details               | - |
| Transaction | /v1/order/openList              | - |
| Transaction | /v1/order/history               | - |
| Transaction | /v1/order/tradeHistory          | ✔ |
| Quote       | /v1/market/depth                | - |
| Quote       | /v1/market/orderBook            | - |
| Quote       | /v1/market/kline                | - |
| Quote       | /v1/market/tickers              | - |
| Quote       | /v1/market/symbolList           | - |

## Reference
[Coinsuper API docs](https://github.com/coinsuperapi/API_docs_en)
