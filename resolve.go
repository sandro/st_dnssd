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
	"bytes"
	"fmt"
	"io"
	"unsafe"
)

type ResolveReplyArgs struct {
	FlagAnalyzer
	IfIndex       C.uint32_t
	ErrorCode     C.DNSServiceErrorType
	FullName      string
	HostTarget    string
	Port          uint16
	TextLength    uint16
	TextRecord    string
	TextRecordMap map[string]string
}

func Resolve(ifIndex uint32, name, regType, domain string, callback func(ResolveReplyArgs)) {
	var service C.DNSServiceRef
	var flags C.DNSServiceFlags = 0
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cRegType := C.CString(regType)
	defer C.free(unsafe.Pointer(cRegType))
	cDomain := C.CString(domain)
	defer C.free(unsafe.Pointer(cDomain))
	errorCode := C.DNSServiceResolve(
		&service,
		flags,
		C.uint32_t(ifIndex),
		cName,
		cRegType,
		cDomain,
		C.DNSServiceResolveReply(C.resolveReplyCallback),
		unsafe.Pointer(&callback),
	)
	if errorCode == C.kDNSServiceErr_NoError {
		go func() {
			for {
				errorCode = C.DNSServiceProcessResult(service)
				fmt.Println("resolveReply", errorCode)
			}
		}()
	}
}

func parseTextRecord(textLength C.uint16_t, textRecord *C.char) map[string]string {
	pairs := make(map[string]string, 0)
	if textLength == 0 {
		return pairs
	}
	stream := C.GoBytes(unsafe.Pointer(textRecord), C.int(textLength))
	buffer := bytes.NewBuffer(stream)
	for {
		len, err := buffer.ReadByte()
		if err == io.EOF {
			break
		}
		data := buffer.Next(int(len))
		fmt.Println(data)
		inter := bytes.Split(data, []byte("="))
		pairs[string(inter[0])] = string(inter[1])
	}
	return pairs
}

//export goResolveReplyCallback
func goResolveReplyCallback(
	service C.DNSServiceRef,
	flags C.DNSServiceFlags,
	ifIndex C.uint32_t,
	errorCode C.DNSServiceErrorType,
	fullName *C.char,
	hostTarget *C.char,
	port C.uint16_t,
	textLength C.uint16_t,
	textRecord *C.char,
	pCallback unsafe.Pointer,
) {
	pairs := parseTextRecord(textLength, textRecord)
	fmt.Println("pairs is", pairs)
	args := ResolveReplyArgs{
		FullName:      C.GoString(fullName),
		HostTarget:    C.GoString(hostTarget),
		Port:          uint16(port),
		TextLength:    uint16(textLength),
		TextRecord:    C.GoString(textRecord),
		TextRecordMap: pairs,
	}
	fmt.Println("resolve reply flags", flags)
	args.FlagAnalyzer.flags = flags
	callback := *(*func(ResolveReplyArgs))(pCallback)
	callback(args)
}
