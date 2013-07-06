package braintree

import (
	"testing"
)

// This test will fail unless you set up your Braintree sandbox account correctly. See TESTING.md for details.
func TestCustomer(t *testing.T) {
	oc := &Customer{
		FirstName: "Lionel",
		LastName:  "Barrow",
		Company:   "Braintree",
		Email:     "lionel.barrow@example.com",
		Phone:     "312.555.1234",
		Fax:       "614.555.5678",
		Website:   "http://www.example.com",
		CreditCard: &CreditCard{
			Number:         testCreditCards["visa"].Number,
			ExpirationDate: "05/14",
			CVV:            "200",
			Options: &CreditCardOptions{
				VerifyCard: true,
			},
		},
	}

	// Create with errors
	_, err := testGateway.Customer().Create(oc)
	if err == nil {
		t.Fatal("Did not receive error when creating invalid customer")
	}

	// Create
	customer, err := testGateway.Customer().Create(&Customer{
		FirstName: "Lionel",
		LastName:  "Barrow",
		Company:   "Braintree",
		Email:     "lionel.barrow@example.com",
		Phone:     "312.555.1234",
		Fax:       "614.555.5678",
		Website:   "http://www.example.com",
		CreditCard: &CreditCard{
			Number:         testCreditCards["visa"].Number,
			ExpirationDate: "05/14",
			CVV:            "",
			Options:        nil,
		},
	})

	t.Log(customer)

	if err != nil {
		t.Fatal(err)
	}
	if customer.Id == "" {
		t.Fatal("invalid customer id")
	}

	// Update
	c2, err := testGateway.Customer().Update(&Customer{
		Id:        customer.Id,
		FirstName: "John",
	})

	t.Log(c2)

	if err != nil {
		t.Fatal(err)
	}
	if c2.FirstName != "John" {
		t.Fatal("first name not changed")
	}

	// Find
	c3, err := testGateway.Customer().Find(customer.Id)

	t.Log(c3)

	if err != nil {
		t.Fatal(err)
	}
	if c3.Id != customer.Id {
		t.Fatal("ids do not match")
	}

	// Delete
	err = testGateway.Customer().Delete(customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	// Test customer 404
	c4, err := testGateway.Customer().Find(customer.Id)
	if err == nil {
		t.Fatal("should return 404")
	}
	if err.Error() != "Not Found (404)" {
		t.Fatal(err)
	}
	if c4 != nil {
		t.Fatal(c4)
	}
}
