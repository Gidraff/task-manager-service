#!/usr/bin/env groovy

pipeline {

  environment {
    CI = 'true'
    XDG_CACHE_HOME = '/tmp/.cache'
    IMAGENAME = "gidraff/taskman"
    REGISTRYCREDENTIAL = 'yenigul-dockerhub'
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
    stage ('BuildImage') {
      agent any
      steps {
        script {
            dockerImage = docker.build IMAGENAME
        }
      }
    }
    stage ('PublishImage') {
      agent any
      steps {
        script {
            docker.withRegistry( '', REGISTRYCREDENTIAL ) {
                dockerImage.push("$BUILD_NUMBER")
                dockerImage.push('latest')
            }
        }
      }
    }
  }
}
