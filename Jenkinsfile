pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                echo 'go version'
                echo 'go build src/main.go'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
                echo 'go test -v src/*.go'
            }
        }
        stage('Docker') {
            steps {
                echo 'Building and Uploading Docker Image..'
                echo 'docker build -t sergividal/go-app .'
                echo 'docker push sergividal/go-app'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                echo 'kubectl create -f deployment2.yaml'
                echo 'kubectl patch service go-app -p \'{"spec":{"selector":{"version": "2.0.0"}}}\''
            }
        }
    }
}