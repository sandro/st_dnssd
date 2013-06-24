extern void serviceRegisterCallback(
  DNSServiceRef service,
  DNSServiceFlags flags,
  DNSServiceErrorType errorCode,
  const char *name,
  const char *registrationType,
  const char *domain,
  void *context
);

extern DNSServiceRegisterReply serviceRegisterCallbackShim();
