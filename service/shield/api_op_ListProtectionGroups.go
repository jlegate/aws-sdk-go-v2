// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the ProtectionGroup objects for the account.
func (c *Client) ListProtectionGroups(ctx context.Context, params *ListProtectionGroupsInput, optFns ...func(*Options)) (*ListProtectionGroupsOutput, error) {
	if params == nil {
		params = &ListProtectionGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProtectionGroups", params, optFns, c.addOperationListProtectionGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProtectionGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProtectionGroupsInput struct {

	// The maximum number of ProtectionGroup objects to return. If you leave this
	// blank, Shield Advanced returns the first 20 results. This is a maximum value.
	// Shield Advanced might return the results in smaller batches. That is, the number
	// of objects returned could be less than MaxResults, even if there are still more
	// objects yet to return. If there are more objects to return, Shield Advanced
	// returns a value in NextToken that you can use in your next request, to get the
	// next batch of objects.
	MaxResults *int32

	// The next token value from a previous call to ListProtectionGroups. Pass null if
	// this is the first call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListProtectionGroupsOutput struct {

	//
	//
	// This member is required.
	ProtectionGroups []types.ProtectionGroup

	// If you specify a value for MaxResults and you have more protection groups than
	// the value of MaxResults, AWS Shield Advanced returns this token that you can use
	// in your next request, to get the next batch of objects.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProtectionGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProtectionGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProtectionGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProtectionGroups(options.Region), middleware.Before); err != nil {
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

// ListProtectionGroupsAPIClient is a client that implements the
// ListProtectionGroups operation.
type ListProtectionGroupsAPIClient interface {
	ListProtectionGroups(context.Context, *ListProtectionGroupsInput, ...func(*Options)) (*ListProtectionGroupsOutput, error)
}

var _ ListProtectionGroupsAPIClient = (*Client)(nil)

// ListProtectionGroupsPaginatorOptions is the paginator options for
// ListProtectionGroups
type ListProtectionGroupsPaginatorOptions struct {
	// The maximum number of ProtectionGroup objects to return. If you leave this
	// blank, Shield Advanced returns the first 20 results. This is a maximum value.
	// Shield Advanced might return the results in smaller batches. That is, the number
	// of objects returned could be less than MaxResults, even if there are still more
	// objects yet to return. If there are more objects to return, Shield Advanced
	// returns a value in NextToken that you can use in your next request, to get the
	// next batch of objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProtectionGroupsPaginator is a paginator for ListProtectionGroups
type ListProtectionGroupsPaginator struct {
	options   ListProtectionGroupsPaginatorOptions
	client    ListProtectionGroupsAPIClient
	params    *ListProtectionGroupsInput
	nextToken *string
	firstPage bool
}

// NewListProtectionGroupsPaginator returns a new ListProtectionGroupsPaginator
func NewListProtectionGroupsPaginator(client ListProtectionGroupsAPIClient, params *ListProtectionGroupsInput, optFns ...func(*ListProtectionGroupsPaginatorOptions)) *ListProtectionGroupsPaginator {
	if params == nil {
		params = &ListProtectionGroupsInput{}
	}

	options := ListProtectionGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProtectionGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProtectionGroupsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListProtectionGroups page.
func (p *ListProtectionGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProtectionGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListProtectionGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListProtectionGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "ListProtectionGroups",
	}
}
