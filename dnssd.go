package main

/*
#include <stdio.h>
#include <errno.h>
#include <stdlib.h>
#include <dns_sd.h>
extern void serviceRegister();
*/
import "C"
import (
  "fmt"
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
  C.serviceRegister()
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
}

func main() {
  fmt.Println("booting...")
  register()
}
