extern void serviceRegisterCallback(
  DNSServiceRef service,
  DNSServiceFlags flags,
  DNSServiceErrorType errorCode,
  char *name,
  char *registrationType,
  char *domain,
  void *context
);

extern DNSServiceRegisterReply serviceRegisterCallbackShim();

extern DNSServiceBrowseReply bCallback(
  void                  *sdRef,
  uint32_t              flags,
  uint32_t              ifIndex,
  DNSServiceErrorType   errorCode,
  void                  *serviceName,
  void                  *regtype,
  void                  *replyDomain,
  void                  *context
);

/* DNSServiceBrowseReply BrowseCallback; */

extern void MyClBk();
