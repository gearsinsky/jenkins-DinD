pipeline {
    agent any

    // 參數區：選擇部署環境
    parameters {
        choice(name: 'ENV', choices: ['dev', 'staging', 'prod'], description: 'Select deployment environment')
    }

    // 全域環境變數
    environment {
        PROJECT_NAME = "go-project"
        // DEPLOY_DIR = "/opt/${env.ENV}/deployments"
    }

    stages {
        stage('Dockerfile-lint') {
            steps {
                echo "Running Dockerfile lint for environment: ${params.ENV}"
                sh '''
                    chmod 755 *
                    ls -la
                    whoami
                    docker info
                    docker run --rm -i \
                    -v $(pwd):/workspace \
                    ghcr.io/hadolint/hadolint \
                    hadolint /workspace/Dockerfile | tee hadolint.out
                '''             
            }
        }
        stage('Build') {
            steps {
                echo "Building image: ${env.PROJECT_NAME}"
                sh 'docker build --platform linux/amd64 -t ${env.PROJECT_NAME}:${BUILD_NUMBER} .'
                echo "image already to use"
                sh 'docker images'
            }
        }

        stage('image-lint') {
            steps {
                echo "Running Image lint"
                sh 'docker run --rm \
                -v /var/run/docker.sock:/var/run/docker.sock \
                aquasec/trivy image ${env.PROJECT_NAME}:${BUILD_NUMBER} '
            }
        }
    }
    post {
        success {
            echo "Pipeline completed successfully for ${params.ENV}"
        }
        failure {
            echo "Pipeline failed for ${params.ENV}"
        }
    }
}