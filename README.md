# Custom application development in the RainMaker AWS account

Many times, customers have a requirement of developing some custom applications and deploying these applications in the same AWS account where RainMaker is deployed. 

One of the common use cases customers have for developing the custom applications is, to build a new API linked to an AWS Lambda function and perform read/write operations on the database tables.

> **Please note:** Deployment of these custom applications should not have any impact on the RainMaker deployment or vice versa.

<hr><br>

This document provides details on how theses applications can be developed and deployed.

While the customers can make use of the AWS Console, CLI or APIs to create AWS resources and deploy their custom code, this approach is error prone and not repeatable. This document focuses on the usage of AWS CloudFormation based YML templates, which are extensible and makes the deployment process easier and repeatable.

This example focuses on the creation of an API,  a new API, an AWS Lambda function, AWS IAM roles and policies and an AWS DynamoDb table. It uses the Go language. However, customers can use any AWS resources (e.g. AWS Lambda, AWS SQS, AWS Kinesis, AWS S3, AWS SNS, even AWS EC2 or containers) and any [supported programming language](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html) of their choice like Python, Java, Node.js, etc. 

This example creates an API Gateway endpoint, which is independent of the RainMaker API Gateway. The deployment and upgrade of the custom application too can be carried out independently.

It consists of three stacks developed using AWS SAM (Serverless Application Module), a deployment framework based on AWS CloudFormation.


1. **custom-base-api**: a base API resource
2. **custom-base**: a sample Database table
3. **custom-service**: an IAM policy, a Lambda function and an API end point

<br><hr><br>

## Steps to deploy the project:

Please refer pre-requisite.txt before proceeding further

1) Create an S3 bucket in aws: 
   
   $ AWS console -> S3 -> create bucket -> enter the bucket name -> next -> next -> untick Block public access -> agree to the acknowledgement -> create bucket. 
   
   Go to bucket -> permission -> bucket policy -> replace and add the following policy 
   
   ```
   { 
       "Version": "2012-10-17", 
       "Statement": [ 
           { 
               "Sid": "AddPerm", 
               "Effect": "Allow", 
               "Principal": "*", 
               "Action": "s3:GetObject", 
               "Resource": “arn:aws:s3:::<YOUR_BUCKET_NAME>/*" 
           } 
       ] 
   } 
   ```

2) Set gosdk path
goSDK:/rainmaker-custom-development 

3) open command line terminal and Build Project 
```
$ make all 
```

> Note: $make help will display all available commands 
 
4) Deploy Project 
```
$ make deploy S3-BUCKET=<YOUR_BUCKET_NAME> 
```

5) Done 

<br><hr><br>

## Steps to test the deployment

Copy the API path from the output section of the deployment. This will be available as the last statement in the command line terminal like:

```
Key                 TestApi                                                                                                                                      
Description         API Gateway endpoint URL                                                                                          
Value               https://<API_GATEWAY_ID>.execute-api.<REGION>.amazonaws.com/dev/{version}/user/custom
```

Run the below command in the Command line terminal using the CURL tool or from a REST client like Postman:

```
curl --request GET https://<API_GATEWAY_ID>.execute-api.<REGION>.amazonaws.com/dev/{version}/user/custom
```