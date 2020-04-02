package responses

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
)

type BankResponse struct {
	Data models.Bank `json:"data"`
	BaseResponse
}

type AccountResponse struct {
	Data models.Account `json:"data"`
	BaseResponse
}

type AccountListResponse struct {
	Data []models.Account `json:"data"`
	BaseResponse
}

type Metadata struct {
	FetchedRecordsCount int32  `json:"fetchedRecordsCount"`
	Bookmark            string `json:"bookmark"`
}

type AccountPageData struct {
	Metadata Metadata         `json:"metadata"`
	Items    []models.Account `json:"items"`
}

type AccountPageResponse struct {
	Data AccountPageData `json:"data"`
	BaseResponse
}

type SenderAddressData struct {
	Address string `json:"address"`
}

type SenderAddressResponse struct {
	Data SenderAddressData `json:"data"`
	BaseResponse
}

type BankListResponse struct {
	Data []models.Bank `json:"data"`
	BaseResponse
}

type BankTotalData struct {
	Total int `json:"total"`
}

type BankTotalResponse struct {
	Data BankTotalData `json:"data"`
	BaseResponse
}

type BankPageData struct {
	Metadata Metadata      `json:"metadata"`
	Items    []models.Bank `json:"items"`
}

type BankPageResponse struct {
	Data BankPageData `json:"data"`
	BaseResponse
}

type SenderIsBankResponse struct {
	Data bool `json:"data"`
	BaseResponse
}

type GetAvailablePlatformsResponse struct {
	Data string `json:"data"`
	BaseResponse
}

type CurrencyListResponse struct {
	Data []models.Currency `json:"data"`
	BaseResponse
}

type CurrencyPageData struct {
	Metadata Metadata          `json:"metadata"`
	Items    []models.Currency `json:"items"`
}

type CurrencyPageResponse struct {
	Data CurrencyPageData `json:"data"`
	BaseResponse
}

type CurrencyResponse struct {
	Data models.Currency `json:"data"`
	BaseResponse
}

type AccountBalanceResponse struct {
	Data AccountBalanceData `json:"data"`
	BaseResponse
}
type AccountBalanceData struct {
	Items []models.AmountOfBank `json:"items"`
	Total int64                 `json:"total"`
}

type WithdrawResultResponse struct {
	Data models.WithdrawResult `json:"data"`
	BaseResponse
}

type WithdrawConfirmResponse struct {
	Data models.TransactionHistory `json:"data"`
	BaseResponse
}

type BankBalanceData struct {
	Liabilities []models.AmountOfBank `json:"liabilities"`
	Claims      []models.AmountOfBank `json:"claims"`
	Issue       int64                 `json:"issue"`
	IssueLimit  int64                 `json:"issueLimit"` // TODO убрать если не быдет использоваться
}

type BankBalanceResponse struct {
	Data BankBalanceData `json:"data"`
	BaseResponse
}

type WithdrawRejectResponse WithdrawConfirmResponse

type ClaimsItemResponse struct {
	CurrencyCode    int             `json:"currencyCode"`
	BankClaims      models.BankInfo `json:"bankClaims"`
	BankLiabilities models.BankInfo `json:"bankLiabilities"`
	Amount          int64           `json:"amount"`
	Unconfirmed     int64           `json:"unconfirmed"`
}

type ClearingResultResponseData struct {
	Id          string                `json:"id"`
	Owner       string                `json:"owner"`
	Claims      int64                 `json:"claims"`
	Liabilities int64                 `json:"liabilities"`
	History     []ClaimsItemResponse  `json:"history"`
	Netting     []models.AmountOfBank `json:"netting"`
	Procedure   []ClaimsItemResponse  `json:"procedure"`
	Created     int64                 `json:"created"`
}

type ClearingResultResponse struct {
	Data ClearingResultResponseData `json:"data"`
	BaseResponse
}

type ClearingListResponse struct {
	Data []ClearingResultResponseData `json:"data"`
	BaseResponse
}

type ClaimsListResponse struct {
	Data []models.ClaimsItem `json:"data"`
	BaseResponse
}

type ClearingPageData struct {
	Metadata Metadata                     `json:"metadata"`
	Items    []ClearingResultResponseData `json:"items"`
}

type ClearingPageResponse struct {
	Data ClearingPageData `json:"data"`
	BaseResponse
}

type ClaimsPageData struct {
	Metadata Metadata            `json:"metadata"`
	Items    []models.ClaimsItem `json:"items"`
}

type ClaimsPageResponse struct {
	Data ClaimsPageData `json:"data"`
	BaseResponse
}

type BankClaimsLiabilities struct {
	Claims      map[string]int64 `json:"claims"`
	Liabilities map[string]int64 `json:"liabilities"`
}
type BankClaimsLiabilitiesResponse struct {
	Data BankClaimsLiabilities `json:"data"`
	BaseResponse
}
