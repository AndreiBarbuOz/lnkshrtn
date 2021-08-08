// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package ports

// Defines values for LinkObjectApiVersion.
const (
	LinkObjectApiVersionV1 LinkObjectApiVersion = "v1"
)

// Defines values for LinkObjectListApiVersion.
const (
	LinkObjectListApiVersionV1 LinkObjectListApiVersion = "v1"
)

// LinkObject defines model for LinkObject.
type LinkObject struct {
	ApiVersion *LinkObjectApiVersion   `json:"apiVersion,omitempty"`
	Metadata   *map[string]interface{} `json:"metadata,omitempty"`
	Spec       *LinkObjectSpec         `json:"spec,omitempty"`
}

// LinkObjectApiVersion defines model for LinkObject.ApiVersion.
type LinkObjectApiVersion string

// LinkObjectList defines model for LinkObjectList.
type LinkObjectList struct {
	ApiVersion *LinkObjectListApiVersion `json:"apiVersion,omitempty"`
	Items      *[]LinkObject             `json:"items,omitempty"`
	Kind       *string                   `json:"kind,omitempty"`
}

// LinkObjectListApiVersion defines model for LinkObjectList.ApiVersion.
type LinkObjectListApiVersion string

// LinkObjectSpec defines model for LinkObjectSpec.
type LinkObjectSpec struct {
	Shortned *string `json:"shortned,omitempty"`
	Url      string  `json:"url"`
}

// PostLinksJSONBody defines parameters for PostLinks.
type PostLinksJSONBody LinkObjectSpec

// PostLinksJSONRequestBody defines body for PostLinks for application/json ContentType.
type PostLinksJSONRequestBody PostLinksJSONBody
