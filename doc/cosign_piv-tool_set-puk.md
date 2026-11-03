## cosign piv-tool set-puk

sets the PUK on a hardware token

```
cosign piv-tool set-puk [flags]
```

### Options

```
  -h, --help             help for set-puk
      --new-puk string   new PUK, uses default if empty
      --old-puk string   existing PUK, uses default if empty
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

* [cosign piv-tool](cosign_piv-tool.md)	 - Provides utilities for managing a hardware token

