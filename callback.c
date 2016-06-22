#include "_cgo_export.h"

void serviceRegisterCallback(
  DNSServiceRef service,
  DNSServiceFlags flags,
  DNSServiceErrorType errorCode,
  char *name,
  char *registrationType,
  char *domain,
  void *context
) {
  printf("in register callback, %p\n", &context);
  goRegistrationCallback(
    service,
    flags,
    errorCode,
    name,
    registrationType,
    domain,
    context
  );
}

DNSServiceRegisterReply serviceRegisterCallbackShim() {
  puts("in shim callback");
  return serviceRegisterCallback;
}

void MyClBk() {
  printf("in C callback, \n");
}

DNSServiceBrowseReply bCallback(
  void                  *sdRef,
  uint32_t              flags,
  uint32_t              ifIndex,
  DNSServiceErrorType   errorCode,
  void                  *serviceName,
  void                  *regtype,
  void                  *replyDomain,
  void                  *context
) {
  printf("in bCallback, \n");
  GoBrowseCallback(context);
}

/* DNSServiceBrowseReply BrowseCallback = (DNSServiceBrowseReply) bCallback; */
