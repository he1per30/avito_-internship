package client

type Client struct {
	ID           string  `json:"id" bson:"id,omitempty"`
	creditAmount float64 `json:"creditAmount" bson:"credit_amount"`
}
