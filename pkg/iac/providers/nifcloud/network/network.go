package network

import iacTypes "vulnscan/pkg/iac/types"

type Network struct {
	ElasticLoadBalancers []ElasticLoadBalancer
	LoadBalancers        []LoadBalancer
	Routers              []Router
	VpnGateways          []VpnGateway
}

type NetworkInterface struct {
	Metadata     iacTypes.Metadata
	NetworkID    iacTypes.StringValue
	IsVipNetwork iacTypes.BoolValue
}
