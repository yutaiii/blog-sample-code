package main

import (
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/compute/metadata"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {
	awsRoleArn := "<作成したロールのARN>"

	// メタデータサーバーから値を取得
	instanceName := getMetadata("instance", "name")
	projectId := getMetadata("project", "project-id")
	projectAndInstanceName := fmt.Sprintf("%s.%s", projectId, instanceName)
	token := getMetadata("instance", "service-accounts/default/identity?format=standard&audience=gcp")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("<リージョン>"),
	})
	if err != nil {
		fmt.Println("NewSession Error", err)
		return
	}

	// STS Clientを作成
	svc := sts.New(sess)
	result, err := svc.AssumeRoleWithWebIdentity(&sts.AssumeRoleWithWebIdentityInput{
		RoleArn:          &awsRoleArn,
		RoleSessionName:  &projectAndInstanceName,
		WebIdentityToken: &token,
	})
	if err != nil {
		fmt.Println("AssumeRoleWithWebIdentity Error", err)
		return
	}

	AccessKeyId := *result.Credentials.AccessKeyId
	SecretAccessKey := *result.Credentials.SecretAccessKey
	SessionToken := *result.Credentials.SessionToken
	Expiration := result.Credentials.Expiration

	fmt.Println(fmt.Sprintf("accesskey:%s, secretkey:%s, sessiontoken:%s, expiration:%v", AccessKeyId, SecretAccessKey, SessionToken, Expiration))

	creds := credentials.NewStaticCredentials(AccessKeyId, SecretAccessKey, SessionToken)
	sess2, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String("<リージョン>")},
	)
	if err != nil {
		fmt.Println("NewSession with credentials Error", err)
		return
	}
	client := s3.New(sess2)

	bucketName := "<バケット名>"
	objectKey := "<ファイル名>"
	obj, err := client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 内容を読み込んで表示
	byteArray, _ := ioutil.ReadAll(obj.Body)
	fmt.Println(string(byteArray))
}

func getMetadata(path, parameter string) string {
	data, _ := metadata.Get(fmt.Sprintf("%s/%s", path, parameter))
	return data
}
