// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/elastic/cloudbeat/resources/providers/awslib"
	"github.com/elastic/elastic-agent-libs/logp"
)

type BucketDescription struct {
	Name             string
	SSEAlgorithm     string
	BucketPolicy     BucketPolicy
	BucketVersioning BucketVersioning
}

// TODO: This can be better typed, but this is a complex object. See this library for example: https://github.com/liamg/iamgo/
type BucketPolicy map[string]any

type BucketVersioning struct {
	Enabled   bool
	MfaDelete bool
}

type Logging struct {
	Enabled      bool   `json:"Enabled"`
	TargetBucket string `json:"TargetBucket"`
}

type S3 interface {
	DescribeBuckets(ctx context.Context) ([]awslib.AwsResource, error)
	GetBucketACL(ctx context.Context, bucketName *string, region string) ([]types.Grant, error)
	GetBucketPolicy(ctx context.Context, bucketName *string, region string) (BucketPolicy, error)
	GetBucketLogging(ctx context.Context, bucketName *string, region string) (Logging, error)
}

type Provider struct {
	log     *logp.Logger
	clients map[string]Client
}

type Client interface {
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	GetBucketEncryption(ctx context.Context, params *s3.GetBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.GetBucketEncryptionOutput, error)
	GetBucketLocation(ctx context.Context, params *s3.GetBucketLocationInput, optFns ...func(*s3.Options)) (*s3.GetBucketLocationOutput, error)
	GetBucketPolicy(ctx context.Context, params *s3.GetBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.GetBucketPolicyOutput, error)
	GetBucketVersioning(ctx context.Context, params *s3.GetBucketVersioningInput, optFns ...func(*s3.Options)) (*s3.GetBucketVersioningOutput, error)
	GetBucketAcl(ctx context.Context, params *s3.GetBucketAclInput, optFns ...func(*s3.Options)) (*s3.GetBucketAclOutput, error)
	GetBucketLogging(ctx context.Context, params *s3.GetBucketLoggingInput, optFns ...func(*s3.Options)) (*s3.GetBucketLoggingOutput, error)
}