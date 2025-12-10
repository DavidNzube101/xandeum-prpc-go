package prpc

import (
	"encoding/json"
	"fmt"
)

// Pod represents a pod in the gossip network
type Pod struct {
	Address             string  `json:"address"`
	IsPublic            *bool   `json:"is_public"`
	LastSeenTimestamp   int64   `json:"last_seen_timestamp"`
	Pubkey              string  `json:"pubkey"`
	RPCPort             int     `json:"rpc_port"`
	StorageCommitted    int64   `json:"storage_committed"`
	StorageUsagePercent float64 `json:"storage_usage_percent"`
	StorageUsed         int64   `json:"storage_used"`
	Uptime              int64   `json:"uptime"`
	Version             string  `json:"version"`
}

// PodsResponse represents the response from get-pods
type PodsResponse struct {
	Pods       []Pod `json:"pods"`
	TotalCount int   `json:"total_count"`
}

// GetPods retrieves the list of pods from a pNode
func (c *Client) GetPods() (*PodsResponse, error) {
	resp, err := c.call("get-pods", nil)
	if err != nil {
		return nil, err
	}

	var podsResp PodsResponse
	resultBytes, err := json.Marshal(resp.Result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result: %w", err)
	}

	if err := json.Unmarshal(resultBytes, &podsResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal pods response: %w", err)
	}

	return &podsResp, nil
}

// GetPodsWithStats retrieves the list of pods with detailed statistics
func (c *Client) GetPodsWithStats() (*PodsResponse, error) {
	resp, err := c.call("get-pods-with-stats", nil)
	if err != nil {
		return nil, err
	}

	var podsResp PodsResponse
	resultBytes, err := json.Marshal(resp.Result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result: %w", err)
	}

	if err := json.Unmarshal(resultBytes, &podsResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal pods response: %w", err)
	}

	return &podsResp, nil
}
