/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  This beta load balancer service is provided free of charge. For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type LoadBalancer struct {
	Errors *[]LoadBalancerError `json:"errors,omitempty"`
	// External load balancer IP address where this load balancer is exposed.
	ExternalAddress *string `json:"externalAddress,omitempty"`
	// List of all listeners which will accept traffic.
	Listeners *[]Listener `json:"listeners,omitempty"`
	// Load balancer name
	Name *string `json:"name,omitempty"`
	// List of networks that listeners and targets reside in. Currently limited to one.
	Networks *[]Network           `json:"networks,omitempty"`
	Options  *LoadBalancerOptions `json:"options,omitempty"`
	// Transient private load balancer IP address that can change any time.
	PrivateAddress *string `json:"privateAddress,omitempty"`
	Status         *string `json:"status,omitempty"`
	// List of all target pools which will be used in the load balancer.
	TargetPools *[]TargetPool `json:"targetPools,omitempty"`
}
