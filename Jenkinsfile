pipeline{
    agent{
        kubernetes{
            cloud 'kubernetes'
            yaml'''
            apiVersion: v1
            kind: Pod
            spec: 
                containers:
                - name: buildah
                  image: quay.io/buildah/stable:latest
                  command:
                    - cat
                  tty: true
                  securityContext:
                    privileged: true
                - name: alpine
                  image: docker.io/library/alpine:3
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
                CREDENTIALS_REGISTRY = credentials 'jesayafn_cred-quay.io'
            }
            steps {
                container('buildah') {
                    sh '''buildah login --username ${CREDENTIALS_REGISTRY_USR} \\
                    --password ${CREDENTIALS_REGISTRY_PSW} --verbose\\
                    ${PROVIDER_REGISTRY}'''
                    sh '''buildah build --compress --file ./Dockerfile \\
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
                PROVIDER_REGISTRY = 'quay.io'
                IMAGE_REGISTRY = 'jesayafn/musics'
            }
            steps {
                container('alpine') {
                    sh 'wget https://github.com/aquasecurity/trivy/releases/download/v0.35.0/trivy_0.35.0_Linux-64bit.tar.gz'
                    sh 'tar -xvzf trivy_0.35.0_Linux-64bit.tar.gz && mv trivy /bin'
                    sh 'trivy image --no-progress --output trivy-report.out ${PROVIDER_REGISTRY}/${IMAGE_REGISTRY}:${BUILD_NUMBER}'
                }
            }
        }
    }
}