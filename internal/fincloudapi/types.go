package fincloudapi

type genericResponse[T any] struct {
	ResponseCode string `json:"responseCode"`
	Description  string `json:"description"`
	Data         T      `json:"data"`
}

type genericListResponse[T any, U any] struct {
	ResponseCode string `json:"responseCode"`
	Description  string `json:"description"`
	Data         T      `json:"data"`
	List         []U    `json:"list"`
}

type InquiryAccountStatementsResponse = genericListResponse[
	struct {
		Index string `json:"index"`
		Next  string `json:"next"`
	}, struct {
		Date            string `json:"date"`
		Time            string `json:"time"`
		ReferenceNumber string `json:"referenceNumber"`
		Amount          string `json:"amount"`
		Sign            string `json:"sign"`
		Description     string `json:"description"`
		EndBalance      string `json:"endBalance"`
	},
]

type GlToGlResponse struct {
	DebitAccount         string `json:"debitAccount"`
	CreditAccStatus      string `json:"creditAccStatus"`
	CreditAccountIDType  string `json:"creditAccountIdType"`
	CreditCustomerType   string `json:"creditCustomerType"`
	DebitAccStatus       string `json:"debitAccStatus"`
	DebitResidentStatus  string `json:"debitResidentStatus"`
	CreditCustomer       string `json:"creditCustomer"`
	CreditNationalIDNo   string `json:"creditNationalIdNo"`
	JournalID            string `json:"journalId"`
	DebitAccountIDType   string `json:"debitAccountIdType"`
	CreditAccount        string `json:"creditAccount"`
	CreditAccBalance     string `json:"creditAccBalance"`
	CreditResidentStatus string `json:"creditResidentStatus"`
	DebitAccBalance      string `json:"debitAccBalance"`
	DebitNationalIDNo    string `json:"debitNationalIdNo"`
	DebitCustomerType    string `json:"debitCustomerType"`
	CreditCustName       string `json:"creditCustName"`
}

type GlToSavingResponse struct {
	DebitAccountIDType   string `json:"debitAccountIdType"`
	DebitResidentStatus  string `json:"debitResidentStatus"`
	DebitCustomerType    string `json:"debitCustomerType"`
	CreditCustomer       string `json:"creditCustomer"`
	CreditAccStatus      string `json:"creditAccStatus"`
	CreditCustomerType   string `json:"creditCustomerType"`
	DebitAccount         string `json:"debitAccount"`
	DebitAccBalance      string `json:"debitAccBalance"`
	DebitAccStatus       string `json:"debitAccStatus"`
	DebitNationalIDNo    string `json:"debitNationalIdNo"`
	CreditNationalIDNo   string `json:"creditNationalIdNo"`
	CreditCustName       string `json:"creditCustName"`
	CreditAccBalance     string `json:"creditAccBalance"`
	CreditAccountIDType  string `json:"creditAccountIdType"`
	CreditResidentStatus string `json:"creditResidentStatus"`
	JournalID            string `json:"journalId"`
	CreditAccount        string `json:"creditAccount"`
}

type LoanTerminationResponse struct {
	BranchCode          string `json:"branchCode"`
	CifNo               string `json:"cifNo"`
	CustomerName        string `json:"customerName"`
	InterestPaid        string `json:"interestPaid"`
	InterestWaive       string `json:"interestWaive"`
	PenaltyWaive        string `json:"penaltyWaive"`
	TransactionID       string `json:"transactionId"`
	JournalID           string `json:"journalId"`
	AccountNumber       string `json:"accountNumber"`
	Status              string `json:"status"`
	TrxReference        string `json:"trxReference"`
	CifNoAlt            string `json:"cifNoAlt"`
	Description         string `json:"description"`
	PrincipalPaid       string `json:"principalPaid"`
	PenaltyPaid         string `json:"penaltyPaid"`
	AlternateNumber     string `json:"alternateNumber"`
	RepaymentAccBalance string `json:"repaymentAccBalance"`
}

type CIFInquiryResponse struct {
	IncomeSource     string `json:"incomeSource"`
	BranchCode       string `json:"branchCode"`
	BranchName       string `json:"branchName"`
	BirthDate        string `json:"birthDate"`
	PhoneNo          string `json:"phoneNo"`
	SMSNo            string `json:"smsNo"`
	Village          string `json:"village"`
	JobPhone         string `json:"jobPhone"`
	Address          string `json:"address"`
	Email            string `json:"email"`
	TaxIDNo          string `json:"taxIdNo"`
	Religion         string `json:"religion"`
	City             string `json:"city"`
	Status           string `json:"status"`
	Nationality      string `json:"nationality"`
	BirthPlace       string `json:"birthPlace"`
	LegalDocExp      string `json:"legalDocExp"`
	SubDistrict      string `json:"subDistrict"`
	MotherMaidenName string `json:"motherMaidenName"`
	LegalDocName     string `json:"legalDocName"`
	CellularNo       string `json:"cellularNo"`
	PostalCode       string `json:"postalCode"`
	MaritalStatus    string `json:"maritalStatus"`
	IncomeRange      string `json:"incomeRange"`
	CustomerNo       string `json:"customerNo"`
	Sector           string `json:"sector"`
	AlternateNumber  string `json:"alternateNumber"`
	CustomerType     string `json:"customerType"`
	Province         string `json:"province"`
	JobType          string `json:"jobType"`
	JobPostalCode    string `json:"jobPostalCode"`
	LastEducation    string `json:"lastEducation"`
	Documents        string `json:"documents"`
	JobAddress       string `json:"jobAddress"`
	CustomerName     string `json:"customerName"`
	Gender           string `json:"gender"`
	LegalDocID       string `json:"legalDocId"`
	Job              string `json:"job"`
	JobCompanyName   string `json:"jobCompanyName"`
}

