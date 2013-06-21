package main

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
  service Service
  flags int
  interfaceIndex int
  name string
  registrationType string
  domain string
  host string
  port int
  textRecord string
}

func (r *Registration) textRecordLength() int {
  return len(r.textRecord)
}

func (r *Registration) aCallback() {
  fmt.Println("GO A CALLBACK");
}

func (r *Registration) registerService() {
  /* var service *C.DNSServiceRef */
  var flags C.DNSServiceFlags = 0

  name := C.CString(r.name)
  registrationType := C.CString(r.registrationType)
  domain := C.CString(r.domain)
  host := C.CString(r.host)
  textRecord := C.CString(r.textRecord)
  defer C.free(unsafe.Pointer(name))
  defer C.free(unsafe.Pointer(registrationType))
  defer C.free(unsafe.Pointer(domain))
  defer C.free(unsafe.Pointer(host))
  defer C.free(unsafe.Pointer(textRecord))

  fmt.Println(r.registrationType, flags, r.interfaceIndex, r.port);
  C.serviceRegister(
    C.uint32_t(r.interfaceIndex),
    name,
    registrationType,
    domain,
    host,
    C.uint16_t(r.port),
    C.uint16_t(r.textRecordLength()),
    textRecord,
    nil,
  )
  /* errorCode := C.DNSServiceRegister( */
  /*   service, */
  /*   flags, */
  /*   C.uint32_t(0), */
  /*   name, */
  /*   registrationType, */
  /*   domain, */
  /*   host, */
  /*   C.uint16_t(56565), */
  /*   C.uint16_t(0), */
  /*   nil, */
  /*   C.serviceRegisterCallbackShim(), */
  /*   nil, */
  /* ) */

  /* fmt.Println("Registration done:", errorCode); */
}

//export goRegistrationCallback
func goRegistrationCallback() {
  fmt.Println("GO REGISTER CALLBACK");
}

func register() {
  fmt.Println("registering")
  r := Registration{registrationType: "_go._tcp", port: 56565}
  r.registerService()

  fmt.Println("length", r.textRecordLength())
  select {}
}

func main() {
  fmt.Println("booting...")
  register()
}
