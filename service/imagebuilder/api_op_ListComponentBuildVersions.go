// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of component build versions for the specified semantic version.
// The semantic version has four nodes: ../. You can assign values for the first
// three, and can filter on all of them. Filtering: When you retrieve or reference
// a resource with a semantic version, you can use wildcards (x) to filter your
// results. When you use a wildcard in any node, all nodes to the right of the
// first wildcard must also be wildcards. For example, specifying "1.2.x", or
// "1.x.x" works to filter list results, but neither "1.x.2", nor "x.2.x" will
// work. You do not have to specify the build - Image Builder automatically uses a
// wildcard for that, if applicable.
func (c *Client) ListComponentBuildVersions(ctx context.Context, params *ListComponentBuildVersionsInput, optFns ...func(*Options)) (*ListComponentBuildVersionsOutput, error) {
	if params == nil {
		params = &ListComponentBuildVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComponentBuildVersions", params, optFns, c.addOperationListComponentBuildVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComponentBuildVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComponentBuildVersionsInput struct {

	// The component version Amazon Resource Name (ARN) whose versions you want to
	// list.
	//
	// This member is required.
	ComponentVersionArn *string

	// The maximum items to return in a request.
	MaxResults int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListComponentBuildVersionsOutput struct {

	// The list of component summaries for the specified semantic version.
	ComponentSummaryList []types.ComponentSummary

	// The next token used for paginated responses. When this is not empty, there are
	// additional elements that the service has not included in this request. Use this
	// token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListComponentBuildVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListComponentBuildVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListComponentBuildVersions{}, middleware.After)
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
	if err = addOpListComponentBuildVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComponentBuildVersions(options.Region), middleware.Before); err != nil {
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

// ListComponentBuildVersionsAPIClient is a client that implements the
// ListComponentBuildVersions operation.
type ListComponentBuildVersionsAPIClient interface {
	ListComponentBuildVersions(context.Context, *ListComponentBuildVersionsInput, ...func(*Options)) (*ListComponentBuildVersionsOutput, error)
}

var _ ListComponentBuildVersionsAPIClient = (*Client)(nil)

// ListComponentBuildVersionsPaginatorOptions is the paginator options for
// ListComponentBuildVersions
type ListComponentBuildVersionsPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListComponentBuildVersionsPaginator is a paginator for
// ListComponentBuildVersions
type ListComponentBuildVersionsPaginator struct {
	options   ListComponentBuildVersionsPaginatorOptions
	client    ListComponentBuildVersionsAPIClient
	params    *ListComponentBuildVersionsInput
	nextToken *string
	firstPage bool
}

// NewListComponentBuildVersionsPaginator returns a new
// ListComponentBuildVersionsPaginator
func NewListComponentBuildVersionsPaginator(client ListComponentBuildVersionsAPIClient, params *ListComponentBuildVersionsInput, optFns ...func(*ListComponentBuildVersionsPaginatorOptions)) *ListComponentBuildVersionsPaginator {
	if params == nil {
		params = &ListComponentBuildVersionsInput{}
	}

	options := ListComponentBuildVersionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListComponentBuildVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListComponentBuildVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListComponentBuildVersions page.
func (p *ListComponentBuildVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListComponentBuildVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListComponentBuildVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListComponentBuildVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "ListComponentBuildVersions",
	}
}
