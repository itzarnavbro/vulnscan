# VulnScan

A comprehensive security scanner and web dashboard for detecting vulnerabilities (CVEs), IaC misconfigurations, and exposed secrets across container images, filesystems, and Git repositories.

## Features
- **Container Image Scanning:** Inspects OS packages and language dependencies for known CVEs.
- **Filesystem & Repository Scanning:** Analyzes source code lockfiles, Dockerfiles, and Terraform/Kubernetes manifests.
- **Secret Detection:** Identifies hardcoded API keys, tokens, and credentials.
- **SBOM Generation:** Exports software inventories in CycloneDX and SPDX formats.

## Build from Source
```bash
go build -ldflags="-s -w -X vulnscan/pkg/version/app.ver=1.0.0" -o vulnscan ./cmd/vulnscan