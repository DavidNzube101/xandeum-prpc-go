# xandeum-prpc-go

A Go client for interacting with Xandeum pNode pRPC APIs.

## Installation

```bash
go get github.com/DavidNzube101/xandeum-prpc-go
```

## Usage

```go
import (
	"fmt"
	prpc "github.com/DavidNzube101/xandeum-prpc-go"
	"log"
)

func main() {
	client := prpc.NewClient("192.190.136.28") // Replace with a pNode IP

	// Get pods with detailed statistics
	podsWithStats, err := client.GetPodsWithStats()
	if err != nil {
		log.Fatalf("Failed to get pods with stats: %v", err)
	}

	fmt.Printf("Found %d pods with stats:\n", podsWithStats.TotalCount)
	for _, pod := range podsWithStats.Pods {
		fmt.Printf("  Pubkey: %s, Address: %s, Uptime: %d, Storage Used: %d bytes\n",
			pod.Pubkey, pod.Address, pod.Uptime, pod.StorageUsed)
	}
}
```

## API

- `NewClient(ip string)` - Creates a new pRPC client for a given pNode IP.
- `GetPods() (*PodsResponse, error)` - Retrieves the list of pods. (Note: Use `GetPodsWithStats` for more data).
- `GetPodsWithStats() (*PodsResponse, error)` - Retrieves the list of pods with detailed statistics.
- `GetStats() (*NodeStats, error)` - Retrieves the statistics for a single node.

