/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// A set of properties that uniquely identify a given Docker image.
type ImageFingerprint struct {

	// Required. The layer ID of the final layer in the Docker image's v1 representation.
	V1Name string `json:"v1_name,omitempty"`

	// Required. The ordered list of v2 blobs that represent a given image.
	V2Blob []string `json:"v2_blob,omitempty"`

	// Output only. The name of the image's v2 blobs computed via:   [bottom] := v2_blob[bottom]   [N] := sha256(v2_blob[N] + \" \" + v2_name[N+1]) Only the name of the final blob is kept.
	V2Name string `json:"v2_name,omitempty"`
}
