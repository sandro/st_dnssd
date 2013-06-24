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
  puts("in register callback");
}

DNSServiceRegisterReply serviceRegisterCallbackShim() {
  puts("in shim callback");
  return serviceRegisterCallback;
}
