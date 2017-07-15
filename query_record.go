package st_dnssd

/*
#include <stdio.h>
#include <errno.h>
#include <stdlib.h>
#include <dns_sd.h>
#include "callback.h"
*/
import "C"

type QueryRecord struct {
}

type QueryRecordReplyArgs struct {
	FlagAnalyzer
	service              C.DNSServiceRef
	ifIndex              C.uint32_t
	errorCode            C.DNSServiceErrorType
	fullName             string
	resourceRecordType   uint16
	resourceRecordClass  uint16
	resourceRecordLength uint16
	resourceRecordData   string
	ttl                  uint32
	QueryRecord          QueryRecord
}
