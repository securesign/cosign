## cosign pkcs11-tool list-keys-uris

list-keys-uris lists URIs of all keys in a PKCS11 token

```
cosign pkcs11-tool list-keys-uris [flags]
```

### Options

```
  -h, --help                 help for list-keys-uris
      --module-path string   absolute path to the PKCS11 module
      --pin string           pin of the PKCS11 slot, uses environment variable COSIGN_PKCS11_PIN if empty
      --slot-id uint         id of the PKCS11 slot, uses 0 if empty
```

### Options inherited from parent commands

```
  -f, --no-input                  skip warnings and confirmations
      --output-file string        log output to a file
  -t, --timeout duration          timeout for commands (default 3m0s)
      --tracing-enabled           enable OpenTelemetry tracing
      --tracing-endpoint string   OTLP gRPC collector endpoint (e.g. host:4317)
      --tracing-insecure          use insecure gRPC connection for tracing (default true)
      --tracing-stdout            export traces to stdout instead of OTLP
  -d, --verbose                   log debug output
```

### SEE ALSO

* [cosign pkcs11-tool](cosign_pkcs11-tool.md)	 - Provides utilities for retrieving information from a PKCS11 token.

