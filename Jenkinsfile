#!/usr/bin/env groovy

pipeline {
  agent {
    docker { image 'golang'}
  }
  environment {
    CI = 'true'
    XDG_CACHE_HOME = '/tmp/.cache'
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
