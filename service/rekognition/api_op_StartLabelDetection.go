// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts asynchronous detection of labels in a stored video. Amazon Rekognition
// Video can detect labels in a video. Labels are instances of real-world entities.
// This includes objects like flower, tree, and table; events like wedding,
// graduation, and birthday party; concepts like landscape, evening, and nature;
// and activities like a person getting out of a car or a person skiing. The video
// must be stored in an Amazon S3 bucket. Use Video to specify the bucket name and
// the filename of the video. StartLabelDetection returns a job identifier (JobId)
// which you use to get the results of the operation. When label detection is
// finished, Amazon Rekognition Video publishes a completion status to the Amazon
// Simple Notification Service topic that you specify in NotificationChannel. To
// get the results of the label detection operation, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetLabelDetection and pass the job identifier (JobId) from the initial call to
// StartLabelDetection.
func (c *Client) StartLabelDetection(ctx context.Context, params *StartLabelDetectionInput, optFns ...func(*Options)) (*StartLabelDetectionOutput, error) {
	if params == nil {
		params = &StartLabelDetectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartLabelDetection", params, optFns, c.addOperationStartLabelDetectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartLabelDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartLabelDetectionInput struct {

	// The video in which you want to detect labels. The video must be stored in an
	// Amazon S3 bucket.
	//
	// This member is required.
	Video *types.Video

	// Idempotent token used to identify the start request. If you use the same token
	// with multiple StartLabelDetection requests, the same JobId is returned. Use
	// ClientRequestToken to prevent the same job from being accidently started more
	// than once.
	ClientRequestToken *string

	// An identifier you specify that's returned in the completion notification that's
	// published to your Amazon Simple Notification Service topic. For example, you can
	// use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string

	// Specifies the minimum confidence that Amazon Rekognition Video must have in
	// order to return a detected label. Confidence represents how certain Amazon
	// Rekognition is that a label is correctly identified.0 is the lowest confidence.
	// 100 is the highest confidence. Amazon Rekognition Video doesn't return any
	// labels with a confidence level lower than this specified value. If you don't
	// specify MinConfidence, the operation returns labels with confidence values
	// greater than or equal to 50 percent.
	MinConfidence *float32

	// The Amazon SNS topic ARN you want Amazon Rekognition Video to publish the
	// completion status of the label detection operation to.
	NotificationChannel *types.NotificationChannel

	noSmithyDocumentSerde
}

type StartLabelDetectionOutput struct {

	// The identifier for the label detection job. Use JobId to identify the job in a
	// subsequent call to GetLabelDetection.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartLabelDetectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartLabelDetection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartLabelDetection{}, middleware.After)
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
	if err = addOpStartLabelDetectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartLabelDetection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartLabelDetection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartLabelDetection",
	}
}
