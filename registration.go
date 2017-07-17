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
	"errors"
	"fmt"
	"log"
	"unsafe"
)

type RegistrationCallbackArgs struct {
	FlagAnalyzer
	service          C.DNSServiceRef
	flags            C.DNSServiceFlags
	ErrorCode        int
	Name             string
	RegistrationType string
	Domain           string
	Registration     *Registration
}

// Registration can be used to register a service
type Registration struct {
	// Service          C.DNSServiceRef
	flags            uint32
	interfaceIndex   int
	Name             string
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

// RegisterService advertises a dnssd service
func (r *Registration) RegisterService(textRecords []string) error {
	var flags C.DNSServiceFlags

	name := C.CString(r.Name)
	defer C.free(unsafe.Pointer(name))
	registrationType := C.CString(r.RegistrationType)
	defer C.free(unsafe.Pointer(registrationType))
	domain := C.CString(r.domain)
	defer C.free(unsafe.Pointer(domain))
	host := C.CString(r.host)
	defer C.free(unsafe.Pointer(host))
	textRecord := makeTextRecord(textRecords)
	log.Println("length of textRecord", len(textRecord), C.uint16_t(len(textRecord)))
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
	var txtPtr unsafe.Pointer
	if len(textRecord) > 0 {
		txtPtr = unsafe.Pointer(&textRecord[0])
	} else {
		txtPtr = nil
	}
	fmt.Printf("%#v\n", txtPtr)
	stateIndex := callbackState.Add(r)
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
		txtPtr,
		// 0,
		// nil,
		C.DNSServiceRegisterReply(C.registrationCallback),
		// unsafe.Pointer(&Z{hh: "string", ary: &[]string{"one"}, ff: func(args RegistrationCallbackArgs) { fmt.Println("ff") }}),
		unsafe.Pointer(&stateIndex),
		// nil,
	)

	log.Println("Registration done:", errorCode)

	if errorCode == C.kDNSServiceErr_NoError {
		go func() {
			for {
				errorCode = C.DNSServiceProcessResult(service)
				fmt.Println("registration process", errorCode)
			}
		}()
		return nil
	}
	return errors.New("unknown error")

}

func makeTextRecord(texts []string) []byte {
	log.Println("makeTextRecord for", texts)
	if len(texts) == 0 {
		log.Println("noTextRecord returning nil")
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
	pstateIndex unsafe.Pointer,
) {
	stateIndex := *(*int)(pstateIndex)
	// var registration *Registration
	registration := callbackState.Get(stateIndex).(*Registration)
	callbackArgs := RegistrationCallbackArgs{
		service:      service,
		ErrorCode:    int(errorCode),
		Registration: registration,
	}
	callbackArgs.FlagAnalyzer.flags = flags
	if errorCode == C.kDNSServiceErr_NoError {
		callbackArgs.Name = C.GoString(name_p)
		callbackArgs.RegistrationType = C.GoString(registrationType_p)
		callbackArgs.Domain = C.GoString(domain_p)
	}
	registration.Callback(callbackArgs)
}
