package dnssd

/*
#include <stdio.h>
#include <errno.h>
#include <stdlib.h>
#include <dns_sd.h>
#include "callback.h"
*/
import "C"
import (
	"fmt"
	"io"
	"os"
	"unsafe"
)

type Browser struct {
	Service          C.DNSServiceRef
	Flags            int
	InterfaceIndex   int
	RegistrationType string
	Domain           string
}

//export GoBrowseCallback
func GoBrowseCallback(browser unsafe.Pointer) {
	b1 := *(*Browser)(browser)
	fmt.Println("in go with browser", browser, b1)
}

func (o *Browser) Browse() {
	var flags C.DNSServiceFlags = 0

	registrationType := C.CString(o.RegistrationType)
	defer C.free(unsafe.Pointer(registrationType))

	domain := C.CString(o.Domain)
	defer C.free(unsafe.Pointer(domain))

	errorCode := C.DNSServiceBrowse(
		&o.Service,
		flags,
		C.uint32_t(o.InterfaceIndex),
		registrationType,
		nil,
		(C.DNSServiceBrowseReply)(C.bCallback),
		unsafe.Pointer(o),
	)
	fmt.Println("Browse done. error code:", errorCode)

	if errorCode == C.kDNSServiceErr_NoError {
		fd := C.DNSServiceRefSockFD(unsafe.Pointer(o.Service))
		fmt.Println("FD", fd)
		// file, err := os.OpenFile("file", os.O_RDWR, os.ModePerm)
		file := os.NewFile(uintptr(fd), "dnssd_fd")
		fmt.Println(file)
		go func() {
			bb := make([]byte, 100)
			fmt.Println("entering loop")
			for {
				n, err := file.Read(bb)
				if n == 0 && err == io.EOF {
					fmt.Println("EOF")
				} else {
					fmt.Println(string(bb))
				}
			}
			errorCode = C.DNSServiceProcessResult(o.Service)
			fmt.Println("process", errorCode)
		}()
	}
}
