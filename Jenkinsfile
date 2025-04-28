pipeline {
    agent any

    // 參數區：選擇部署環境
    parameters {
        choice(name: 'ENV', choices: ['dev', 'staging', 'prod'], description: 'Select deployment environment')
    }

    // 全域環境變數
    environment {
        PROJECT_NAME = "python-project"
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
                        -e DOCKER_HOST=tcp://docker-dind-daemon:2375 \
                        aquasec/trivy image --docker-host tcp://docker-dind-daemon:2375 ${env.PROJECT_NAME}:${BUILD_NUMBER}
                    """
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