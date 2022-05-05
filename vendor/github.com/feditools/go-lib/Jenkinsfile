pipeline {
  environment {
    BUILD_IMAGE = 'gobuild:1.17'
    PATH = '/go/bin:~/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/usr/local/go/bin'
  }

  agent any

  stages {

    stage('Check Formatting') {
      agent {
        docker {
          image "${BUILD_IMAGE}"
          args '-e HOME=${WORKSPACE}'
          reuseNode true
        }
      }
      steps {
        script {
          sh "make check"
        }
      }
    }

    stage('Test') {
      agent {
        docker {
          image "${BUILD_IMAGE}"
          args '-e HOME=${WORKSPACE}'
          reuseNode true
        }
      }
      steps {
        script {
          withCredentials([
            string(credentialsId: 'codecov-feditools-go-lib', variable: 'CODECOV_TOKEN')
          ]) {
            sh """#!/bin/bash
            go test -race -coverprofile=coverage.txt -covermode=atomic ./...
            RESULT=\$?
            bash <(curl -s https://codecov.io/bash)
            exit \$RESULT
            """
          }
        }
      }
    }

  }

}
