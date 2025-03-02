/*
STACKIT Secrets Manager API

This API provides endpoints for managing the Secrets-Manager.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretsmanager

type CreateInstancePayload struct {
	// A user chosen name to distinguish multiple secrets manager instances.
	// REQUIRED
	Name *string `json:"name"`
}
