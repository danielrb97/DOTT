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
    }
    stages{
        stage('Build'){
            steps{
                //clone the repo from my git
                git 'https://github.com/danielrb97/DOTT.git'
                //building the "api" file from repo in docker <3
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
                          -Dsonar.host.url=http://ec2-18-222-250-154.us-east-2.compute.amazonaws.com:9000 \
                          -Dsonar.login=71578fd0cb7dcf663c19a6eec38f0a00e8986ba4"
                    }
                }
            }
        }
        stage('Unit test'){
            steps{
                sh 'pwd ${GOPATH}'
                echo "installing plugin/dependencies"
                sh 'go version'
                sh 'go get github.com/gorilla/mux'
                script{ 
                   dir ('/var/lib/jenkins/workspace/sonarqube-pipeline/cidr_convert_api/go'){
                        sh 'go test'
                   }
                }
            }
        }
    }
    
}