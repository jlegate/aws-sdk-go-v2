// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a version of a SageMaker image.
func (c *Client) DescribeImageVersion(ctx context.Context, params *DescribeImageVersionInput, optFns ...func(*Options)) (*DescribeImageVersionOutput, error) {
	if params == nil {
		params = &DescribeImageVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImageVersion", params, optFns, c.addOperationDescribeImageVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImageVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImageVersionInput struct {

	// The name of the image.
	//
	// This member is required.
	ImageName *string

	// The version of the image. If not specified, the latest version is described.
	Version *int32

	noSmithyDocumentSerde
}

type DescribeImageVersionOutput struct {

	// The registry path of the container image on which this image version is based.
	BaseImage *string

	// The registry path of the container image that contains this image version.
	ContainerImage *string

	// When the version was created.
	CreationTime *time.Time

	// When a create or delete operation fails, the reason for the failure.
	FailureReason *string

	// The Amazon Resource Name (ARN) of the image the version is based on.
	ImageArn *string

	// The ARN of the version.
	ImageVersionArn *string

	// The status of the version.
	ImageVersionStatus types.ImageVersionStatus

	// When the version was last modified.
	LastModifiedTime *time.Time

	// The version number.
	Version *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeImageVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeImageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeImageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeImageVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImageVersion(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeImageVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeImageVersion",
	}
}