type CreateCIFResponse struct {
	Name            string `json:"name"`
	City            string `json:"city"`
	CifNumber       string `json:"cifNumber"`
	AlternateNumber string `json:"alternateNumber"`
}

type PortfolioInquiryResponse struct {
	Savings []PortfolioKind `json:"savings"`
	Deposit []PortfolioKind `json:"deposit"`
}

type PortfolioKind struct {
	Name          string   `json:"name"`
	Account       string   `json:"account"`
	ProductCode   string   `json:"productCode"`
	ProductName   string   `json:"productName"`
	Status        string   `json:"status"`
	Type          string   `json:"type"`
	Currency      string   `json:"currency"`
	MidRate       float64  `json:"midRate"`
	LedgerBalance float64  `json:"ledgerBalance"`
	Plafond       float64  `json:"plafond"`
	MinBalance    float64  `json:"minBalance"`
	IssueDate     *string  `json:"issueDate,omitempty"`
	MaturityDate  *string  `json:"maturityDate,omitempty"`
	IntRate       *float64 `json:"intRate,omitempty"`
}

type SavingBalanceInquiryResponse struct {
	CustomerName       string `json:"customerName"`
	DocumentStatus     string `json:"documentStatus"`
	MinimumBalance     string `json:"minimumBalance"`
	FacilityLimit      string `json:"facilityLimit"`
	AccountNumber      string `json:"accountNumber"`
	LedgerBalance      string `json:"ledgerBalance"`
	AvailableBalance   string `json:"availableBalance"`
	LockBalance        string `json:"lockBalance"`
	CustNationalIDNo   string `json:"custNationalIdNo"`
	CustomerType       string `json:"customerType"`
	AccountIDType      string `json:"accountIdType"`
	CustomerNumber     string `json:"customerNumber"`
	ProductName        string `json:"productName"`
	AvailableFacility  string `json:"availableFacility"`
	CustResidentStatus string `json:"custResidentStatus"`
	ProductID          string `json:"productId"`
	Currency           string `json:"currency"`
}

type SavingBalanceInquiry2Response struct {
	Account          string `json:"account"`
	Plafond          string `json:"plafond"`
	ClearBalance     string `json:"clearBalance"`
	AvailableBalance string `json:"availableBalance"`
	InactiveMarker   string `json:"inactiveMarker"`
	CustomerName     string `json:"customerName"`
	ProductCode      string `json:"productCode"`
	ProductName      string `json:"productName"`
	Currency         string `json:"currency"`
}

type LoanInquiryResponse struct {
	PrincipalArrears        string `json:"principalArrears"`
	SaForLoanRepayment      string `json:"saForLoanRepayment"`
	ProductName             string `json:"productName"`
	StartPeriod             string `json:"startPeriod"`
	CurrencyCode            string `json:"currencyCode"`
	InterestWriteOff        string `json:"interestWriteOff"`
	CustomerName            string `json:"customerName"`
	ChargeOffDate           string `json:"chargeOffDate"`
	TermPeriod              string `json:"termPeriod"`
	CreditAgreementInterest string `json:"creditAgreementInterest"`
	ProductID               string `json:"productID"`
	Term                    string `json:"term"`
	InstallmentAmount       string `json:"installmentAmount"`
	InterestArrears         string `json:"interestArrears"`
	PenaltyWriteOff         string `json:"penaltyWriteOff"`
	Collectability          string `json:"collectability"`
	NextDueDate             string `json:"nextDueDate"`
	CifNoAlt                string `json:"cifNoAlt"`
	LoanOutStanding         string `json:"loanOutStanding"`
	PenaltyArrears          string `json:"penaltyArrears"`
	WriteOffDate            string `json:"writeOffDate"`
	WriteOffAmount          string `json:"writeOffAmount"`
	Dpd                     string `json:"dpd"`
	CifNo                   string `json:"cifNo"`
	CreditLimit             string `json:"creditLimit"`
	PrincipalWriteOff       string `json:"principalWriteOff"`
	Status                  string `json:"status"`
	BranchCode              string `json:"branchCode"`
	AltNumber               string `json:"altNumber"`
	ChargeOffAmount         string `json:"chargeOffAmount"`
	ClosedDate              string `json:"closedDate"`
	AccrueInterest          string `json:"accrueInterest"`
	EndPeriod               string `json:"endPeriod"`
	InterestRate            string `json:"interestRate"`
	LoanPrincipal           string `json:"loanPrincipal"`
	LastDueDate             string `json:"lastDueDate"`
	CompanyID               string `json:"companyID"`
	AccountNumber           string `json:"accountNumber"`
}

type LoanDisbursementResponse struct {
	LoanNo             string `json:"loanNo"`
	FincloudLoanNo     string `json:"fincloudLoanNo"`
	DisburseNo         string `json:"disburseNo"`
	SourceCustomerName string `json:"sourceCustomerName"`
	AccountNo          string `json:"accountNo"`
	Fullname           string `json:"fullname"`
	PrincipalTotal     string `json:"principalTotal"`
	Installment        string `json:"installment"`
	Cif                string `json:"cif"`
}

type LoanRepaymentResponse struct {
	Description     string `json:"description"`
	JournalID       string `json:"journalId"`
	AlternateNumber string `json:"alternateNumber"`
	CifNoAlt        string `json:"cifNoAlt"`
	CifNo           string `json:"cifNo"`
	CustomerName    string `json:"customerName"`
	Amount          string `json:"amount"`
	Status          string `json:"status"`
	BranchCode      string `json:"branchCode"`
	TransactionID   string `json:"transactionId"`
	TrxReference    string `json:"trxReference"`
	AccountNumber   string `json:"accountNumber"`
}
