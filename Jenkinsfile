pipeline {
    agent any

    parameters {
        choice(name: 'ENV', choices: ['dev', 'staging', 'prod'], description: 'Select deployment environment')
    }

    environment {
        PROJECT_NAME = "python-project"
        S_MESSAGE = "✅ 成功：$JOB_NAME #$BUILD_NUMBER"
        F_MESSAGE = "❌ 失敗：$JOB_NAME #$BUILD_NUMBER"
        DOCKER_REPO = "owen0522/jenkins_dind"

    }

    stages {
        stage('Build') {
            steps {
                echo "Building image: ${env.PROJECT_NAME}"
                sh """
                    docker build --platform linux/amd64 -t ${env.PROJECT_NAME}:${BUILD_NUMBER} .
                """
                echo "image already to use"
                sh 'docker images'
            }
        }

        stage('image-lint') {
            steps {
                echo "Running Image lint"
                sh """
                    docker run --rm \
                    --add-host=docker-dind-daemon:172.18.0.3 \
                    aquasec/trivy image --docker-host tcp://docker-dind-daemon:2375 ${env.PROJECT_NAME}:${BUILD_NUMBER}
                    env
                """
            }
        }
        stage('Debug Branch') {
            steps {
                echo "Current Git Branch: ${env.BRANCH_NAME}"
            }
        }
        stage('docker push') {
            when {
                anyOf {
                    // branch 'main'
                    branch 'login-error'
                    branch 'hotfix'
                }
            }
            steps {
                echo "Reading push image to docker repo"
                withCredentials([usernamePassword(credentialsId: 'docker-repo', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASWD')]) {
                    sh """
                        echo "$DOCKER_PASWD" | docker login -u "$DOCKER_USER" --password-stdin
                        docker tag ${PROJECT_NAME}:${BUILD_NUMBER} ${DOCKER_REPO}:${BUILD_NUMBER}
                        docker tag ${PROJECT_NAME}:${BUILD_NUMBER} ${DOCKER_REPO}:latest
                        docker push ${DOCKER_REPO}:${BUILD_NUMBER}
                        docker push ${DOCKER_REPO}:latest
                    """
                }
            }
        }
    }

    post {
        success {
            echo "Pipeline completed successfully for ${params.ENV}"
            sh '''
                curl -s -X POST https://api.telegram.org/bot$TOKEN/sendMessage \
                -d chat_id="$CHAT_ID" \
                -d text="$S_MESSAGE"
            '''
        }
        failure {
            echo "Pipeline failed for ${params.ENV}"
            sh '''
                curl -s -X POST https://api.telegram.org/bot$TOKEN/sendMessage \
                -d chat_id="$CHAT_ID" \
                -d text="$F_MESSAGE"
            '''
        }
    }
}