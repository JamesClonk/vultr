package lib

// why does Vultr return only the balance as a json number? Beats me..

// AccountInfo of Vultr account
type AccountInfo struct {
	Balance           float64 `json:"balance"`
	PendingCharges    float64 `json:"pending_charges"`
	LastPaymentDate   string  `json:"last_payment_date"`
	LastPaymentAmount string  `json:"last_payment_amount"`
}

func (c *Client) GetAccountInfo() (info AccountInfo, err error) {
	if err := c.get(`account/info`, &info); err != nil {
		return AccountInfo{}, err
	}
	return
}
