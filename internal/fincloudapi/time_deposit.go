package fincloudapi

import (
	"context"
	"net/http"
)

func (c *Client) InquiryTimeDepositProducts(
	ctx context.Context,
) ([]TimeDepositProduct, error) {
	return doAPIList[TimeDepositProduct](
		c,
		ctx,
		http.MethodGet,
		"/deposit/product/list",
		nil,
		nil,
	)
}
