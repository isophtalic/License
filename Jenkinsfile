pipeline {
  environment {
    registryUrl = "https://hub.cyradar.com"
    ImageName = "hub.cyradar.com/license/backend"
    registryCredential = "jenkins"
    version = "$GIT_COMMIT".substring(0,8)
    DOCKER_BUILDKIT='1'
    // GITLAB_TOKEN = credentials('gitlab-token')
  }
  agent any
  stages {
    stage ('Building image') {
        steps {
            script{
            if (env.BRANCH_NAME == "dev" || env.BRANCH_NAME == "feature/cicd" || env.BRANCH_NAME == "master"  || env.BRANCH_NAME == "staging"){
                dockerImage = docker.build("$ImageName" + ":" + "$version", "-f build/Dockerfile .")
                }
            }
        }
    }
    stage('Deploy Image') {
        steps{
            script {
                if (env.BRANCH_NAME == "dev" || env.BRANCH_NAME == "staging") {
                    docker.withRegistry(registryUrl, registryCredential ) {
                        dockerImage.push(env.BRANCH_NAME)
                    }
                }
                else if (env.BRANCH_NAME == "feature/cicd") {
                    docker.withRegistry(registryUrl, registryCredential ) {
                        dockerImage.push("cicd")
                    }
                }
                else if (env.BRANCH_NAME == "master") {
                    docker.withRegistry(registryUrl, registryCredential ){
                        dockerImage.push(env.BRANCH_NAME + "." + version)
                        dockerImage.push("latest")
                    }
                }
            }
        }
    }
    stage('Remove docker image') {
      steps {
        script {
            if (env.BRANCH_NAME == "dev" || env.BRANCH_NAME == "feature/cicd" || env.BRANCH_NAME == "master"  || env.BRANCH_NAME == "staging"){
                sh "docker rmi $ImageName:$version"
            }
            if (env.BRANCH_NAME == "master"){
                sh "docker rmi $ImageName:master.$version"
            }
        }
      }
    }
  }
}