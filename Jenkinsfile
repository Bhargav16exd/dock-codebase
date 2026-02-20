pipeline {
  agent any

  environment {
    VERSION = "build-${BUILD_NUMBER}"
  }
  tools {
    go 'Go'  
  }
  stages {
    stage('Build') {
      steps {
        sh '''
        export CGO_ENABLED=0
        export GOOS=linux
        export GOARCH=arm64
        go build -o dock
        '''
      }
    }

    stage('Package') {
      steps {
        sh 'mv dock dock-build-${VERSION}'
        archiveArtifacts artifacts: "dock-build-${VERSION}"
      }
    }
  }
}
