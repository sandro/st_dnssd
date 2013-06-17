#include "_cgo_export.h"

void serviceRegisterCallback() {
  printf("in c callback\n");
  goRegistrationCallback();
}

void serviceRegister() {
  printf("regisetr...\n");
  DNSServiceRef service;
  DNSServiceFlags flags = 0;
  uint32_t interfaceIndex = 0;
  char *name = NULL;
  char *registrationType = "_go._tcp";
  char *domain = NULL;
  char *host = NULL;
  uint16_t port = 56565;
  uint16_t textLength = 0;
  void *textRecord = NULL;
  /* DNSServiceRegisterReply callback; */
  void *context = NULL;
  DNSServiceErrorType registerError;
  registerError = DNSServiceRegister(&service, flags, interfaceIndex, name, registrationType, domain, host, port, textLength, textRecord, serviceRegisterCallback, context);
  printf("%s %d err\n", registrationType, registerError);
  if (registerError == kDNSServiceErr_NoError) {
    printf("no error\n");
    int err = DNSServiceRefSockFD(service);
    printf("%d error\n", err);
    registerError = DNSServiceProcessResult(service);
    printf("return %d\n", registerError);
  }
}
