/*
 * Sample API
 *
 * Production: http://api.example.com/v1  Staging: http://staging-api.example.com/v1 
 *
 * API version: 0.1.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UserName struct {

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`
}
