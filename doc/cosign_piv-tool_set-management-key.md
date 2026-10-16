## cosign piv-tool set-management-key

sets the management key of a hardware token

```
cosign piv-tool set-management-key [flags]
```

### Options

```
  -h, --help                    help for set-management-key
      --new-key string          new management key, uses default if empty
      --old-key string          existing management key, uses default if empty
      --random-management-key   if set to true, generates a new random management key and deletes it after
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

