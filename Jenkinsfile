pipeline{
    agent any
    tools{
        //name of the tool of Golang, actually i use the 1.14
        go  'go1.15'
    }
    environment {
        //variables de ambiente, no las borres igual y las usa
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        // the path of go installation 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages{
        stage('Build'){
            steps{
                //clone the repo from my git
                git 'https://github.com/danielrb97/DOTT.git'
                //building the "api" file from repo in docker image <3 so anyone can use it
                sh 'cd  ${WORKSPACE}/cidr_convert_api/go/ && sudo docker build -t api .'
            }
        }
        stage('Sonar analysis'){
            steps{
                script{
                    //taking the tool of sonarscanner 4.7.0
                    def scannerHome = tool 'SonarQubeScanner-4.7.0';
                    //enviroment and version  of  mmy sonarqube 
                    withSonarQubeEnv('sonarqube-7.9.3'){
                        sh "${scannerHome}/bin/sonar-scanner \
                          -Dsonar.projectKey=Final-project1 \
                          -Dsonar.sources=cidr_convert_api/go/\
                          -Dsonar.host.url=http://ec2-3-144-188-140.us-east-2.compute.amazonaws.com:9000 \
                          -Dsonar.login=71578fd0cb7dcf663c19a6eec38f0a00e8986ba4"
                    }
                }
            }
        }
        stage('pre-test'){
            steps{
                //installation of de the packages, the go version its just to let me know which version of i using
                sh 'go version'
                sh 'go get -v -u github.com/gorilla/mux  && go get -v -u github.com/stretchr/testify/assert'
                sh 'go install github.com/gorilla/mux && go install github.com/stretchr/testify/assert'
            }          
        }
        stage('Unit test'){
            steps{
                script{ 
                    // with this, im saying to  jenkins that he has to werk in that path 
                   dir ('/var/lib/jenkins/workspace/sonarqube-pipeline/cidr_convert_api/go'){
                       sh 'go test -v'  
                   }
                }
            }
        } 
        stage('deployment'){
            steps{
                //again, copy the app .NOTE: WORKSPACE is like  a path that jenkins create
                git 'https://github.com/danielrb97/DOTT.git'
                   //      whe mi go files are                   building the image                                          putting a tag    the ports where the image is running 5001 is the best   "new sever" will be the name of the image and "go-docker" the repository                                                                                            
                sh 'cd ${WORKSPACE}/cidr_convert_api/go  && sudo docker build  --tag go-docker1 . && sudo docker tag go-docker1:latest go-docker1:v2.0.0 &&  sudo docker run -d -p 5001:5001 --name DevOps go-docker1'
            }
        }
    }
    
}