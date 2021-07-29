// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the Shield Response Team's (SRT) access to your Amazon Web Services
// account. To make a DisassociateDRTRole request, you must be subscribed to the
// Business Support plan (https://aws.amazon.com/premiumsupport/business-support/)
// or the Enterprise Support plan
// (https://aws.amazon.com/premiumsupport/enterprise-support/). However, if you are
// not subscribed to one of these support plans, but had been previously and had
// granted the SRT access to your account, you can submit a DisassociateDRTRole
// request to remove this access.
func (c *Client) DisassociateDRTRole(ctx context.Context, params *DisassociateDRTRoleInput, optFns ...func(*Options)) (*DisassociateDRTRoleOutput, error) {
	if params == nil {
		params = &DisassociateDRTRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateDRTRole", params, optFns, c.addOperationDisassociateDRTRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateDRTRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateDRTRoleInput struct {
	noSmithyDocumentSerde
}

type DisassociateDRTRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateDRTRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateDRTRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateDRTRole{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateDRTRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateDRTRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "DisassociateDRTRole",
	}
}
