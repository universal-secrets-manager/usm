# Jenkins Integration

This directory contains the Jenkins integration for USM.

## Usage

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