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
                sh 'go version'
                sh 'go build'
            }
        }
    }
}