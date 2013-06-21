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

void serviceRegister(
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
