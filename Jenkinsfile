pipeline{
    agent{
        kubernetes{
            cloud 'kubernetes'
            yaml'''
            apiVersion: v1
            kind: Pod
            spec: 
                container:
                - name: buildah
                  image: quay.io/buildah/stable:latest
                  command:
                    - cat
                  tty: true
                  securityContext:
                    privileged: true
                - name: trivy
                  image: docker.io/aquasec/trivy:0.34.0
                  command:
                    - cat
                  tty: true
            '''
        }
    }
    stages{
        stage('Build'){
            environment {
                PROVIDER_REGISTRY = 'quay.io'
                IMAGE_REGISTRY = 'jesayafn/musics'
                CREDENTIALS_REGISTRY = credentials 'jesayafn@quay.io'
            }
            steps {
                container('buildah') {
                    sh '''buildah login --username ${CREDENTIALS_REGISTRY_USR} \\
                    --password ${CREDENTIALS_REGISTRY_PSW} --verbose\\
                    ${PROVIDER_REGISTRY}'''
                    sh '''buildah build \\
                    --compress --file ./Dockerfile \\
                    --tag ${PROVIDER_REGISTRY}/${IMAGE_REGISTRY}:${BUILD_NUMBER} \\
                    --tag ${PROVIDER_REGISTRY}/${IMAGE_REGISTRY}:latest'''
                    sh 'buildah push ${PROVIDER_REGISTRY}/${IMAGE_REGISTRY}:${BUILD_NUMBER}'
                    sh 'buildah push ${PROVIDER_REGISTRY}/${IMAGE_REGISTRY}:latest'
                }
            }
            post {
                always {
                    container('buildah') {
                        sh 'buildah logout ${PROVIDER_REGISTRY}'
                    }
                }
            }
        }
        stage('Scan image') {
            environment {
                PROVIDER_REGISTRY = quay.io
                IMAGE_REGISTRY = jesayafn/musics
            }
            steps {
                container(trivy) {
                    sh 'image ${PROVIDER_REGISTRY}/${IMAGE_REGISTRY}:${BUILD_NUMBER}'
                }
            }
        }
    }
}