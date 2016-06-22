package main

import "dnssd"

func main() {
	// fmt.Println("registering")
	// r := dnssd.Registration{RegistrationType: "_go._tcp", Port: 56565, TextRecord: "5hello"}
	// r.RegisterService()

	// fmt.Println("length", r.TextRecordLength())

	b := dnssd.Browser{RegistrationType: "_go._tcp"}
	b.Browse()
	select {}
}
