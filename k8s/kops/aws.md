

## IAM

aws iam create-group --group-name kops

aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonEC2FullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonRoute53FullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonS3FullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/IAMFullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonVPCFullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonSQSFullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonEventBridgeFullAccess --group-name kops

aws iam create-user --user-name kops

aws iam add-user-to-group --user-name kops --group-name kops

aws iam create-access-key --user-name kops

export AWS_ACCESS_KEY_ID=$(aws configure get aws_access_key_id)
export AWS_SECRET_ACCESS_KEY=$(aws configure get aws_secret_access_key)

aws s3api create-bucket \
    --bucket zirain-kops-test-state-store \
    --region us-east-1

aws s3api create-bucket \
--bucket zirain-kops-test-state-store \
--region us-east-1 \
--acl public-read

kops create cluster kops-test.k8s.local \
    --zones=us-east-1a \
    --cloud=aws \
    --networking=amazonvpc \
    --state=s3://zirain-kops-test-state-store \
    --node-count=2 \
    --master-count=1 


kops delete cluster --name kops-test.k8s.local --yes --state=s3://zirain-kops-test-state-store

kops update cluster --name kops-test.k8s.local --yes --admin --state=s3://zirain-kops-test-state-store