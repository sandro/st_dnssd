#include "_cgo_export.h"

void serviceRegisterCallback(
  DNSServiceRef service,
  DNSServiceFlags flags,
  DNSServiceErrorType errorCode,
  const char *name,
  const char *registrationType,
  const char *domain,
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
