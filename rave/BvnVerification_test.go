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

var bvn = rave.BVN{
	r,
}


func TestBvn(t *testing.T) {
	
		error, response := bvn.Bvn("12345678901")
		if error != nil{
			t.Fatalf("BVN Verification failed with error %v",error)
		}
		if response["status"] != "success"{
			t.Fatalf("BVN Verification status: %v, Details: %v",response["status"], response)
		}

}
