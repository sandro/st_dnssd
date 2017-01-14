package st_dnssd

func OrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
