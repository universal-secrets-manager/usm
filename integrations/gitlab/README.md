# GitLab CI Integration

This directory contains the GitLab CI integration for USM.

## Usage

```yaml
include:
  - project: universal-secrets-manager/usm
    file: usm.gitlab.yml

job:
  script:
    - usm-export.sh
```