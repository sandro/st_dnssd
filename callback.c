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

void resolveReplyCallback(
  void                  *sdRef,
  uint32_t              flags,
  uint32_t              ifIndex,
  DNSServiceErrorType   errorCode,
  char *fullName,
  char *hostTarget,
  uint16_t port,
  uint16_t txtLen,
  char *txtRecord,
  void *context
) {
  goResolveReplyCallback(
    sdRef,
    flags,
    ifIndex,
    errorCode,
    fullName,
    hostTarget,
    port,
    txtLen,
    txtRecord,
    context
  );
}
