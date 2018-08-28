node('docker') {
    checkout scm
    stage('Build') {
        docker.image('golang').inside {
            sh 'go version'
            sh 'ls $GOPATH/src'
            sh 'go get github.com/ffalcn85/todo-api'
            sh 'cd $GOPATH/src/github.com/ffalcn85/todo-api/cmd/server'
            sh 'go build -v'
        }
    }

}

// goTemplate {
//     def buildPath = "/home/jenkins/go/src/github.com/ffalcn85/todo-api/cmd/server"
//     dir(buildPath) {
//     container('golang-build') {
//         stage('install dependencies') {
//             sh 'go version'
//             sh 'go get -u golang.org/x/lint/golint'
//             sh 'go get -u github.com/golang/dep/cmd/dep'
//             }

//         stage('format') {
//             // point to lint script, govet
//             sh 'gofmt'
//         }

//         stage('lint') {
//             // point to lint script, govet
//             sh 'go vet'
//         }

//         stage('test') {
//             // point to test script, code coverage
//             sh 'go test'
//         }
//     }
// }