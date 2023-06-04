/*
 * Otito API documentation
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.0
 * Contact: contact@ayinke.ventures
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ServerSwaggerMessageRespMessages struct {
	ApplicationId string `json:"application_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Id string `json:"id,omitempty"`
	// This is for the request
	IpAddress string `json:"ip_address,omitempty"`
	Method string `json:"method,omitempty"`
	OrganisationId string `json:"organisation_id,omitempty"`
	Path string `json:"path,omitempty"`
	Reference string `json:"reference,omitempty"`
	Request *ServerMessageHttpDefinition `json:"request,omitempty"`
	Response *ServerMessageHttpDefinition `json:"response,omitempty"`
	StatusCode int32 `json:"status_code,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}