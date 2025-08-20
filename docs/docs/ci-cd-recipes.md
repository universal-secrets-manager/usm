# CI/CD Recipes

## GitHub Actions

```yaml
- uses: usm/action@v1
  with:
    profile: prod
```

## GitLab CI

```yaml
include:
  - project: universal-secrets-manager/usm
    file: usm.gitlab.yml

job:
  script:
    - usm-export.sh
```

## Jenkins

```groovy
pipeline {
    agent any
    stages {
        stage('Export Secrets') {
            steps {
                script {
                    usmExport()
                }
            }
        }
    }
}
```