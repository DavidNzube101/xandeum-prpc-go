# xandeum-prpc-go

A Go client for interacting with Xandeum pNode pRPC APIs.

## Installation

```bash
go get github.com/DavidNzube101/xandeum-prpc-go
```

## Usage

```go
import "github.com/DavidNzube101/xandeum-prpc-go"

client := prpc.NewClient("192.190.136.28")
stats, err := client.GetStats()
```

