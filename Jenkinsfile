#!/usr/bin/env groovy

pipeline {
  agent any
  environment {
    CI = 'true'
    XDG_CACHE_HOME = '/tmp/.cache'
  }
  stages {
    agent {
        docker { image 'golang'}
      }
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
    stage ('BuildAndPublish') {
      agent any
      steps {
        sh 'make image'
        sh 'make push'
      }
    }
  }
}
