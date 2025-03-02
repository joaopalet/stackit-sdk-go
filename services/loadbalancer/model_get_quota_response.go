/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  This beta load balancer service is provided free of charge. For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type GetQuotaResponse struct {
	// The maximum number of load balancing servers in this project. Unlimited if set to -1.
	MaxLoadBalancers *int32 `json:"maxLoadBalancers,omitempty"`
	// Project identifier
	ProjectId *string `json:"projectId,omitempty"`
}
