/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type TimeZoneInformation struct {
	Bias int32 `json:"Bias,omitempty"`

	StandardName string `json:"StandardName,omitempty"`

	StandardDate *SystemTime `json:"StandardDate,omitempty"`

	StandardBias int32 `json:"StandardBias,omitempty"`

	DaylightName string `json:"DaylightName,omitempty"`

	DaylightDate *SystemTime `json:"DaylightDate,omitempty"`

	DaylightBias int32 `json:"DaylightBias,omitempty"`
}
