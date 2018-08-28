def buildPath = "/go/src/github.com/ffalcn85/todo-api"
dir(buildPath) {
    node('docker') {
        checkout scm
        stage('Build') {
            docker.image('golang').inside {
                sh 'go version'
                sh 'ls'
                sh 'cd cmd/server && go build -v'
            }
    }
}