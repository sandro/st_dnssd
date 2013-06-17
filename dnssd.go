package main

/*
#include <stdio.h>
#include <errno.h>
#include <stdlib.h>
#include <dns_sd.h>
extern void serviceRegister(
    uint32_t interfaceIndex,
    char *name,
    char *registrationType,
    char *domain,
    char *host,
    uint16_t port,
    uint16_t textLength,
    char *textRecord,
    void *context
  );
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
  name := C.CString(r.name)
  registrationType := C.CString(r.name)
  name := C.CString(r.name)
  name := C.CString(r.name)
  defer C.free(unsafe.Pointer(name))
  C.serviceRegister(
    C.uint32_t(r.interfaceIndex),
    name,
  )
}

//export goRegistrationCallback
func goRegistrationCallback() {
  fmt.Println("GO REGISTER CALLBACK");
}

func register() {
  fmt.Println("registering")
  r := Registration{textRecord: "hi"}
  r.registerService()

  fmt.Println("length", r.textRecordLength())
  select {}
}

func main() {
  fmt.Println("booting...")
  register()
}
