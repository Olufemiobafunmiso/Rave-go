package rave

import (
	"Rave-go/rave"
	"testing"
)

var r = rave.Rave {
	false,
	"FLWPUBK_TEST-9070a8d235099f5ac7db26a0356e92b3-X",
	"FLWSECK_TEST-9e54889bc26206218e96152ce4f477f9-X",
	}


var Settlement = rave.Settlements{
	r,
}

func TestSettlement(t *testing.T) {

	payload := rave.ListSettlementData{
		From: "from",
		To:   "to",
		Page: "page",
		Subaccountid: "subaccountid",
		// Id:   "id",
		// Seckey       : "seckey",
	}

	// for _, payload := range payloads {
	error, response := Settlement.List(payload)
	if error != nil {
		t.Fatalf("Settlement List failed with error %v", error)
	}
	if response["status"] != "success" {
		t.Fatalf("Settlement List status: %v, Details: %v", response["status"], response)
	}
	// }

}
