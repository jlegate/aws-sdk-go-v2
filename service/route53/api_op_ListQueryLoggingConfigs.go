// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the configurations for DNS query logging that are associated with the
// current account or the configuration that is associated with a specified hosted
// zone. For more information about DNS query logs, see CreateQueryLoggingConfig
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateQueryLoggingConfig.html).
// Additional information, including the format of DNS query logs, appears in
// Logging DNS Queries
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html) in
// the Amazon Route 53 Developer Guide.
func (c *Client) ListQueryLoggingConfigs(ctx context.Context, params *ListQueryLoggingConfigsInput, optFns ...func(*Options)) (*ListQueryLoggingConfigsOutput, error) {
	if params == nil {
		params = &ListQueryLoggingConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQueryLoggingConfigs", params, optFns, c.addOperationListQueryLoggingConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQueryLoggingConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQueryLoggingConfigsInput struct {

	// (Optional) If you want to list the query logging configuration that is
	// associated with a hosted zone, specify the ID in HostedZoneId. If you don't
	// specify a hosted zone ID, ListQueryLoggingConfigs returns all of the
	// configurations that are associated with the current account.
	HostedZoneId *string

	// (Optional) The maximum number of query logging configurations that you want
	// Amazon Route 53 to return in response to the current request. If the current
	// account has more than MaxResults configurations, use the value of NextToken
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListQueryLoggingConfigs.html#API_ListQueryLoggingConfigs_RequestSyntax)
	// in the response to get the next page of results. If you don't specify a value
	// for MaxResults, Route 53 returns up to 100 configurations.
	MaxResults *int32

	// (Optional) If the current account has more than MaxResults query logging
	// configurations, use NextToken to get the second and subsequent pages of results.
	// For the first ListQueryLoggingConfigs request, omit this value. For the second
	// and subsequent requests, get the value of NextToken from the previous response
	// and specify that value for NextToken in the request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListQueryLoggingConfigsOutput struct {

	// An array that contains one QueryLoggingConfig
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_QueryLoggingConfig.html)
	// element for each configuration for DNS query logging that is associated with the
	// current account.
	//
	// This member is required.
	QueryLoggingConfigs []types.QueryLoggingConfig

	// If a response includes the last of the query logging configurations that are
	// associated with the current account, NextToken doesn't appear in the response.
	// If a response doesn't include the last of the configurations, you can get more
	// configurations by submitting another ListQueryLoggingConfigs
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListQueryLoggingConfigs.html)
	// request. Get the value of NextToken that Amazon Route 53 returned in the
	// previous response and include it in NextToken in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListQueryLoggingConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListQueryLoggingConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListQueryLoggingConfigs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQueryLoggingConfigs(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListQueryLoggingConfigsAPIClient is a client that implements the
// ListQueryLoggingConfigs operation.
type ListQueryLoggingConfigsAPIClient interface {
	ListQueryLoggingConfigs(context.Context, *ListQueryLoggingConfigsInput, ...func(*Options)) (*ListQueryLoggingConfigsOutput, error)
}

var _ ListQueryLoggingConfigsAPIClient = (*Client)(nil)

// ListQueryLoggingConfigsPaginatorOptions is the paginator options for
// ListQueryLoggingConfigs
type ListQueryLoggingConfigsPaginatorOptions struct {
	// (Optional) The maximum number of query logging configurations that you want
	// Amazon Route 53 to return in response to the current request. If the current
	// account has more than MaxResults configurations, use the value of NextToken
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListQueryLoggingConfigs.html#API_ListQueryLoggingConfigs_RequestSyntax)
	// in the response to get the next page of results. If you don't specify a value
	// for MaxResults, Route 53 returns up to 100 configurations.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListQueryLoggingConfigsPaginator is a paginator for ListQueryLoggingConfigs
type ListQueryLoggingConfigsPaginator struct {
	options   ListQueryLoggingConfigsPaginatorOptions
	client    ListQueryLoggingConfigsAPIClient
	params    *ListQueryLoggingConfigsInput
	nextToken *string
	firstPage bool
}

// NewListQueryLoggingConfigsPaginator returns a new
// ListQueryLoggingConfigsPaginator
func NewListQueryLoggingConfigsPaginator(client ListQueryLoggingConfigsAPIClient, params *ListQueryLoggingConfigsInput, optFns ...func(*ListQueryLoggingConfigsPaginatorOptions)) *ListQueryLoggingConfigsPaginator {
	if params == nil {
		params = &ListQueryLoggingConfigsInput{}
	}

	options := ListQueryLoggingConfigsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListQueryLoggingConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListQueryLoggingConfigsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListQueryLoggingConfigs page.
func (p *ListQueryLoggingConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListQueryLoggingConfigsOutput, error) {
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

	result, err := p.client.ListQueryLoggingConfigs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListQueryLoggingConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListQueryLoggingConfigs",
	}
}
