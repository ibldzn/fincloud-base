package fincloudapi

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) InquirySavingBalance(
	ctx context.Context,
	accountNumber string,
) (*SavingBalanceInquiryResponse, error) {
	return doAPI[SavingBalanceInquiryResponse](
		c,
		ctx,
		http.MethodGet,
		"/saving/inq/balance",
		nil,
		func(req *http.Request) error {
			q := req.URL.Query()
			addNonEmptyQuery(q, "accountNumber", accountNumber)
			req.URL.RawQuery = q.Encode()

			return nil
		},
	)
}

func (c *Client) InquirySavingTransactionHistory(
	ctx context.Context,
	payload SavingStatementInquiryRequest,
) (*InquiryAccountStatementsResponse, error) {
	return doAPI[InquiryAccountStatementsResponse](
		c,
		ctx,
		http.MethodGet,
		"/account/statement",
		nil,
		func(req *http.Request) error {
			q := req.URL.Query()
			addNonEmptyQuery(q, "accountNo", payload.AccountNumber)
			addNonEmptyQuery(q, "startDate", payload.StartDate)
			addNonEmptyQuery(q, "endDate", payload.EndDate)
			addNonEmptyQuery(q, "rowPerPage", fmt.Sprintf("%d", payload.RowPerPage))
			addNonEmptyQuery(q, "index", fmt.Sprintf("%d", payload.Index))
			req.URL.RawQuery = q.Encode()

			return nil
		},
	)
}
