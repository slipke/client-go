/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// DiscoveredContinuousAnalysis : Whether the resource is continuously analyzed.   - CONTINUOUS_ANALYSIS_UNSPECIFIED: Unknown.  - ACTIVE: The resource is continuously analyzed.  - INACTIVE: The resource is ignored for continuous analysis.
type DiscoveredContinuousAnalysis string

// List of DiscoveredContinuousAnalysis
const (
	CONTINUOUS_ANALYSIS_UNSPECIFIED DiscoveredContinuousAnalysis = "CONTINUOUS_ANALYSIS_UNSPECIFIED"
	ACTIVE DiscoveredContinuousAnalysis = "ACTIVE"
	INACTIVE DiscoveredContinuousAnalysis = "INACTIVE"
)