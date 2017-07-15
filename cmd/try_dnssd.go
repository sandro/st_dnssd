package main

import (
	"fmt"
	"st_dnssd"
)

func main() {
	// fmt.Println("registering")
	// rCallback := func(args st_dnssd.RegistrationCallbackArgs) {
	// 	fmt.Println("callback")
	// 	fmt.Println(args)
	// 	fmt.Println(args.FlagIsAdd())
	// }
	// r := st_dnssd.Registration{RegistrationType: "_go._tcp", Port: 56565, Callback: rCallback}
	// r.RegisterService([]string{"version=1", "base=2"})

	// fmt.Println("length", r.TextRecordLength())

	bCallback := func(args st_dnssd.BrowserCallbackArgs) {
		fmt.Println("callback more coming", args.FlagIsMoreComing())
		fmt.Println(args)
		rCallback := func(args st_dnssd.ResolveReplyArgs) {
			if !args.FlagIsMoreComing() {
				fmt.Println("resolve args", args)
			}
		}
		// resolver := &st_dnssd.Resolver{
		// 	// Name:             args.ServiceName,
		// 	// RegistrationType: args.RegistrationType,
		// 	// Domain:           args.ReplyDomain,
		// 	Callback: rCallback,
		// }
		if !args.FlagIsMoreComing() {
			st_dnssd.Resolve(args.IfIndex, args.ServiceName, args.RegistrationType, args.ReplyDomain, rCallback)
		}
	}
	b := st_dnssd.Browser{RegistrationType: "_go._tcp", Callback: bCallback}
	b.Browse()
	select {}
}

// b := st_dnssd.Browser{RegistrationType: "_http._tcp"}
