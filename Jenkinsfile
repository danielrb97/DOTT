pipeline{
    agent any
    tools{
        go  'go1.15'
    }
    environment {
        //go modules go114
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        // the path of go installation
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        GOROOT = ""
    }
    stages{
        /*
        stage('Build'){
            steps{
                //clone the repo from my git
                git 'https://github.com/danielrb97/DOTT.git'
                //building the "api" file from repo in docker <3s
                sh 'cd  ${WORKSPACE}/cidr_convert_api/go/ && sudo docker build -t api .'
            }
        }
        stage('Sonar analysis'){
            steps{
                script{
                    //taking the tool of sonarscanner d4.7.0
                    def scannerHome = tool 'SonarQubeScanner-4.7.0';
                    //enviroment of sonarqube
                    withSonarQubeEnv('sonarqube-7.9.3'){
                        sh "${scannerHome}/bin/sonar-scanner \
                          -Dsonar.projectKey=Final-project1 \
                          -Dsonar.sources=cidr_convert_api/go/\
                          -Dsonar.host.url=http://ec2-3-145-113-193.us-east-2.compute.amazonaws.com:9000 \
                          -Dsonar.login=71578fd0cb7dcf663c19a6eec38f0a00e8986ba4"
                    }
                }
            }
        }
        */
        stage('Unit test'){
            steps{
                dir('/var/lib/jenkins/tools/org.jenkinsci.plugins.golang.GolangInstallation/go1.15/src'){
                    sh 'go get -v -u github.com/gorilla/mux'
                    sh 'go install  -mod=readonly  github.com/gorilla/mux'
                    //sh 'go get -v -u github.com/pkg/errors'
                    //sh 'go get -v -u github.com/stretchr/testify/assert'
                }
                script{ 
                   dir ('/var/lib/jenkins/workspace/sonarqube-pipeline/cidr_convert_api/go'){
                       sh 'go test -v'  
                   }
                }
            }
        } 
        stage('deployment'){
            steps{
                git 'https://github.com/danielrb97/DOTT.git'
                sh 'cd ${WORKSPACE}/cidr_convert_api/go  && sudo docker build  --tag go-docker . && sudo docker tag go-docker:latest go-docker:v1.0.0 &&  sudo docker run -d -p 5001:5001 --name new-server go-docker'
            }
        }
    }
    
}