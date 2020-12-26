#!/usr/bin/env groovy

pipeline {
  agent {
    docker { image 'golang'}
  }
  environment {
    CI = 'true'
    GOCACHE = 'off'
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
