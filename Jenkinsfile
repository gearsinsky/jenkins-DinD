pipeline {
    agent any

    // 參數區：選擇部署環境
    parameters {
        choice(name: 'ENV', choices: ['dev', 'staging', 'prod'], description: 'Select deployment environment')
    }

    // 全域環境變數
    environment {
        PROJECT_NAME = "python-project"
        S_MESSAGE = "✅ 成功：$JOB_NAME #$BUILD_NUMBER"
        F_MESSAGE = "❌ 失敗：$JOB_NAME #$BUILD_NUMBER"
        // DEPLOY_DIR = "/opt/${env.ENV}/deployments"
    }

    stages {
        // stage('Dockerfile-lint') {
        //     steps {
        //         echo "Running Dockerfile lint for environment: ${params.ENV}"
        //         sh '''
        //             pwd
        //             docker run --rm -i \
        //             -v Dockerfile:/workspace/Dockerfile \
        //             ghcr.io/hadolint/hadolint  \
        //             hadolint /workspace/Dockerfile 
        //         '''             
        //     }
        // }
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
            stage('docker push') {
                when {
                    anyOf {
                        branch 'pushrepo'
                        branch 'stg'
                        branch 'rel'
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
    post {
        success {
            echo "Pipeline completed successfully for ${params.ENV}"
            sh '''
                curl -s -X POST https://api.telegram.org/bot$TOKEN/sendMessage \
                -d chat_id="$CHAT_ID" \
                -d text="$S_MESSAGE
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
}