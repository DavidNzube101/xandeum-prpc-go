package prpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type RPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	ID      interface{} `json:"id"`
}

type RPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
	ID      interface{} `json:"id"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Client struct {
	httpClient *http.Client
	baseURL    string
}

const (
	DefaultTimeout = 8 * time.Second
)

var (
	DefaultSeedIPs = []string{
		"173.212.220.65",
		"161.97.97.41",
		"192.190.136.36",
		"192.190.136.38",
		"207.244.255.1",
		"192.190.136.28",
		"192.190.136.29",
		"173.212.203.145",
	}
)

func NewClient(ip string, timeout ...time.Duration) *Client {
	clientTimeout := DefaultTimeout
	if len(timeout) > 0 {
		clientTimeout = timeout[0]
	}
	return &Client{
		httpClient: &http.Client{
			Timeout: clientTimeout,
		},
		baseURL: fmt.Sprintf("http://%s:6000/rpc", ip),
	}
}

func (c *Client) call(method string, params interface{}) (*RPCResponse, error) {
	req := RPCRequest{
		JSONRPC: "2.0",
		Method:  method,
		ID:      1,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.httpClient.Post(c.baseURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	var rpcResp RPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if rpcResp.Error != nil {
		return nil, fmt.Errorf("RPC error: %s", rpcResp.Error.Message)
	}

	return &rpcResp, nil
}

type FindPNodeOptions struct {
	AddSeeds     []string
	ReplaceSeeds []string
	Timeout      time.Duration
}

func FindPNode(nodeId string, options *FindPNodeOptions) (*Pod, error) {
	opts := options
	if opts == nil {
		opts = &FindPNodeOptions{}
	}

	timeout := opts.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	seeds := DefaultSeedIPs
	if opts.ReplaceSeeds != nil {
		seeds = opts.ReplaceSeeds
	} else if opts.AddSeeds != nil {
		seeds = append(seeds, opts.AddSeeds...)
	}

	resultChan := make(chan *Pod, len(seeds))
	errChan := make(chan error, len(seeds))

	for _, seedIP := range seeds {
		go func(ip string) {
			client := NewClient(ip, timeout)
			podsResp, err := client.GetPods()
			if err != nil {
				errChan <- fmt.Errorf("failed to get pods from seed %s: %w", ip, err)
				return
			}
			for _, pod := range podsResp.Pods {
				if pod.Pubkey == nodeId {
					resultChan <- &pod
					return
				}
			}
			errChan <- nil
		}(seedIP)
	}

	for i := 0; i < len(seeds); i++ {
		select {
		case pod := <-resultChan:
			return pod, nil
		case err := <-errChan:
			if err != nil {
				log.Printf("Seed query error: %v", err)
			}
		case <-ctx.Done():
			return nil, fmt.Errorf("timed out waiting for pNode %s from seeds", nodeId)
		}
	}

	return nil, fmt.Errorf("pNode %s not found on any of the seeds", nodeId)
}
