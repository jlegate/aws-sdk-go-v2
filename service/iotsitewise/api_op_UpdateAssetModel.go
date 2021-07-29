// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an asset model and all of the assets that were created from the model.
// Each asset created from the model inherits the updated asset model's property
// and hierarchy definitions. For more information, see Updating assets and models
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/update-assets-and-models.html)
// in the IoT SiteWise User Guide. This operation overwrites the existing model
// with the provided model. To avoid deleting your asset model's properties or
// hierarchies, you must include their IDs and definitions in the updated asset
// model payload. For more information, see DescribeAssetModel
// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeAssetModel.html).
// If you remove a property from an asset model, IoT SiteWise deletes all previous
// data for that property. If you remove a hierarchy definition from an asset
// model, IoT SiteWise disassociates every asset associated with that hierarchy.
// You can't change the type or data type of an existing property.
func (c *Client) UpdateAssetModel(ctx context.Context, params *UpdateAssetModelInput, optFns ...func(*Options)) (*UpdateAssetModelOutput, error) {
	if params == nil {
		params = &UpdateAssetModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAssetModel", params, optFns, c.addOperationUpdateAssetModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAssetModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAssetModelInput struct {

	// The ID of the asset model to update.
	//
	// This member is required.
	AssetModelId *string

	// A unique, friendly name for the asset model.
	//
	// This member is required.
	AssetModelName *string

	// The composite asset models that are part of this asset model. Composite asset
	// models are asset models that contain specific properties. Each composite model
	// has a type that defines the properties that the composite model supports. Use
	// composite asset models to define alarms on this asset model.
	AssetModelCompositeModels []types.AssetModelCompositeModel

	// A description for the asset model.
	AssetModelDescription *string

	// The updated hierarchy definitions of the asset model. Each hierarchy specifies
	// an asset model whose assets can be children of any other assets created from
	// this asset model. For more information, see Asset hierarchies
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html)
	// in the IoT SiteWise User Guide. You can specify up to 10 hierarchies per asset
	// model. For more information, see Quotas
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the
	// IoT SiteWise User Guide.
	AssetModelHierarchies []types.AssetModelHierarchy

	// The updated property definitions of the asset model. For more information, see
	// Asset properties
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html)
	// in the IoT SiteWise User Guide. You can specify up to 200 properties per asset
	// model. For more information, see Quotas
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the
	// IoT SiteWise User Guide.
	AssetModelProperties []types.AssetModelProperty

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string

	noSmithyDocumentSerde
}

type UpdateAssetModelOutput struct {

	// The status of the asset model, which contains a state (UPDATING after
	// successfully calling this operation) and any error message.
	//
	// This member is required.
	AssetModelStatus *types.AssetModelStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAssetModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAssetModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAssetModel{}, middleware.After)
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
	if err = addEndpointPrefix_opUpdateAssetModelMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opUpdateAssetModelMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateAssetModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAssetModel(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateAssetModelMiddleware struct {
}

func (*endpointPrefix_opUpdateAssetModelMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateAssetModelMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opUpdateAssetModelMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opUpdateAssetModelMiddleware{}, `OperationSerializer`, middleware.After)
}

type idempotencyToken_initializeOpUpdateAssetModel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateAssetModel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateAssetModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateAssetModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateAssetModelInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateAssetModelMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateAssetModel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateAssetModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "UpdateAssetModel",
	}
}
