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
    stage ('Build') {
      agent { docker { image 'golang'} }
      steps {
        sh 'go build ./...'
      }
    }
    stage ('Test') {
      agent { docker { image 'golang'} }
      steps {
        sh 'go test ./...'
      }
    }
    stage ('BuildAndPublish') {
      agent any
      steps {
        script {
            dockerImage = docker.build IMAGENAME
        }
      }
    }
  }
}
