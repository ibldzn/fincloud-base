package fincloudapi

import (
	"context"
	"net/http"
)

func (c *Client) InquiryLoan(
	ctx context.Context,
	payload LoanInquiryDTO,
) (*LoanInquiryResponse, error) {
	return doAPI[LoanInquiryResponse](
		c,
		ctx,
		http.MethodPost,
		"/inquiry/detail/loan",
		payload,
		nil,
	)
}

func (c *Client) DisburseLoan(
	ctx context.Context,
	payload LoanDisbursementDTO,
) (*LoanDisbursementResponse, error) {
	return doAPI[LoanDisbursementResponse](
		c,
		ctx,
		http.MethodPost,
		"/disbursement",
		payload,
		nil,
	)
}

func (c *Client) TerminateLoan(
	ctx context.Context,
	payload LoanTerminationDTO,
) (*LoanTerminationResponse, error) {
	return doAPI[LoanTerminationResponse](
		c,
		ctx,
		http.MethodPost,
		"/loan/earlytermination/",
		payload,
		nil,
	)
}
