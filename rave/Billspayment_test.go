package rave

import (
	"testing"
	"Rave-go/rave"
)

// var r = rave.Rave {
// 	false,
// 	"FLWPUBK_TEST-e29f6951e9883f99aa4d0290495d8d07-X",
// 	"FLWSECK_TEST-81e77d815f391cbe7c37acf4505eb1f7-X",
// 	}

var bill = rave.Billpayment{
	r,
}


func TestBill(t *testing.T) {
	 
	payload := rave.FlyBuyData{
			Service:        "fly_recurring",
			ServiceMethod:  "get",
			ServiceVersion: "v1",
			ServiceChannel: "rave",
	}
	
	// for _, payload := range payloads {
		error, response := bill.Bill(payload)
		if error != nil{
			t.Fatalf("Bills Payment failed with error %v",error)
		}
		if response["status"] != "success"{
			t.Fatalf("Bills Payment status: %v, Details: %v",response["status"], response)
		}
	// }

}
