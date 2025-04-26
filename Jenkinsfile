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
        stage('install docker') {
            steps {
                echo "Install Dcoker ..."
                sh ' apt-get update \
                    apt-get install ca-certificates curl \
                    install -m 0755 -d /etc/apt/keyrings \
                    curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc \
                    chmod a+r /etc/apt/keyrings/docker.asc'
                sh 'echo \
                    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/debian \
                    $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
                      tee /etc/apt/sources.list.d/docker.list > /dev/null'
                sh ' apt-get update && apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin'
            }
        }
        stage('Dockerfile-lint') {
            steps {
                echo "Running Dockerfile lint for environment: ${params.ENV}"
                sh 'docker run --rm -i \
                ghcr.io/hadolint/hadolint < Dockerfile'
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