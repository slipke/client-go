/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// AliasContextKind : The type of an alias.   - KIND_UNSPECIFIED: Unknown.  - FIXED: Git tag.  - MOVABLE: Git branch.  - OTHER: Used to specify non-standard aliases. For example, if a Git repo has a ref named \"refs/foo/bar\".
type AliasContextKind string

// List of AliasContextKind
const (
	KIND_UNSPECIFIED AliasContextKind = "KIND_UNSPECIFIED"
	FIXED AliasContextKind = "FIXED"
	MOVABLE AliasContextKind = "MOVABLE"
	OTHER AliasContextKind = "OTHER"
)
