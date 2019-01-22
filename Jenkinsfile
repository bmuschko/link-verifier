pipeline {
    agent any
    tools {
        go 'go-1.11'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
        stage('Test') {
            environment {
                CODECOV_TOKEN = credentials('CODECOV_TOKEN')
            }
            steps {
                sh 'go test ./... -coverprofile=coverage.txt'
                sh "curl -s https://codecov.io/bash"
            }
        }
    }
}