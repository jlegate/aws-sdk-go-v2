// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more of your network ACLs. For more information, see Network
// ACLs (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html) in the
// Amazon Virtual Private Cloud User Guide.
func (c *Client) DescribeNetworkAcls(ctx context.Context, params *DescribeNetworkAclsInput, optFns ...func(*Options)) (*DescribeNetworkAclsOutput, error) {
	if params == nil {
		params = &DescribeNetworkAclsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNetworkAcls", params, optFns, c.addOperationDescribeNetworkAclsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNetworkAclsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNetworkAclsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * association.association-id - The ID of an association ID
	// for the ACL.
	//
	// * association.network-acl-id - The ID of the network ACL involved
	// in the association.
	//
	// * association.subnet-id - The ID of the subnet involved in
	// the association.
	//
	// * default - Indicates whether the ACL is the default network
	// ACL for the VPC.
	//
	// * entry.cidr - The IPv4 CIDR range specified in the entry.
	//
	// *
	// entry.icmp.code - The ICMP code specified in the entry, if any.
	//
	// *
	// entry.icmp.type - The ICMP type specified in the entry, if any.
	//
	// *
	// entry.ipv6-cidr - The IPv6 CIDR range specified in the entry.
	//
	// *
	// entry.port-range.from - The start of the port range specified in the entry.
	//
	// *
	// entry.port-range.to - The end of the port range specified in the entry.
	//
	// *
	// entry.protocol - The protocol specified in the entry (tcp | udp | icmp or a
	// protocol number).
	//
	// * entry.rule-action - Allows or denies the matching traffic
	// (allow | deny).
	//
	// * entry.rule-number - The number of an entry (in other words,
	// rule) in the set of ACL entries.
	//
	// * network-acl-id - The ID of the network
	// ACL.
	//
	// * owner-id - The ID of the AWS account that owns the network ACL.
	//
	// * tag:
	// - The key/value combination of a tag assigned to the resource. Use the tag key
	// in the filter name and the tag value as the filter value. For example, to find
	// all resources that have a tag with the key Owner and the value TeamA, specify
	// tag:Owner for the filter name and TeamA for the filter value.
	//
	// * tag-key - The
	// key of a tag assigned to the resource. Use this filter to find all resources
	// assigned a tag with a specific key, regardless of the tag value.
	//
	// * vpc-id - The
	// ID of the VPC for the network ACL.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// One or more network ACL IDs. Default: Describes all your network ACLs.
	NetworkAclIds []string

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeNetworkAclsOutput struct {

	// Information about one or more network ACLs.
	NetworkAcls []types.NetworkAcl

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNetworkAclsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeNetworkAcls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeNetworkAcls{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNetworkAcls(options.Region), middleware.Before); err != nil {
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

// DescribeNetworkAclsAPIClient is a client that implements the DescribeNetworkAcls
// operation.
type DescribeNetworkAclsAPIClient interface {
	DescribeNetworkAcls(context.Context, *DescribeNetworkAclsInput, ...func(*Options)) (*DescribeNetworkAclsOutput, error)
}

var _ DescribeNetworkAclsAPIClient = (*Client)(nil)

// DescribeNetworkAclsPaginatorOptions is the paginator options for
// DescribeNetworkAcls
type DescribeNetworkAclsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeNetworkAclsPaginator is a paginator for DescribeNetworkAcls
type DescribeNetworkAclsPaginator struct {
	options   DescribeNetworkAclsPaginatorOptions
	client    DescribeNetworkAclsAPIClient
	params    *DescribeNetworkAclsInput
	nextToken *string
	firstPage bool
}

// NewDescribeNetworkAclsPaginator returns a new DescribeNetworkAclsPaginator
func NewDescribeNetworkAclsPaginator(client DescribeNetworkAclsAPIClient, params *DescribeNetworkAclsInput, optFns ...func(*DescribeNetworkAclsPaginatorOptions)) *DescribeNetworkAclsPaginator {
	if params == nil {
		params = &DescribeNetworkAclsInput{}
	}

	options := DescribeNetworkAclsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeNetworkAclsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeNetworkAclsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeNetworkAcls page.
func (p *DescribeNetworkAclsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeNetworkAclsOutput, error) {
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

	result, err := p.client.DescribeNetworkAcls(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeNetworkAcls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeNetworkAcls",
	}
}
