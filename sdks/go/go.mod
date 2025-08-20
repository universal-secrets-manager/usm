module github.com/universal-secrets-manager/usm/sdks/go

go 1.21

require github.com/universal-secrets-manager/usm/core/crypto v0.0.0

require gopkg.in/yaml.v3 v3.0.1 // indirect

replace github.com/universal-secrets-manager/usm/core/crypto => ../../core/crypto
