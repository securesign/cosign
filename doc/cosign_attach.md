## cosign attach

Provides utilities for attaching artifacts to other artifacts in a registry

### Options

```
  -h, --help   help for attach
```

### Options inherited from parent commands

```
      --output-file string        log output to a file
  -t, --timeout duration          timeout for commands (default 3m0s)
      --tracing-enabled           enable OpenTelemetry tracing
      --tracing-endpoint string   OTLP gRPC collector endpoint (e.g. host:4317)
      --tracing-insecure          use insecure gRPC connection for tracing (default true)
      --tracing-stdout            export traces to stdout instead of OTLP
  -d, --verbose                   log debug output
```

### SEE ALSO

* [cosign](cosign.md)	 - A tool for Container Signing, Verification and Storage in an OCI registry.
* [cosign attach attestation](cosign_attach_attestation.md)	 - Attach attestation to the supplied container image
* [cosign attach sbom](cosign_attach_sbom.md)	 - DEPRECATED: Attach sbom to the supplied container image
* [cosign attach signature](cosign_attach_signature.md)	 - Attach signatures to the supplied container image

