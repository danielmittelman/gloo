
---
title: "filter.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `envoy.config.filter.http.aws_lambda.v2` 
#### Types:


- [AWSLambdaPerRoute](#awslambdaperroute)
- [AWSLambdaProtocolExtension](#awslambdaprotocolextension)
- [AWSLambdaConfig](#awslambdaconfig)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/aws/filter.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/external/envoy/extensions/aws/filter.proto)





---
### AWSLambdaPerRoute

 
AWS Lambda contains the configuration necessary to perform transform regular
http calls to AWS Lambda invocations.

```yaml
"name": string
"qualifier": string
"async": bool
"emptyBodyOverride": .google.protobuf.StringValue

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` | The name of the function. |  |
| `qualifier` | `string` | The qualifier of the function (defaults to $LATEST if not specified). |  |
| `async` | `bool` | Invocation type - async or regular. |  |
| `emptyBodyOverride` | [.google.protobuf.StringValue](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/string-value) | Optional default body if the body is empty. By default on default body is used if the body empty, and an empty body will be sent upstream. |  |




---
### AWSLambdaProtocolExtension



```yaml
"host": string
"region": string
"accessKey": string
"secretKey": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `host` | `string` | The host header for AWS this cluster. |  |
| `region` | `string` | The region for this cluster. |  |
| `accessKey` | `string` | The access_key for AWS this cluster. |  |
| `secretKey` | `string` | The secret_key for AWS this cluster. |  |




---
### AWSLambdaConfig



```yaml
"useDefaultCredentials": .google.protobuf.BoolValue

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `useDefaultCredentials` | [.google.protobuf.BoolValue](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/bool-value) | Use AWS default credentials chain to get credentials. This will search environment variables, ECS metadata and instance metadata to get the credentials. credentials will be rotated automatically. If credentials are provided on the cluster (using the AWSLambdaProtocolExtension), it will override these credentials. This defaults to false, but may change in the future to true. |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->