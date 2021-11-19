# Custom application development in AWS accounts

Many times customers have a requirement of developing some custom applications and deploying these applications in the same AWS account where RainMaker is deployed.

One of the important requirement for such applications is, the deployment of these custom applications should not have any impact on RainMaker deployment and vice versa.

This document provides the details about how theses applications can be developed and deployed.

While the customers can make use of AWS Console, CLI  or APIs and create AWS resources and deploy their custom code, this approach is error prone and not repeatable.

This document focuses on the usage of CloudFormation based YML templates, which are extensible and makes the deployment process easier and repeatable.

One of the common use case customers have for developing the custom applications is, to build a new API, which will be linked to a lambda function and perform read/write operations on the database tables.

While the customers can make use of any of the AWS resources( e.g. Lambda, SQS, Kinesis, S3, SNS, even EC2 or containers),  this example focuses on how a new API can be created and linked to a Lambda function and creation of IAM Role/Policy and a DynamoDb table. 

While this example uses the Go language, the customers can choose any programming language of their choice like Python, Java, Node.js, etc.

This example creates an API Gateway endpoint, which is completely different from the RainMaker API Gateway. The deployment and upgrade of the custom application or RainMaker can be carried out independently.

This examples consists of three stacks developed using Serverless Application Module (SAM), which is a deployment framework based on AWS CloudFormation. 


*1. A SAM template which will create a Base API resource ( custom-base-api)
2. A SAM template which will create a sample Database table ( custom-base)
3. A SAM template which will create a IAM policy, Lambda function and an API end point (custom-service)*


##Steps to deploy the project:

Please refer pre-requisite.txt before proceeding further

1) Create s3 bucket in aws: 
   
   $ Aws console -> s3 -> create bucket -> enter bucket name -> next -> next -> untick Block public access -> agree to the acknowledgement -> create bucket. 
   
   Go to bucket -> permission -> bucket policy -> replace and add following policy 
   
   { 
       "Version": "2012-10-17", 
       "Statement": [ 
           { 
               "Sid": "AddPerm", 
               "Effect": "Allow", 
               "Principal": "*", 
               "Action": "s3:GetObject", 
               "Resource": “arn:aws:s3:::<Your Bucket Name>/*" 
           } 
       ] 
   } 

2) Set gosdk path
goSDK:/rainmaker-custom-development 

3) open command line terminal and Build Project 
$ make all 
Note: $make help will guide you other commands available. 
 
4) Deploy Project 
 $ make deploy S3-BUCKET=< YOUR_BUCKET_NAME> 
 
DONE!!! 


##Steps to test the deployment

Copy API path from the output section of the deployment[Will be available as last statement in command line prompt] 
Something similar as below :

Key                 TestApi                                                                                                                                      
Description         API Gateway endpoint URL                                                                                          
Value               https://<API_GATEWAY_ID>.execute-api.<REGION>.amazonaws.com/dev/{version}/user/custom

Run the below command in the Command line terminal using the CURL tool or from the REST client like Postman.

curl --request GET https://<API_GATEWAY_ID>.execute-api.<REGION>.amazonaws.com/dev/{version}/user/custom
