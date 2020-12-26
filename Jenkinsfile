#!/usr/bin/env groovy

pipeline {
  agent {
    docker {
      image 'golang:1.14.2-alpine3.11'
    }
  }
  environment {
    CI = 'true'
  }
  stages {
    stage ('Build') {
      steps {
        sh 'go build ./...'
      }
    }
    stage ('Test') {
      steps {
        sh 'go test ./...'
      }
    }
  }
}
