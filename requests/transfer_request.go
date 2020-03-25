package requests

type TransferRequest struct {
	AddressFrom   string  `json:"addressFrom" validate:"required,validHex40"`
	To            string  `json:"to" validate:"required,validHex40or64"`
	CurrencyCode  int     `json:"currencyCode" validate:"required,min=0"`
	Amount        int64   `json:"amount" validate:"required"`
	Payload       string  `json:"payload"`
	MsgHash       string  `json:"msgHash" validate:"required"`
	Sig           SignDto `json:"sig" validate:"required,dive"`
	InvoiceNumber string  `json:"invoiceNumber" validate:"max=255"`
	TransactionId int     `json:"transactionId" validate:"required,min=0"`
}