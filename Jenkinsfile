#!/usr/bin/env groovy

pipeline {

  environment {
    CI = 'true'
    XDG_CACHE_HOME = '/tmp/.cache'
    IMAGENAME = "gidraff/taskman"
    REGISTRYCREDENTIALS = 'yenigul-dockerhub'
  }
  agent any
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
      steps {
        script {
            dockerImage = docker.build IMAGENAME
        }
      }
    }
  }
}
