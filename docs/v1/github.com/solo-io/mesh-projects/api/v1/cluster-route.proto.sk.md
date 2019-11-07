
---
title: "cluster-route.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `zephyr.solo.io` 
#### Types:


- [MeshBridge](#meshbridge) **Top-Level Resource**
  



##### Source File: [github.com/solo-io/mesh-projects/api/v1/cluster-route.proto](https://github.com/solo-io/mesh-projects/blob/master/api/v1/cluster-route.proto)





---
### MeshBridge



```yaml
"status": .core.solo.io.Status
"metadata": .core.solo.io.Metadata
"sourceMesh": .core.zephyr.solo.io.ClusterResourceRef
"targetMesh": .core.zephyr.solo.io.ClusterResourceRef
"glooNamespace": string
"source": string
"target": .core.zephyr.solo.io.ClusterResourceRef

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `status` | [.core.solo.io.Status](../../../../solo-kit/api/v1/status.proto.sk/#status) | Status indicates the validation status of this resource. Status is read-only by clients, and set by gloo during validation. |  |
| `metadata` | [.core.solo.io.Metadata](../../../../solo-kit/api/v1/metadata.proto.sk/#metadata) | Metadata contains the object metadata for this resource. |  |
| `sourceMesh` | [.core.zephyr.solo.io.ClusterResourceRef](../core/ref.proto.sk/#clusterresourceref) | Source mesh from which to create the bridge. |  |
| `targetMesh` | [.core.zephyr.solo.io.ClusterResourceRef](../core/ref.proto.sk/#clusterresourceref) | Target mesh to which the bridge will be pointing. |  |
| `glooNamespace` | `string` | Namespace in which to find gloo in the target mesh cluster. |  |
| `source` | `string` | Source service which will be using the bridge. Currently a string as it assumes that the service which will need this bridge is located in the same namespace as the bridge object. |  |
| `target` | [.core.zephyr.solo.io.ClusterResourceRef](../core/ref.proto.sk/#clusterresourceref) | Target service in a different mesh to which this bridge will allow traffic from the source service. This source pointer is important as the name of this service (e.g. `gloo.gloo-system.svc.cluster.local`) is translated into `gloo.gloo-system.global`. Which allows is to be routed to and understood by istio as being in a foreign cluster, or other non-local space. |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->