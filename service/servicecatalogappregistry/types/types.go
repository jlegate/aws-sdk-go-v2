// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Represents a Service Catalog AppRegistry application that is the top-level node
// in a hierarchy of related cloud resource abstractions.
type Application struct {

	// The Amazon resource name (ARN) that specifies the application across services.
	Arn *string

	// The ISO-8601 formatted timestamp of the moment when the application was created.
	CreationTime *time.Time

	// The description of the application.
	Description *string

	// The identifier of the application.
	Id *string

	// The ISO-8601 formatted timestamp of the moment when the application was last
	// updated.
	LastUpdateTime *time.Time

	// The name of the application. The name must be unique in the region in which you
	// are creating the application.
	Name *string

	// Key-value pairs you can use to associate with the application.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Summary of a Service Catalog AppRegistry application.
type ApplicationSummary struct {

	// The Amazon resource name (ARN) that specifies the application across services.
	Arn *string

	// The ISO-8601 formatted timestamp of the moment when the application was created.
	CreationTime *time.Time

	// The description of the application.
	Description *string

	// The identifier of the application.
	Id *string

	// The ISO-8601 formatted timestamp of the moment when the application was last
	// updated.
	LastUpdateTime *time.Time

	// The name of the application. The name must be unique in the region in which you
	// are creating the application.
	Name *string

	noSmithyDocumentSerde
}

// Represents a Service Catalog AppRegistry attribute group that is rich metadata
// which describes an application and its components.
type AttributeGroup struct {

	// The Amazon resource name (ARN) that specifies the attribute group across
	// services.
	Arn *string

	// The ISO-8601 formatted timestamp of the moment the attribute group was created.
	CreationTime *time.Time

	// The description of the attribute group that the user provides.
	Description *string

	// The globally unique attribute group identifier of the attribute group.
	Id *string

	// The ISO-8601 formatted timestamp of the moment the attribute group was last
	// updated. This time is the same as the creationTime for a newly created attribute
	// group.
	LastUpdateTime *time.Time

	// The name of the attribute group.
	Name *string

	// Key-value pairs you can use to associate with the attribute group.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Summary of a Service Catalog AppRegistry attribute group.
type AttributeGroupSummary struct {

	// The Amazon resource name (ARN) that specifies the attribute group across
	// services.
	Arn *string

	// The ISO-8601 formatted timestamp of the moment the attribute group was created.
	CreationTime *time.Time

	// The description of the attribute group that the user provides.
	Description *string

	// The globally unique attribute group identifier of the attribute group.
	Id *string

	// The ISO-8601 formatted timestamp of the moment the attribute group was last
	// updated. This time is the same as the creationTime for a newly created attribute
	// group.
	LastUpdateTime *time.Time

	// The name of the attribute group.
	Name *string

	noSmithyDocumentSerde
}

// Information about the resource.
type ResourceInfo struct {

	// The Amazon resource name (ARN) that specifies the resource across services.
	Arn *string

	// The name of the resource.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
