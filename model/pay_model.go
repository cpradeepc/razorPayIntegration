package model

//page data interate with razorpay gui , these should be migrate with DB
type PageVariables struct {
	UsrKey      string  `json:"user_key,omitempty"`
	OrderId     string  `json:"order_id,omitempty"`
	PayId       string  `json:"pay_id,omitempty"`
	Email       string  `json:"email"`
	Name        string  `json:"name"`
	Amount      int     `json:"amount"`
	Contact     string  `json:"contact"`
	Currency    string  `json:"currency"`
	Receipt     string  `json:"receipt"`
	OrgName     string  `json:"org_name,omitempty"`
	Description string  `json:"description,omitempty"`
	ImgUrl      string  `json:"img_url,omitempty"`
	Status      string  `json:"status,omitempty"`
	Created_At  float64 `json:"created_at,omitempty"`
}

//payment success
type PaySuccess struct {
	PaymentId string
	OrderId   string
	Signature string
}

//payment failed
type PayFailed struct {
	OrderId     string
	PayementId  string
	Code        string
	Description string
	Source      string
	Step        string
	Reason      string
}
