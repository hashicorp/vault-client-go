// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// SysWriteToolsRandomUrlbytesRequest struct for SysWriteToolsRandomUrlbytesRequest
type SysWriteToolsRandomUrlbytesRequest struct {
	// The number of bytes to generate (POST body parameter). Defaults to 32 (256 bits).
	Bytes int32 `json:"bytes"`

	// Encoding format to use. Can be \"hex\" or \"base64\". Defaults to \"base64\".
	Format string `json:"format"`

	// Which system to source random data from, ether \"platform\", \"seal\", or \"all\".
	Source string `json:"source"`
}

// NewSysWriteToolsRandomUrlbytesRequestWithDefaults instantiates a new SysWriteToolsRandomUrlbytesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSysWriteToolsRandomUrlbytesRequestWithDefaults() *SysWriteToolsRandomUrlbytesRequest {
	var this SysWriteToolsRandomUrlbytesRequest

	this.Bytes = 32
	this.Format = "base64"
	this.Source = "platform"

	return &this
}
