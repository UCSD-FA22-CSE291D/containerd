/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

import "github.com/Microsoft/hcsshim/internal/protocol/guestrequest"

type ModifySettingRequest struct {
	ResourcePath string `json:"ResourcePath,omitempty"`

	RequestType guestrequest.RequestType `json:"RequestType,omitempty"` // NOTE: Swagger generated as string. Locally updated.

	Settings interface{} `json:"Settings,omitempty"` // NOTE: Swagger generated as *interface{}. Locally updated

	GuestRequest interface{} `json:"GuestRequest,omitempty"` // NOTE: Swagger generated as *interface{}. Locally updated
}
