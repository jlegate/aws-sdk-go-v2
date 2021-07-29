// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycluster

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycluster/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsAwsjson10_serializeOpGetRoutingControlState struct {
}

func (*awsAwsjson10_serializeOpGetRoutingControlState) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpGetRoutingControlState) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetRoutingControlStateInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("ToggleCustomerAPI.GetRoutingControlState")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentGetRoutingControlStateInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpUpdateRoutingControlState struct {
}

func (*awsAwsjson10_serializeOpUpdateRoutingControlState) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpUpdateRoutingControlState) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateRoutingControlStateInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("ToggleCustomerAPI.UpdateRoutingControlState")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentUpdateRoutingControlStateInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpUpdateRoutingControlStates struct {
}

func (*awsAwsjson10_serializeOpUpdateRoutingControlStates) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpUpdateRoutingControlStates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateRoutingControlStatesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("ToggleCustomerAPI.UpdateRoutingControlStates")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentUpdateRoutingControlStatesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson10_serializeDocumentUpdateRoutingControlStateEntries(v []types.UpdateRoutingControlStateEntry, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson10_serializeDocumentUpdateRoutingControlStateEntry(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson10_serializeDocumentUpdateRoutingControlStateEntry(v *types.UpdateRoutingControlStateEntry, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.RoutingControlArn != nil {
		ok := object.Key("RoutingControlArn")
		ok.String(*v.RoutingControlArn)
	}

	if len(v.RoutingControlState) > 0 {
		ok := object.Key("RoutingControlState")
		ok.String(string(v.RoutingControlState))
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentGetRoutingControlStateInput(v *GetRoutingControlStateInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.RoutingControlArn != nil {
		ok := object.Key("RoutingControlArn")
		ok.String(*v.RoutingControlArn)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentUpdateRoutingControlStateInput(v *UpdateRoutingControlStateInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.RoutingControlArn != nil {
		ok := object.Key("RoutingControlArn")
		ok.String(*v.RoutingControlArn)
	}

	if len(v.RoutingControlState) > 0 {
		ok := object.Key("RoutingControlState")
		ok.String(string(v.RoutingControlState))
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentUpdateRoutingControlStatesInput(v *UpdateRoutingControlStatesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.UpdateRoutingControlStateEntries != nil {
		ok := object.Key("UpdateRoutingControlStateEntries")
		if err := awsAwsjson10_serializeDocumentUpdateRoutingControlStateEntries(v.UpdateRoutingControlStateEntries, ok); err != nil {
			return err
		}
	}

	return nil
}
