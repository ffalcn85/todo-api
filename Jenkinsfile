node('docker') {
    checkout scm
    stage('Build') {
        docker.image('golang:1.10-alpine3.7').inside {
            sh 'go version'
            sh 'go build -v /usr/src/github.com/ffalcn85/todo-api'
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