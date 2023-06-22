# Sample Go Application

This is a sample go application based on DDD.


## Start app on local without Docker

Required: PostgreSQL is running on your local machine.

```sh
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_PORT=5432
export DB_NAME=go-app-sample
make start-api
```

## Start app on local using Docker

Requirement: Docker is running on your local machine.

```sh
make run
```


// Creating ECS Cluster

```cmd
aws ec2 create-security-group --group-name example-ecs-sg --description example-ecs-sg

aws ec2 authorize-security-group-ingress --group-name example-ecs-sg --protocol tcp --port 22 --cidr 0.0.0.0/0

ecs-cli up --cluster example-ecs-cluster --instance-role ecsInstanceRole --keypair dev.pem --size 1 --security-group example-ecs-sg --subnets subnet-dafc6f93,subnet-6b8f2b30,subnet-eed8a38b,subnet-f38b2bde --vpc vpc-340b1053 --instance-type t2.small --launch-type EC2

```

// Creating Task Definition

```cmd
aws ecs register-task-definition --cli-input-json file://task-definition.json
```

```yaml
{
    "containerDefinitions": [
        {
            "name": "go-app",
            "image": "644928968953.dkr.ecr.ap-southeast-1.amazonaws.com/go-api:v5",
            "memory": 512,
            "cpu": 128,
            "portMappings": [
                {
                    "containerPort": 8080,
                    "hostPort": 8080,
                    "protocol": "tcp"
                }
            ],
            "environment": [
                {
                    "name": "DB_HOST",
                    "value": "database-1.ckxnccosaddo.ap-southeast-1.rds.amazonaws.com"
                },
                {
                    "name": "DB_NAME",
                    "value": "go-app-sample"
                },
                {
                    "name": "DB_PASSWORD",
                    "value": "postgres"
                },
                {
                    "name": "DB_USER",
                    "value": "postgres"
                },
                {
                    "name": "DB_PORT",
                    "value": "5432"
                }
            ],
            "command": [
                "./app"
            ],
            "essential": true
        }
    ],
    "family": "go-api",
    "taskRoleArn": "arn:aws:iam::644928968953:role/ecsTaskExecutionRole"
}
```

// Creating Elastic Load Balancer (ELB) with Target Group

```
aws ec2 create-security-group --group-name example-elb-sg --description example-elb-sg

aws ec2 authorize-security-group-ingress --group-name example-elb-sg --protocol tcp --port 80 --cidr 0.0.0.0/0

aws ec2 authorize-security-group-ingress --group-name example-ecs-sg --source-group example-elb-sg --protocol tcp --port 1-65535

aws elbv2 create-target-group --name example-target-group --port 80 --protocol HTTP --target-type instance --vpc-id vpc-340b1053 --health-check-protocol HTTP --health-check-path /hello-world

aws elbv2 create-load-balancer --name example-elb --subnets subnet-dafc6f93 subnet-6b8f2b30 subnet-eed8a38b subnet-f38b2bde --security-groups example-elb-sg --scheme internet-facing --type application

aws elbv2 create-listener --load-balancer-arn arn:aws:elasticloadbalancing:us-east-1:548754742764:loadbalancer/app/example-elb/3cb7c0ce850338d6 --protocol HTTP --port 80 --default-actions Type=forward,TargetGroupArn=arn:aws:elasticloadbalancing:us-east-1:548754742764:targetgroup/example-target-group/73e3f1b663022983

```

// Creating ECS Service

```
aws ecs create-service --cluster example-ecs-cluster --service-name example-ecs-service --task-definition nodejs-task-def --desired-count 1 --launch-type EC2 --load-balancers targetGroupArn=arn:aws:elasticloadbalancing:us-east-1:548754742764:targetgroup/example-target-group/73e3f1b663022983,containerName=node-app,containerPort=8080
```