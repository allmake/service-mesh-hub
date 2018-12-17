<!-- Code generated by protoc-gen-solo-kit. DO NOT EDIT. -->

## Package:
supergloo.solo.io

## Source File:
routing.proto 

## Description:  

## Contents:
- Messages:  
	- [RoutingRule](#RoutingRule)  
	- [TrafficShifting](#TrafficShifting)  
	- [WeightedDestination](#WeightedDestination)  
	- [HeaderManipulation](#HeaderManipulation)  
	- [Percent](#Percent)

---
  
### <a name="RoutingRule">RoutingRule</a>

Description: rules to add features such as Fault Injection and Retries to a mesh

```yaml
"status": .core.solo.io.Status
"metadata": .core.solo.io.Metadata
"target_mesh": .core.solo.io.ResourceRef
"sources": [.core.solo.io.ResourceRef]
"destinations": [.core.solo.io.ResourceRef]
"request_matchers": [.gloo.solo.io.Matcher]
"traffic_shifting": .supergloo.solo.io.TrafficShifting
"fault_injection": .networking.istio.io.HTTPFaultInjection
"timeout": .google.protobuf.Duration
"retries": .networking.istio.io.HTTPRetry
"cors_policy": .networking.istio.io.CorsPolicy
"mirror": .gloo.solo.io.Destination
"header_manipulaition": .supergloo.solo.io.HeaderManipulation

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| status | [.core.solo.io.Status](routing.proto.sk.md#RoutingRule) | Status indicates the validation status of this resource. Status is read-only by clients, and set by gloo during validation |  |
| metadata | [.core.solo.io.Metadata](routing.proto.sk.md#RoutingRule) | Metadata contains the object metadata for this resource |  |
| target_mesh | [.core.solo.io.ResourceRef](routing.proto.sk.md#RoutingRule) | target where we apply this rule |  |
| sources | [[.core.solo.io.ResourceRef]](routing.proto.sk.md#RoutingRule) | source upstreams to apply the rule to. if empty, applies to all sources. |  |
| destinations | [[.core.solo.io.ResourceRef]](routing.proto.sk.md#RoutingRule) | destination upstreams for which this rule applies. if empty, applies to all destinations |  |
| request_matchers | [[.gloo.solo.io.Matcher]](routing.proto.sk.md#RoutingRule) | if specified, this rule will only apply to http requests in the mesh matching these parameters |  |
| traffic_shifting | [.supergloo.solo.io.TrafficShifting](routing.proto.sk.md#RoutingRule) | configuration to enable traffic shifting, e.g. by percentage or for alternate destinations |  |
| fault_injection | [.networking.istio.io.HTTPFaultInjection](routing.proto.sk.md#RoutingRule) | configuration to enable fault injection for this rule |  |
| timeout | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) | Timeout for this rule |  |
| retries | [.networking.istio.io.HTTPRetry](routing.proto.sk.md#RoutingRule) | Retry policy for for this rule |  |
| cors_policy | [.networking.istio.io.CorsPolicy](routing.proto.sk.md#RoutingRule) | Cross-Origin Resource Sharing policy (CORS) for this rule. Refer to https://developer.mozilla.org/en-US/docs/Web/HTTP/Access_control_CORS for further details about cross origin resource sharing. |  |
| mirror | [.gloo.solo.io.Destination](routing.proto.sk.md#RoutingRule) | Mirror HTTP traffic to a another destination for this rule. Traffic will still be sent to its original destination as normal. |  |
| header_manipulaition | [.supergloo.solo.io.HeaderManipulation](routing.proto.sk.md#RoutingRule) | manipulate request and response headers for this rule |  |
  
### <a name="TrafficShifting">TrafficShifting</a>

Description: enable traffic shifting for any http requests sent to one of the destinations on this rule

```yaml
"destinations": [.supergloo.solo.io.WeightedDestination]

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| destinations | [[.supergloo.solo.io.WeightedDestination]](routing.proto.sk.md#TrafficShifting) | split traffic between these subsets based on their weights weights should add to 100 |  |
  
### <a name="WeightedDestination">WeightedDestination</a>

Description: WeightedDestination attaches a weight to a single destination.

```yaml
"upstream": .core.solo.io.ResourceRef
"weight": int

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| upstream | [.core.solo.io.ResourceRef](routing.proto.sk.md#WeightedDestination) |  |  |
| weight | int | Weight must be greater than zero Routing to each destination will be balanced by the ratio of the destination's weight to the total weight on a route |  |
  
### <a name="HeaderManipulation">HeaderManipulation</a>

Description: manipulate request and response headers

```yaml
"remove_response_headers": [string]
"append_response_headers": [.supergloo.solo.io.HeaderManipulation.AppendResponseHeadersEntry]
"remove_request_headers": [string]
"append_request_headers": [.supergloo.solo.io.HeaderManipulation.AppendRequestHeadersEntry]

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| remove_response_headers | [string] | HTTP headers to remove before returning a response to the caller. |  |
| append_response_headers | [[.supergloo.solo.io.HeaderManipulation.AppendResponseHeadersEntry]](routing.proto.sk.md#HeaderManipulation) | Additional HTTP headers to add before returning a response to the caller. |  |
| remove_request_headers | [string] | HTTP headers to remove before forwarding a request to the destination service. |  |
| append_request_headers | [[.supergloo.solo.io.HeaderManipulation.AppendRequestHeadersEntry]](routing.proto.sk.md#HeaderManipulation) | Additional HTTP headers to add before forwarding a request to the destination service. |  |
  
### <a name="Percent">Percent</a>

Description: Percent specifies a percentage in the range of [0.0, 100.0].

```yaml
"value": float

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| value | float |  |  |

