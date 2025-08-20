module github.com/universal-secrets-manager/usm/cli/usm

go 1.22

toolchain go1.24.5

require (
	github.com/spf13/cobra v1.7.0
	github.com/universal-secrets-manager/usm/core/crypto v0.0.0
)

require (
	github.com/aws/aws-sdk-go-v2 v1.38.0 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.31.1 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.18.5 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.38.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.28.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.33.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.37.1 // indirect
	github.com/aws/smithy-go v1.22.5 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.12.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/universal-secrets-manager/usm/core/crypto => ../../core/crypto
