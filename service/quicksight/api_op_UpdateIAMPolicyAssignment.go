// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing IAM policy assignment. This operation updates only the
// optional parameter or parameters that are specified in the request. This
// overwrites all of the users included in Identities.
func (c *Client) UpdateIAMPolicyAssignment(ctx context.Context, params *UpdateIAMPolicyAssignmentInput, optFns ...func(*Options)) (*UpdateIAMPolicyAssignmentOutput, error) {
	if params == nil {
		params = &UpdateIAMPolicyAssignmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIAMPolicyAssignment", params, optFns, c.addOperationUpdateIAMPolicyAssignmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIAMPolicyAssignmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIAMPolicyAssignmentInput struct {

	// The name of the assignment, also called a rule. This name must be unique within
	// an Amazon Web Services account;.
	//
	// This member is required.
	AssignmentName *string

	// The ID of the Amazon Web Services account; that contains the IAM policy
	// assignment.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace of the assignment.
	//
	// This member is required.
	Namespace *string

	// The status of the assignment. Possible values are as follows:
	//
	// * ENABLED -
	// Anything specified in this assignment is used when creating the data source.
	//
	// *
	// DISABLED - This assignment isn't used when creating the data source.
	//
	// * DRAFT -
	// This assignment is an unfinished draft and isn't used when creating the data
	// source.
	AssignmentStatus types.AssignmentStatus

	// The QuickSight users, groups, or both that you want to assign the policy to.
	Identities map[string][]string

	// The ARN for the IAM policy to apply to the QuickSight users and groups specified
	// in this assignment.
	PolicyArn *string

	noSmithyDocumentSerde
}

type UpdateIAMPolicyAssignmentOutput struct {

	// The ID of the assignment.
	AssignmentId *string

	// The name of the assignment or rule.
	AssignmentName *string

	// The status of the assignment. Possible values are as follows:
	//
	// * ENABLED -
	// Anything specified in this assignment is used when creating the data source.
	//
	// *
	// DISABLED - This assignment isn't used when creating the data source.
	//
	// * DRAFT -
	// This assignment is an unfinished draft and isn't used when creating the data
	// source.
	AssignmentStatus types.AssignmentStatus

	// The QuickSight users, groups, or both that the IAM policy is assigned to.
	Identities map[string][]string

	// The ARN for the IAM policy applied to the QuickSight users and groups specified
	// in this assignment.
	PolicyArn *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIAMPolicyAssignmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateIAMPolicyAssignment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateIAMPolicyAssignment{}, middleware.After)
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
	if err = addOpUpdateIAMPolicyAssignmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIAMPolicyAssignment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateIAMPolicyAssignment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateIAMPolicyAssignment",
	}
}
