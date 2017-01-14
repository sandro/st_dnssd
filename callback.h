extern void registrationCallback(
  DNSServiceRef service,
  DNSServiceFlags flags,
  DNSServiceErrorType errorCode,
  char *name,
  char *registrationType,
  char *domain,
  void *context
);

extern void browseCallback(
  void                  *sdRef,
  uint32_t              flags,
  uint32_t              ifIndex,
  DNSServiceErrorType   errorCode,
  void                  *serviceName,
  void                  *regType,
  void                  *replyDomain,
  void                  *context
);
