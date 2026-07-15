package fincloud

import (
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) GetLoanAccountFromAltNumber(ctx context.Context, altNumber string) (string, error) {
	resp, err := c.DoRequest(ctx, func() (*http.Request, error) {
		req, err := c.NewRequest(ctx, http.MethodGet, "/pinjaman/inquiry/rekening/cari", nil)
		if err != nil {
			return nil, err
		}

		q := req.URL.Query()
		q.Set("cabang", "ALL")
		q.Set("noalt", altNumber)
		q.Set("pagesize", "50")
		req.URL.RawQuery = q.Encode()

		return req, nil
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", ErrDataFetchFailed
	}

	var intermediate struct {
		Data struct {
			Result []struct {
				AccountNumber string `json:"id"`
			} `json:"result"`
		} `json:"data"`
		Status string `json:"status"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&intermediate); err != nil {
		return "", err
	}

	if intermediate.Status != "ok" {
		return "", ErrDataFetchFailed
	}

	if len(intermediate.Data.Result) != 1 {
		return "", ErrDataNotFound
	}

	return intermediate.Data.Result[0].AccountNumber, nil
}
