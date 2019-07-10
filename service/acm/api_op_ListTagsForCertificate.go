// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/ListTagsForCertificateRequest
type ListTagsForCertificateInput struct {
	_ struct{} `type:"structure"`

	// String that contains the ARN of the ACM certificate for which you want to
	// list the tags. This must have the following form:
	//
	// arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	//
	// CertificateArn is a required field
	CertificateArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForCertificateInput"}

	if s.CertificateArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateArn"))
	}
	if s.CertificateArn != nil && len(*s.CertificateArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/ListTagsForCertificateResponse
type ListTagsForCertificateOutput struct {
	_ struct{} `type:"structure"`

	// The key-value pairs that define the applied tags.
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s ListTagsForCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTagsForCertificate = "ListTagsForCertificate"

// ListTagsForCertificateRequest returns a request value for making API operation for
// AWS Certificate Manager.
//
// Lists the tags that have been applied to the ACM certificate. Use the certificate's
// Amazon Resource Name (ARN) to specify the certificate. To add a tag to an
// ACM certificate, use the AddTagsToCertificate action. To delete a tag, use
// the RemoveTagsFromCertificate action.
//
//    // Example sending a request using ListTagsForCertificateRequest.
//    req := client.ListTagsForCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/ListTagsForCertificate
func (c *Client) ListTagsForCertificateRequest(input *ListTagsForCertificateInput) ListTagsForCertificateRequest {
	op := &aws.Operation{
		Name:       opListTagsForCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagsForCertificateInput{}
	}

	req := c.newRequest(op, input, &ListTagsForCertificateOutput{})
	return ListTagsForCertificateRequest{Request: req, Input: input, Copy: c.ListTagsForCertificateRequest}
}

// ListTagsForCertificateRequest is the request type for the
// ListTagsForCertificate API operation.
type ListTagsForCertificateRequest struct {
	*aws.Request
	Input *ListTagsForCertificateInput
	Copy  func(*ListTagsForCertificateInput) ListTagsForCertificateRequest
}

// Send marshals and sends the ListTagsForCertificate API request.
func (r ListTagsForCertificateRequest) Send(ctx context.Context) (*ListTagsForCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsForCertificateResponse{
		ListTagsForCertificateOutput: r.Request.Data.(*ListTagsForCertificateOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsForCertificateResponse is the response type for the
// ListTagsForCertificate API operation.
type ListTagsForCertificateResponse struct {
	*ListTagsForCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsForCertificate request.
func (r *ListTagsForCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}