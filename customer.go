package braintree

type Customer struct {
	XMLName     string       `xml:"customer"`
	Id          string       `xml:"id,omitempty"`
	FirstName   string       `xml:"first-name,omitempty"`
	LastName    string       `xml:"last-name,omitempty"`
	Company     string       `xml:"company,omitempty"`
	Email       string       `xml:"email,omitempty"`
	Phone       string       `xml:"phone,omitempty"`
	Fax         string       `xml:"fax,omitempty"`
	Website     string       `xml:"website,omitempty"`
	CreditCard  *CreditCard  `xml:"credit-card,omitempty"`
	CreditCards *CreditCards `xml:"credit-cards,omitempty"`
}

// DefaultPaymentToken returns the payment token of the first creditcard, if any.
func (c *Customer) DefaultPaymentToken() string {
	for _, card := range c.CreditCards.CreditCard {
		if card.Default {
			return card.Token
		}
	}
	return ""
}
