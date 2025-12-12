# xandeum-prpc-go

A Go client for interacting with Xandeum pNode pRPC APIs.

## Installation

```bash
go get github.com/DavidNzube101/xandeum-prpc-go
```

## Usage

### Basic Usage
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

### With Custom Timeout
```go
import (
	"fmt"
	prpc "github.com/DavidNzube101/xandeum-prpc-go"
	"log"
	"time"
)

func main() {
	// Creates a client with a 10-second timeout
	client := prpc.NewClient("192.190.136.28", 10*time.Second)

	stats, err := client.GetStats()
	if err != nil {
		log.Fatalf("Failed to get stats: %v", err)
	}
	fmt.Printf("Node uptime: %d seconds\n", stats.Uptime)
}
```

### Finding a pNode
The library includes a helper function to concurrently search a list of seed nodes to find a specific pNode by its public key.

```go
import (
	"fmt"
	prpc "github.com/DavidNzube101/xandeum-prpc-go"
	"log"
)

func main() {
	// Find a node using the default seed list

pod, err := prpc.FindPNode("2asTHq4vVGazKrmEa3YTXKuYiNZBdv1cQoLc1Tr2kvaw", nil)
	if err != nil {
		log.Fatalf("Failed to find pNode: %v", err)
	}
	fmt.Printf("Found pod: %+v\n", pod)

	// Find a node using a custom seed list
	options := &prpc.FindPNodeOptions{
		ReplaceSeeds: []string{"192.190.136.28"},
	}

pod, err = prpc.FindPNode("GCoCP7CLvVivuWUH1sSA9vMi9jjaJcXpMwVozMVA6yBg", options)
	if err != nil {
		log.Fatalf("Failed to find pNode on custom seed: %v", err)
	}
	fmt.Printf("Found pod on custom seed: %+v\n", pod)
}
```

The default seed IPs are:
```
"173.212.220.65", "161.97.97.41", "192.190.136.36", "192.190.136.38", 
"207.244.255.1", "192.190.136.28", "192.190.136.29", "173.212.203.145"
```

## API

- `NewClient(ip string, timeout ...time.Duration)` - Creates a new pRPC client. Accepts an optional `time.Duration` to set the HTTP timeout. Defaults to 8 seconds.
- `FindPNode(nodeId string, options *FindPNodeOptions) (*Pod, error)` - Concurrently searches a list of seed IPs to find a specific pNode by its public key.
- `GetPods() (*PodsResponse, error)` - Retrieves the list of pods. (Note: Use `GetPodsWithStats` for more data).
- `GetPodsWithStats() (*PodsResponse, error)` - Retrieves the list of pods with detailed statistics.
- `GetStats() (*NodeStats, error)` - Retrieves the statistics for a single node.

```

