#include "_cgo_export.h"

void registrationCallback(
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

void browseCallback(
  void                  *sdRef,
  uint32_t              flags,
  uint32_t              ifIndex,
  DNSServiceErrorType   errorCode,
  void                  *serviceName,
  void                  *regType,
  void                  *replyDomain,
  void                  *context
) {
  printf("in browse Callback, \n");
  goBrowseCallback(
    sdRef,
    flags,
    ifIndex,
    errorCode,
    serviceName,
    regType,
    replyDomain,
    context
  );
}
