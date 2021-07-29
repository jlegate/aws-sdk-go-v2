// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers your account as a publisher of public extensions in the CloudFormation
// registry. Public extensions are available for use by all CloudFormation users.
// This publisher ID applies to your account in all Regions. For information on
// requirements for registering as a public extension publisher, see Registering
// your account to publish CloudFormation extensions
// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html#publish-extension-prereqs)
// in the CloudFormation CLI User Guide.
func (c *Client) RegisterPublisher(ctx context.Context, params *RegisterPublisherInput, optFns ...func(*Options)) (*RegisterPublisherOutput, error) {
	if params == nil {
		params = &RegisterPublisherInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterPublisher", params, optFns, c.addOperationRegisterPublisherMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterPublisherOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterPublisherInput struct {

	// Whether you accept the Terms and Conditions
	// (https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf)
	// for publishing extensions in the CloudFormation registry. You must accept the
	// terms and conditions in order to register to publish public extensions to the
	// CloudFormation registry. The default is false.
	AcceptTermsAndConditions *bool

	// If you are using a Bitbucket or GitHub account for identity verification, the
	// Amazon Resource Name (ARN) for your connection to that account. For more
	// information, see Registering your account to publish CloudFormation extensions
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html#publish-extension-prereqs)
	// in the CloudFormation CLI User Guide.
	ConnectionArn *string

	noSmithyDocumentSerde
}

type RegisterPublisherOutput struct {

	// The ID assigned this account by CloudFormation for publishing extensions.
	PublisherId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterPublisherMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRegisterPublisher{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRegisterPublisher{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterPublisher(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterPublisher(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "RegisterPublisher",
	}
}
