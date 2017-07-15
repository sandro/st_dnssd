package st_dnssd

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
	"unsafe"
)

type Browser struct {
	Service          C.DNSServiceRef
	InterfaceIndex   int
	RegistrationType string
	Domain           string
	Callback         func(BrowserCallbackArgs)
}

type BrowserCallbackArgs struct {
	FlagAnalyzer
	service          C.DNSServiceRef
	IfIndex          uint32
	ErrorCode        C.DNSServiceErrorType
	ServiceName      string
	RegistrationType string
	ReplyDomain      string
	Browser          Browser
}

type FlagAnalyzer struct {
	flags C.DNSServiceFlags
}

func (o *FlagAnalyzer) FlagIsMoreComing() bool {
	return o.flags&C.kDNSServiceFlagsMoreComing == C.kDNSServiceFlagsMoreComing
}
func (o *FlagAnalyzer) FlagIsAdd() bool {
	return o.flags&C.kDNSServiceFlagsAdd == C.kDNSServiceFlagsAdd
}
func (o *FlagAnalyzer) FlagIsRemove() bool {
	return !o.FlagIsAdd()
}

//export goBrowseCallback
func goBrowseCallback(
	service C.DNSServiceRef,
	flags C.DNSServiceFlags,
	ifIndex C.uint32_t,
	errorCode C.DNSServiceErrorType,
	serviceName *C.char,
	registrationType *C.char,
	replyDomain *C.char,
	pBrowser unsafe.Pointer,
) {
	browser := *(*Browser)(pBrowser)
	callbackArgs := BrowserCallbackArgs{
		service:   service,
		IfIndex:   uint32(ifIndex),
		ErrorCode: errorCode,
		Browser:   browser,
	}
	callbackArgs.FlagAnalyzer.flags = flags
	if errorCode == C.kDNSServiceErr_NoError {
		callbackArgs.ServiceName = C.GoString(serviceName)
		callbackArgs.RegistrationType = C.GoString(registrationType)
		callbackArgs.ReplyDomain = C.GoString(replyDomain)
	}
	browser.Callback(callbackArgs)
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
		C.DNSServiceBrowseReply(C.browseCallback),
		unsafe.Pointer(o),
	)
	fmt.Println("Browse done. error code:", errorCode)

	if errorCode == C.kDNSServiceErr_NoError {
		// srv := unsafe.Pointer(o.Service)
		// fd := C.DNSServiceRefSockFD(C.DNSServiceRef(srv))
		// fmt.Println("FD", fd)
		// file, err := os.OpenFile("file", os.O_RDWR, os.ModePerm)
		// file := os.NewFile(uintptr(fd), "dnssd_fd")
		// fmt.Println(file)
		go func() {
			// bb := make([]byte, 100)
			fmt.Println("entering loop")
			for {
				fmt.Println("in loop")
				// n, err := file.Read(bb)
				// if n == 0 && err == io.EOF {
				// 	fmt.Println("EOF")
				// } else {
				// fmt.Println("printing", string(bb))
				errorCode = C.DNSServiceProcessResult(o.Service)
				fmt.Println("process", errorCode)
				// }
			}
		}()
	}
}
