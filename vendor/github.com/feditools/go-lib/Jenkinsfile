pipeline {
  environment {
    PATH = '/go/bin:~/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/usr/local/go/bin'
  }

  agent any

  stages {

    stage('Test') {
      agent {
        docker {
          image 'gobuild:1.18'
          args '-e HOME=${WORKSPACE} -v /var/lib/jenkins/go:/go'
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
