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

// Registration can be used to register a service
type Registration struct {
	Service          C.DNSServiceRef
	flags            int
	interfaceIndex   int
	name             string
	RegistrationType string
	domain           string
	host             string
	Port             int
	TextRecord       string
}

// RegisterService advertises a dnssd service
func (r *Registration) RegisterService() {
	var flags C.DNSServiceFlags

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
		&r.Service,
		flags,
		C.uint32_t(r.interfaceIndex),
		name,
		registrationType,
		domain,
		host,
		C.uint16_t(r.Port),
		// C.uint16_t(r.textRecordLength()),
		C.uint16_t(len(textRecord)),
		unsafe.Pointer(&textRecord[0]),
		(C.DNSServiceRegisterReply)(C.registrationCallback),
		unsafe.Pointer(r),
	)

	fmt.Println("Registration done:", errorCode)

	if errorCode == C.kDNSServiceErr_NoError {
		errorCode = C.DNSServiceProcessResult(r.Service)
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
	pRegistration unsafe.Pointer,
) {
	name := C.GoString(name_p)
	registrationType := C.GoString(registrationType_p)
	domain := C.GoString(domain_p)
	registration := *(*Registration)(pRegistration)
	fmt.Println("GO REGISTER CALLBACK", registration.RegistrationType, service, flags, errorCode, name, registrationType, domain)
}
