package xendit

import (
	"fmt"
	"time"
)

type VirtualAccount struct {
	client *Client
}

type AvailableVirtualAccountResponse struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CreateFixedVirtualAccountRequest struct {
	ExternalId           string    `json:"external_id,omitempty"`
	BankCode             string    `json:"bank_code,omitempty"`
	Name                 string    `json:"name,omitempty"`
	VirtualAccountNumber string    `json:"virtual_account_number,omitempty"`
	SuggestedAmount      float64   `json:"suggested_amount,omitempty"`
	IsClosed             bool      `json:"is_closed,omitempty"`
	ExpectedAmount       float64   `json:"expected_amount,omitempty"`
	ExpirationDate       time.Time `json:"expiration_date,omitempty"`
	IsSingleUse          bool      `json:"is_single_use,omitempty"`
}

type CreateFixedVirtualAccountResponse struct {
	OwnerId         string    `json:"owner_id"`
	ExternalId      string    `json:"external_id"`
	BankCode        string    `json:"bank_code"`
	MerchantCode    string    `json:"merchant_code"`
	Name            string    `json:"name"`
	AccountNumber   string    `json:"account_number"`
	SuggestedAmount float64   `json:"suggested_amount"`
	IsClosed        bool      `json:"is_closed"`
	ExpectedAmount  float64   `json:"expected_amount"`
	Id              string    `json:"id"`
	IsSingleUse     bool      `json:"is_single_use"`
	ExpirationDate  time.Time `json:"expiration_date"`
	Status          string    `json:"status"`
}

type UpdateFixedVirtualAccountRequest struct {
	SuggestedAmount float64   `json:"suggested_amount,omitempty"`
	ExpectedAmount  float64   `json:"expected_amount,omitempty"`
	ExpirationDate  time.Time `json:"expiration_date,omitempty"`
	IsSingleUse     bool      `json:"is_single_use,omitempty"`
}

type UpdateFixedVirtualAccountResponse struct {
	OwnerId         string  `json:"owner_id"`
	ExternalId      string  `json:"external_id"`
	BankCode        string  `json:"bank_code"`
	MerchantCode    string  `json:"merchant_code"`
	Name            string  `json:"name"`
	AccountNumber   string  `json:"account_number"`
	SuggestedAmount float64 `json:"suggested_amount"`
	IsClosed        bool    `json:"is_closed"`
	ExpectedAmount  float64 `json:"expected_amount"`
	Id              string  `json:"id"`
	IsSingleUse     bool    `json:"is_single_use"`
	Status          string  `json:"status"`
}

type FixedVirtualAccountPaymentResponse struct {
	Id                       string    `json:"id"`
	PaymentId                string    `json:"payment_id"`
	CallbackVirtualAccountId string    `json:"callback_virtual_account_id"`
	ExternalId               string    `json:"external_id"`
	MerchantCode             string    `json:"merchant_code"`
	AccountNumber            string    `json:"account_number"`
	BankCode                 string    `json:"bank_code"`
	Amount                   float64   `json:"amount"`
	TransactionTimestamp     time.Time `json:"transaction_timestamp"`
}

type GetFixedVirtualAccountResponse struct {
	OwnerId       string `json:"owner_id"`
	ExternalId    string `json:"external_id"`
	BankCode      string `json:"bank_code"`
	MerchantCode  string `json:"merchant_code"`
	Name          string `json:"name"`
	AccountNumber string `json:"account_number"`
	IsSingleUse   bool   `json:"is_single_use"`
	Status        string `json:"status"`
	IsClosed      bool   `json:"is_closed"`
	Id            string `json:"id"`
}

type FixedVirtualAccountCallbackRequest struct {
	PaymentId                string  `json:"payment_id,omitempty"`
	CallbackVirtualAccountId string  `json:"callback_virtual_account_id,omitempty"`
	OwnerId                  string  `json:"owner_id,omitempty"`
	ExternalId               string  `json:"external_id,omitempty"`
	AccountNumber            string  `json:"account_number,omitempty"`
	BankCode                 string  `json:"bank_code,omitempty"`
	Amount                   float64 `json:"amount,omitempty"`
	MerchantCode             string  `json:"merchant_code,omitempty"`
	Id                       string  `json:"id,omitempty"`
}

func (c *VirtualAccount) GetAvailableVirtualAccount() ([]AvailableVirtualAccountResponse, error) {
	res := new([]AvailableVirtualAccountResponse)

	endpoint := "available_virtual_account_banks"

	err := c.client.Request("GET", endpoint, nil, res)
	return *res, err
}

func (c *VirtualAccount) CreateFixedVirtualAccount(body CreateFixedVirtualAccountRequest) (CreateFixedVirtualAccountResponse, error) {
	res := new(CreateFixedVirtualAccountResponse)

	endpoint := "callback_virtual_accounts"

	err := c.client.Request("POST", endpoint, body, res)
	return *res, err
}

func (c *VirtualAccount) UpdateFixedVirtualAccount(fixedVirtualAccountId string, body UpdateFixedVirtualAccountRequest) (UpdateFixedVirtualAccountResponse, error) {
	res := new(UpdateFixedVirtualAccountResponse)

	endpoint := fmt.Sprintf("callback_virtual_accounts/%s", fixedVirtualAccountId)

	err := c.client.Request("PATCH", endpoint, body, res)
	return *res, err
}

func (c *VirtualAccount) GetFixedVirtualAccountPayment(paymentId string) (FixedVirtualAccountPaymentResponse, error) {
	res := new(FixedVirtualAccountPaymentResponse)

	endpoint := fmt.Sprintf("callback_virtual_account_payments/paymeny_id=%s", paymentId)

	err := c.client.Request("GET", endpoint, nil, res)
	return *res, err
}

func (c *VirtualAccount) GetFixedVirtualAccount(callbackVirtualAccountId string) (GetFixedVirtualAccountResponse, error) {
	res := new(GetFixedVirtualAccountResponse)

	endpoint := fmt.Sprintf("callback_virtual_accounts/%s", callbackVirtualAccountId)

	err := c.client.Request("GET", endpoint, nil, res)
	return *res, err
}
