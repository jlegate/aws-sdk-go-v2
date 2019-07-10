// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/PutDestinationRequest
type PutDestinationInput struct {
	_ struct{} `type:"structure"`

	// A name for the destination.
	//
	// DestinationName is a required field
	DestinationName *string `locationName:"destinationName" min:"1" type:"string" required:"true"`

	// The ARN of an IAM role that grants CloudWatch Logs permissions to call the
	// Amazon Kinesis PutRecord operation on the destination stream.
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" min:"1" type:"string" required:"true"`

	// The ARN of an Amazon Kinesis stream to which to deliver matching log events.
	//
	// TargetArn is a required field
	TargetArn *string `locationName:"targetArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutDestinationInput"}

	if s.DestinationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationName"))
	}
	if s.DestinationName != nil && len(*s.DestinationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationName", 1))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 1))
	}

	if s.TargetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetArn"))
	}
	if s.TargetArn != nil && len(*s.TargetArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/PutDestinationResponse
type PutDestinationOutput struct {
	_ struct{} `type:"structure"`

	// The destination.
	Destination *Destination `locationName:"destination" type:"structure"`
}

// String returns the string representation
func (s PutDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutDestination = "PutDestination"

// PutDestinationRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Creates or updates a destination. A destination encapsulates a physical resource
// (such as an Amazon Kinesis stream) and enables you to subscribe to a real-time
// stream of log events for a different account, ingested using PutLogEvents.
// Currently, the only supported physical resource is a Kinesis stream belonging
// to the same account as the destination.
//
// Through an access policy, a destination controls what is written to its Kinesis
// stream. By default, PutDestination does not set any access policy with the
// destination, which means a cross-account user cannot call PutSubscriptionFilter
// against this destination. To enable this, the destination owner must call
// PutDestinationPolicy after PutDestination.
//
//    // Example sending a request using PutDestinationRequest.
//    req := client.PutDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/PutDestination
func (c *Client) PutDestinationRequest(input *PutDestinationInput) PutDestinationRequest {
	op := &aws.Operation{
		Name:       opPutDestination,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutDestinationInput{}
	}

	req := c.newRequest(op, input, &PutDestinationOutput{})
	return PutDestinationRequest{Request: req, Input: input, Copy: c.PutDestinationRequest}
}

// PutDestinationRequest is the request type for the
// PutDestination API operation.
type PutDestinationRequest struct {
	*aws.Request
	Input *PutDestinationInput
	Copy  func(*PutDestinationInput) PutDestinationRequest
}

// Send marshals and sends the PutDestination API request.
func (r PutDestinationRequest) Send(ctx context.Context) (*PutDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutDestinationResponse{
		PutDestinationOutput: r.Request.Data.(*PutDestinationOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutDestinationResponse is the response type for the
// PutDestination API operation.
type PutDestinationResponse struct {
	*PutDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutDestination request.
func (r *PutDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}