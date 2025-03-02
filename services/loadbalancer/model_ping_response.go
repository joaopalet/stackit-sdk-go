/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  This beta load balancer service is provided free of charge. For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type PingResponse struct {
	// Same project identifier as passed in on request.
	ProjectId *string `json:"projectId,omitempty"`
}
