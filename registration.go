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
	"unsafe"
)

type Service struct {
}

type Registration struct {
	service          Service
	flags            int
	interfaceIndex   int
	name             string
	RegistrationType string
	domain           string
	host             string
	Port             int
	TextRecord       string
}

func (r *Registration) TextRecordLength() int {
	return len(r.TextRecord)
}

func (r *Registration) aCallback() {
	fmt.Println("GO A CALLBACK")
}

func (r *Registration) RegisterService() {
	var service C.DNSServiceRef
	var flags C.DNSServiceFlags = 0

	name := C.CString(r.name)
	registrationType := C.CString(r.RegistrationType)
	domain := C.CString(r.domain)
	host := C.CString(r.host)
	// textRecord := C.CString(r.TextRecord)
	textRecord := make([]byte, 0, 4)
	textRecord = append(textRecord, 3)
	textRecord = append(textRecord, []byte("a=b")...)
	defer C.free(unsafe.Pointer(name))
	defer C.free(unsafe.Pointer(registrationType))
	defer C.free(unsafe.Pointer(domain))
	defer C.free(unsafe.Pointer(host))
	// defer C.free(unsafe.Pointer(textRecord))

	fmt.Println(r.RegistrationType, flags, r.interfaceIndex, r.Port)
	errorCode := C.DNSServiceRegister(
		&service,
		flags,
		C.uint32_t(r.interfaceIndex),
		name,
		registrationType,
		domain,
		host,
		C.uint16_t(r.Port),
		// C.uint16_t(r.TextRecordLength()),
		C.uint16_t(len(textRecord)),
		unsafe.Pointer(&textRecord[0]),
		C.serviceRegisterCallbackShim(),
		unsafe.Pointer(r),
	)

	fmt.Println("Registration done:", errorCode)

	if errorCode == C.kDNSServiceErr_NoError {
		errorCode = C.DNSServiceProcessResult(service)
		fmt.Println("process", errorCode)
	}

}

//export goRegistrationCallback
func goRegistrationCallback(
	service C.DNSServiceRef,
	flags C.DNSServiceFlags,
	errorCode C.DNSServiceErrorType,
	name_p *C.char,
	registrationType_p *C.char,
	domain_p *C.char,
	pfoo unsafe.Pointer,
) {
	name := C.GoString(name_p)
	registrationType := C.GoString(registrationType_p)
	domain := C.GoString(domain_p)
	foo := *(*Registration)(pfoo)
	fmt.Println("GO REGISTER CALLBACK", foo.RegistrationType, service, flags, errorCode, name, registrationType, domain)
}
