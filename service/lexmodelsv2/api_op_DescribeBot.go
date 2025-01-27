// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides metadata information about a bot.
func (c *Client) DescribeBot(ctx context.Context, params *DescribeBotInput, optFns ...func(*Options)) (*DescribeBotOutput, error) {
	if params == nil {
		params = &DescribeBotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBot", params, optFns, c.addOperationDescribeBotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBotInput struct {

	// The unique identifier of the bot to describe.
	//
	// This member is required.
	BotId *string

	noSmithyDocumentSerde
}

type DescribeBotOutput struct {

	// The unique identifier of the bot.
	BotId *string

	// The name of the bot.
	BotName *string

	// The current status of the bot. When the status is Available the bot is ready to
	// be used in conversations with users.
	BotStatus types.BotStatus

	// A timestamp of the date and time that the bot was created.
	CreationDateTime *time.Time

	// Settings for managing data privacy of the bot and its conversations with users.
	DataPrivacy *types.DataPrivacy

	// The description of the bot.
	Description *string

	// The maximum time in seconds that Amazon Lex retains the data gathered in a
	// conversation.
	IdleSessionTTLInSeconds *int32

	// A timestamp of the date and time that the bot was last updated.
	LastUpdatedDateTime *time.Time

	// The Amazon Resource Name (ARN) of an IAM role that has permission to access the
	// bot.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBot{}, middleware.After)
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
	if err = addOpDescribeBotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DescribeBot",
	}
}
