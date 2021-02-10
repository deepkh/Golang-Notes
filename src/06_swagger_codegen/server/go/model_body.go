/*
 * Sample API
 *
 * Production: http://api.example.com/v1  Staging: http://staging-api.example.com/v1 
 *
 * API version: 0.1.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"os"
)

type Body struct {

	OrderId int32 `json:"orderId,omitempty"`

	UserId int32 `json:"userId,omitempty"`

	FileName **os.File `json:"fileName,omitempty"`
}
