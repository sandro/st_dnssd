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
	// Service          C.DNSServiceRef
	flags            uint32
	interfaceIndex   int
	name             string
	RegistrationType string
	domain           string
	host             string
	Port             int
	Callback         func(RegistrationCallbackArgs)
}

type Z struct {
	hh  string
	ary *[]string
	ff  func(RegistrationCallbackArgs)
}

type RegistrationCallbackArgs struct {
	FlagAnalyzer
	service          C.DNSServiceRef
	flags            C.DNSServiceFlags
	errorCode        C.DNSServiceErrorType
	name             string
	registrationType string
	domain           string
	Registration     Registration
}

// RegisterService advertises a dnssd service
func (r *Registration) RegisterService(textRecords []string) {
	var flags C.DNSServiceFlags

	name := C.CString(r.name)
	defer C.free(unsafe.Pointer(name))
	registrationType := C.CString(r.RegistrationType)
	defer C.free(unsafe.Pointer(registrationType))
	domain := C.CString(r.domain)
	defer C.free(unsafe.Pointer(domain))
	host := C.CString(r.host)
	defer C.free(unsafe.Pointer(host))
	textRecord := makeTextRecord(textRecords)
	// var vv bytes.Buffer
	// vv.WriteByte(3)
	// vv.Write([]byte("a=b"))
	// vv.WriteString("3a=b")
	// slice := vv.Bytes()
	// fmt.Println("buffer", slice, vv.Len())
	// zz := copy(vv, textRecord)
	// vv = []byte("3a=b")
	// vv = append(vv, []byte("3a=b")...)
	// fmt.Println("data", vv, len(vv), zz, string(vv))
	// fmt.Println(r.RegistrationType, flags, r.interfaceIndex, r.Port)
	// vv := make([]byte, 0)
	// vv = append(vv, 3)
	// fmt.Println(vv)
	// vv = append(vv, []byte("a=b")...)
	var service C.DNSServiceRef
	errorCode := C.DNSServiceRegister(
		&service,
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
		// 0,
		// nil,
		C.DNSServiceRegisterReply(C.registrationCallback),
		// unsafe.Pointer(&Z{hh: "string", ary: &[]string{"one"}, ff: func(args RegistrationCallbackArgs) { fmt.Println("ff") }}),
		unsafe.Pointer(r),
		// nil,
	)

	fmt.Println("Registration done:", errorCode)

	if errorCode == C.kDNSServiceErr_NoError {
		errorCode = C.DNSServiceProcessResult(service)
		fmt.Println("process", errorCode)
	}

}

func makeTextRecord(texts []string) []byte {
	if len(texts) == 0 {
		return nil
	}
	textRecord := make([]byte, 0)
	for _, txt := range texts {
		// textRecord = append(textRecord, []byte(strconv.Itoa(len(txt)))...)
		textRecord = append(textRecord, byte(len(txt)))
		textRecord = append(textRecord, []byte(txt)...)
	}
	return textRecord
	// fmt.Println("makeText", text)
	// textRecord := make([]byte, 0, len(text))
	// fmt.Println("textRecord is", textRecord, len(textRecord), len(text))
	// zz := copy(textRecord, text)
	// fmt.Println("copy done", zz, string(textRecord), len(textRecord))

	// textRecord := C.CString(r.TextRecord)
	// textRecord = append(textRecord, 3)
	// textRecord = append(textRecord, []byte("a=b")...)
	// defer C.free(unsafe.Pointer(textRecord))
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
	registration := *(*Registration)(pRegistration)
	callbackArgs := RegistrationCallbackArgs{
		service:      service,
		errorCode:    errorCode,
		Registration: registration,
	}
	callbackArgs.FlagAnalyzer.flags = flags
	if errorCode == C.kDNSServiceErr_NoError {
		callbackArgs.name = C.GoString(name_p)
		callbackArgs.registrationType = C.GoString(registrationType_p)
		callbackArgs.domain = C.GoString(domain_p)
	}
	registration.Callback(callbackArgs)
}
