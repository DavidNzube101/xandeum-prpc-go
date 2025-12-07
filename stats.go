package prpc

import (
	"encoding/json"
	"fmt"
)

// NodeStats represents the statistics of a pNode
type NodeStats struct {
	ActiveStreams   int     `json:"active_streams"`
	CPUPercent      float64 `json:"cpu_percent"`
	CurrentIndex    int     `json:"current_index"`
	FileSize        int64   `json:"file_size"`
	LastUpdated     int64   `json:"last_updated"`
	PacketsReceived int     `json:"packets_received"`
	PacketsSent     int     `json:"packets_sent"`
	RAMTotal        int64   `json:"ram_total"`
	RAMUsed         int64   `json:"ram_used"`
	TotalBytes      int64   `json:"total_bytes"`
	TotalPages      int     `json:"total_pages"`
	Uptime          int64   `json:"uptime"`
}

// GetStats retrieves the statistics from a pNode
func (c *Client) GetStats() (*NodeStats, error) {
	resp, err := c.call("get-stats", nil)
	if err != nil {
		return nil, err
	}

	var stats NodeStats
	resultBytes, err := json.Marshal(resp.Result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result: %w", err)
	}

	if err := json.Unmarshal(resultBytes, &stats); err != nil {
		return nil, fmt.Errorf("failed to unmarshal stats response: %w", err)
	}

	return &stats, nil
}
