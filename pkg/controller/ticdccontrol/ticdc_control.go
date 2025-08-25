// Package ticdccontrol provides control logic for TiCDC
package ticdccontrol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	"github.com/pingcap/tidb-operator/pkg/controller"
	"github.com/pingcap/tidb-operator/pkg/manager/pdapi"
)

// defaultTiCDCControl is the default implementation of TiCDCControlInterface.
type defaultTiCDCControl struct {
	pdControl  pdapi.PDControl
	ticdcCache *ticdcCache
}

type ticdcCache struct{}

// NewDefaultTiCDCControl returns a defaultTiCDCControl instance
func NewDefaultTiCDCControl(pdControl pdapi.PDControl) controller.TiCDCControlInterface {
	return &defaultTiCDCControl{
		pdControl:  pdControl,
		ticdcCache: &ticdcCache{},
	}
}

// GetChangeFeeds gets the list of changefeeds
func (c *defaultTiCDCControl) GetChangeFeeds(tc *v1alpha1.TidbCluster) ([]controller.ChangeFeedInfo, error) {
	if tc.Spec.TiCDC == nil {
		return nil, fmt.Errorf("TidbCluster %s/%s does not have TiCDC spec", tc.Namespace, tc.Name)
	}

	// Use the first TiCDC pod as the client
	cdcClient := c.getTiCDCClient(tc)

	// Make API call to get changefeeds
	// This is a simplified implementation; in practice, this would call the TiCDC API
	// to get the list of changefeeds
	url := fmt.Sprintf("http://%s:8301/api/v1/changefeeds", cdcClient)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get changefeeds, status code: %d", resp.StatusCode)
	}

	var result map[string]controller.ChangeFeedInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	changefeeds := make([]controller.ChangeFeedInfo, 0, len(result))
	for _, cf := range result {
		changefeeds = append(changefeeds, cf)
	}

	return changefeeds, nil
}

// CreateChangeFeed creates a new changefeed with the specified sink
func (c *defaultTiCDCControl) CreateChangeFeed(tc *v1alpha1.TidbCluster, changefeedName string, sinkURI string) error {
	if tc.Spec.TiCDC == nil {
		return fmt.Errorf("TidbCluster %s/%s does not have TiCDC spec", tc.Namespace, tc.Name)
	}

	// Use the first TiCDC pod as the client
	cdcClient := c.getTiCDCClient(tc)

	// Prepare the changefeed creation request
	changefeedConfig := map[string]interface{}{
		"sink-uri": sinkURI,
		"filter-rules": []string{
			"*.*",
		},
	}

	payload, err := json.Marshal(changefeedConfig)
	if err != nil {
		return err
	}

	// Make API call to create changefeed
	// This is a simplified implementation; in practice, this would call the TiCDC API
	// to create a changefeed
	url := fmt.Sprintf("http://%s:8301/api/v1/changefeeds/%s", cdcClient, changefeedName)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to create changefeed, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}

// getTiCDCClient returns the client address for accessing TiCDC APIs
func (c *defaultTiCDCControl) getTiCDCClient(tc *v1alpha1.TidbCluster) string {
	return fmt.Sprintf("%s-ticdc-0.%s-ticdc.%s.svc", tc.Name, tc.Name, tc.Namespace)
}

// IsTiCDCSynced checks if TiCDC is synced
func (c *defaultTiCDCControl) IsTiCDCSynced(tc *v1alpha1.TidbCluster) bool {
	// In a real implementation, this would check the status of all TiCDC instances
	// and verify they are properly synced with the cluster
	// For simplicity, we're returning true
	return true
}
