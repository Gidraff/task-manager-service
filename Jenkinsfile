#!/usr/bin/env groovy

pipeline {
  agent any
  environment {
    CI = 'true'
    XDG_CACHE_HOME = '/tmp/.cache'
    imagename = "gidraff/taskman"
    registryCredential = 'yenigul-dockerhub'
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
      steps {
        script {
            dockerImage = docker.build imagename
        }
        script {
          docker.withRegistry( '', registryCredential ) {
            dockerImage.push("$BUILD_NUMBER")
             dockerImage.push('latest')

          }
        }
      }
    }
  }
}
