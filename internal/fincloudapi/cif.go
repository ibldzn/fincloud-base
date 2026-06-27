package fincloudapi

import (
	"context"
	"net/http"
)

func (c *Client) InquiryCIF(
	ctx context.Context,
	payload CIFInquiryRequest,
) (*CIFInquiryResponse, error) {
	return doAPI[CIFInquiryResponse](
		c,
		ctx,
		http.MethodGet,
		"/customer",
		nil,
		func(req *http.Request) error {
			q := req.URL.Query()
			addNonEmptyQuery(q, "accountNumber", payload.AccountNumber)
			addNonEmptyQuery(q, "nationalIdNo", payload.NationalIdNumber)
			addNonEmptyQuery(q, "cifNo", payload.CIFNumber)
			req.URL.RawQuery = q.Encode()

			return nil
		},
	)
}
