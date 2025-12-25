package rating

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func New(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// RebuildCategoryTop вызывает сервис rating для пересчёта топа категории
func (c *Client) RebuildCategoryTop(category string) error {
	u := fmt.Sprintf(
		"%s/get_category_top?categoryName=%s",
		c.baseURL,
		url.QueryEscape(category),
	)

	resp, err := c.client.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("rating service returned status %d", resp.StatusCode)
	}

	return nil
}
