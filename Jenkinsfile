pipeline {
    agent any

    environment {
    DOCKER_IMAGE = 'atulsoni9/scientificcalc'
    DOCKER_TAG = "build-${env.BUILD_ID}"
	}	

    stages {
        stage('Checkout') {
            steps {
                script {
                    try {
                        checkout scm
                    } catch (Exception e) {
                        currentBuild.description = "Failed Stage: Checkout. Reason: ${e.getMessage()}"
                        throw e
                    }
                }
            }
        }

        stage('Build') {
            steps {
                script {
                    try {
                        // Compiles your scientific calculator code
                        bat 'go build -o bin/app.exe main.go'
                    } catch (Exception e) {
                        currentBuild.description = 'Failed Stage: Build. Reason: Go compilation failed.'
                        throw e
                    }
                }
            }
        }

        stage('Test') {
            steps {
                script {
                    try {
                        // Runs automated tests (ensure main_test.go exists!)
                        bat 'go test ./...'
                    } catch (Exception e) {
                        currentBuild.description = 'Failed Stage: Test. Reason: Go tests failed.'
                        throw e
                    }
                }
            }
        }

        stage('Package') {
            steps {
                script {
                    try {
                        bat 'if not exist dist mkdir dist'
                        bat 'copy bin\\app.exe dist\\'
                    } catch (Exception e) {
                        currentBuild.description = 'Failed Stage: Package'
                        throw e
                    }
                }
            }
        }

        stage('Docker Build') {
            steps {
                script {
                    try {
                        // Builds image using your multi-stage Dockerfile
                        bat "docker build -t %DOCKER_IMAGE%:%DOCKER_TAG% ."
                    } catch (Exception e) {
                        currentBuild.description = 'Failed Stage: Docker Build'
                        throw e
                    }
                }
            }
        }

        stage('Docker Push') {
            steps {
                script {
                    // This pushes your image to hub.docker.com/u/atulsoni9
                    docker.withRegistry('https://index.docker.io/v1/', 'dockerhub-credentials-id') {
                        bat "docker push %DOCKER_IMAGE%:%DOCKER_TAG%"
                        bat "docker tag %DOCKER_IMAGE%:%DOCKER_TAG% %DOCKER_IMAGE%:latest"
                        bat "docker push %DOCKER_IMAGE%:latest"
                    }
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    try {
                // Using powershell (pwsh) to manage Docker directly
                powershell '''
                    # 1. Prune old images (Resource Pruning)
                    docker image prune -f

                    # 2. Pull the latest image
                    docker pull atulsoni9/scientificcalc:latest

                    # 3. Clean up any existing container
                    $container = docker ps -aq -f name=scientific-calc-cli
                    if ($container) { 
                        docker rm -f scientific-calc-cli 
                    }

                    Write-Host "Deployment Successful! Ready to run."
                '''
            } catch (Exception e) {
                        currentBuild.description = 'Failed Stage: Deploy'
                        throw e
                    }
                }
            }
        }
    }

    post {
		success {
            mail to: 'atul9soni@gmail.com',
                 subject: "SUCCESS: Build #${env.BUILD_ID}",
                 body: "The Scientific Calculator build passed. View it here: ${env.BUILD_URL}"
        }
        failure {
            mail to: 'atul9soni@gmail.com',
                 subject: "FAILURE: Build #${env.BUILD_ID}",
                 body: "The build failed. Check the console: ${env.BUILD_URL}console"
        }
    }
}