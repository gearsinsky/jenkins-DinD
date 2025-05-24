@Library('jenkins-shared-lib') _

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
        DOCKER_HOST = 'tcp://docker-dind-daemon:2375'
    }

    stages {
        stage('Debug Docker') {
            steps {
                debugdocker()
            }
        }

        stage('Build') {
            steps {
                script{
                    build()
                }
            }
        }

        stage('image-lint') {
            steps {
                imagelint()
            }
        }

        stage('Debug Branch') {
            steps {
                debugbranch()
            }
        }

        stage('docker push') {
            when {
                anyOf {
                    branch 'hotfix'
                }
            }
            steps {
                dockerpush()
            }
        }
    }

    post {
        success {
            script {
                sendtelegrammsg(env.S_MESSAGE)
            }
        }
        failure {
            script {
                sendtelegrammsg(env.F_MESSAGE)
            }
        }
    }
}
