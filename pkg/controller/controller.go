package controller

// CaptureStatus holds CDC server information
// Copyright 2023 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
)

// CaptureStatusInfo holds CDC server information
type CaptureStatusInfo struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

// ChangeFeedInfo represents information about a TiCDC changefeed
type ChangeFeedInfo struct {
	ID      string `json:"id"`
	SinkURI string `json:"sink-uri"`
	State   string `json:"state"`
}

// TiCDCControlInterface implements the control logic for interacting with TiCDC
// It is implemented as an interface to allow for testing.
type TiCDCControlInterface interface {
	// GetChangeFeeds returns a list of changefeeds
	GetChangeFeeds(tc *v1alpha1.TidbCluster) ([]ChangeFeedInfo, error)
	// CreateChangeFeed creates a new changefeed with the specified sink
	CreateChangeFeed(tc *v1alpha1.TidbCluster, changefeedName string, sinkURI string) error
}

// CaptureInfo holds CDC captured info
type CaptureInfo struct {
	ID            string `json:"id"`
	IsOwner       bool   `json:"is-owner"`
	AdvertiseAddr string `json:"address"`
}

// ChangeFeedInfo holds information about a TiCDC changefeed
type ChangeFeedInfo struct {
	ID          string            `json:"id"`
	State       string            `json:"state"`
	SinkURI     string            `json:"sink-uri"`
	CreatedAt   string            `json:"create-time"`
	TableFilter []string          `json:"table-filter,omitempty"`
	Config      map[string]string `json:"config,omitempty"`
}
