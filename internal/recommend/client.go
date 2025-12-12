package recommend

import (
	"fmt"
	"net/http"
	"net/url"
)

func (s *Service) callRatingUpdate(category string) error {
	// экранируем category
	u := fmt.Sprintf("%s/get_category_top?categoryName=%s", s.ratingBase, url.QueryEscape(category))
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("rating returned status %d", resp.StatusCode)
	}
	return nil
}
