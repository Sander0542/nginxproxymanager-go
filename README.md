# Go API client for nginxproxymanager

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.12.2
- Package version: 1.0.0
- Generator version: 7.11.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import nginxproxymanager "github.com/sander0542/nginxproxymanager-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `nginxproxymanager.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), nginxproxymanager.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `nginxproxymanager.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), nginxproxymanager.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `nginxproxymanager.ContextOperationServerIndices` and `nginxproxymanager.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), nginxproxymanager.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), nginxproxymanager.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessListsAPI* | [**CreateAccessList**](docs/AccessListsAPI.md#createaccesslist) | **Post** /nginx/access-lists | Create a Access List
*AccessListsAPI* | [**DeleteAccessList**](docs/AccessListsAPI.md#deleteaccesslist) | **Delete** /nginx/access-lists/{listID} | Delete a Access List
*AccessListsAPI* | [**GetAccessList**](docs/AccessListsAPI.md#getaccesslist) | **Get** /nginx/access-lists/{listID} | Get a access List
*AccessListsAPI* | [**GetAccessLists**](docs/AccessListsAPI.md#getaccesslists) | **Get** /nginx/access-lists | Get all access lists
*AccessListsAPI* | [**UpdateAccessList**](docs/AccessListsAPI.md#updateaccesslist) | **Put** /nginx/access-lists/{listID} | Update a Access List
*AuditLogAPI* | [**GetAuditLog**](docs/AuditLogAPI.md#getauditlog) | **Get** /audit-log | Get Audit Log
*CertificatesAPI* | [**CreateCertificate**](docs/CertificatesAPI.md#createcertificate) | **Post** /nginx/certificates | Create a Certificate
*CertificatesAPI* | [**DeleteCertificate**](docs/CertificatesAPI.md#deletecertificate) | **Delete** /nginx/certificates/{certID} | Delete a Certificate
*CertificatesAPI* | [**DownloadCertificate**](docs/CertificatesAPI.md#downloadcertificate) | **Get** /nginx/certificates/{certID}/download | Downloads a Certificate
*CertificatesAPI* | [**GetCertificate**](docs/CertificatesAPI.md#getcertificate) | **Get** /nginx/certificates/{certID} | Get a Certificate
*CertificatesAPI* | [**GetCertificates**](docs/CertificatesAPI.md#getcertificates) | **Get** /nginx/certificates | Get all certificates
*CertificatesAPI* | [**RenewCertificate**](docs/CertificatesAPI.md#renewcertificate) | **Post** /nginx/certificates/{certID}/renew | Renews a Certificate
*CertificatesAPI* | [**TestHttpReach**](docs/CertificatesAPI.md#testhttpreach) | **Get** /nginx/certificates/test-http | Test HTTP Reachability
*CertificatesAPI* | [**UploadCertificate**](docs/CertificatesAPI.md#uploadcertificate) | **Post** /nginx/certificates/{certID}/upload | Uploads a custom Certificate
*CertificatesAPI* | [**ValidateCertificates**](docs/CertificatesAPI.md#validatecertificates) | **Post** /nginx/certificates/validate | Validates given Custom Certificates
*Class404HostsAPI* | [**Create404Host**](docs/Class404HostsAPI.md#create404host) | **Post** /nginx/dead-hosts | Create a 404 Host
*Class404HostsAPI* | [**DeleteDeadHost**](docs/Class404HostsAPI.md#deletedeadhost) | **Delete** /nginx/dead-hosts/{hostID} | Delete a 404 Host
*Class404HostsAPI* | [**DisableDeadHost**](docs/Class404HostsAPI.md#disabledeadhost) | **Post** /nginx/dead-hosts/{hostID}/disable | Disable a 404 Host
*Class404HostsAPI* | [**EnableDeadHost**](docs/Class404HostsAPI.md#enabledeadhost) | **Post** /nginx/dead-hosts/{hostID}/enable | Enable a 404 Host
*Class404HostsAPI* | [**GetDeadHost**](docs/Class404HostsAPI.md#getdeadhost) | **Get** /nginx/dead-hosts/{hostID} | Get a 404 Host
*Class404HostsAPI* | [**GetDeadHosts**](docs/Class404HostsAPI.md#getdeadhosts) | **Get** /nginx/dead-hosts | Get all 404 hosts
*Class404HostsAPI* | [**UpdateDeadHost**](docs/Class404HostsAPI.md#updatedeadhost) | **Put** /nginx/dead-hosts/{hostID} | Update a 404 Host
*ProxyHostsAPI* | [**CreateProxyHost**](docs/ProxyHostsAPI.md#createproxyhost) | **Post** /nginx/proxy-hosts | Create a Proxy Host
*ProxyHostsAPI* | [**DeleteProxyHost**](docs/ProxyHostsAPI.md#deleteproxyhost) | **Delete** /nginx/proxy-hosts/{hostID} | Delete a Proxy Host
*ProxyHostsAPI* | [**DisableProxyHost**](docs/ProxyHostsAPI.md#disableproxyhost) | **Post** /nginx/proxy-hosts/{hostID}/disable | Disable a Proxy Host
*ProxyHostsAPI* | [**EnableProxyHost**](docs/ProxyHostsAPI.md#enableproxyhost) | **Post** /nginx/proxy-hosts/{hostID}/enable | Enable a Proxy Host
*ProxyHostsAPI* | [**GetProxyHost**](docs/ProxyHostsAPI.md#getproxyhost) | **Get** /nginx/proxy-hosts/{hostID} | Get a Proxy Host
*ProxyHostsAPI* | [**GetProxyHosts**](docs/ProxyHostsAPI.md#getproxyhosts) | **Get** /nginx/proxy-hosts | Get all proxy hosts
*ProxyHostsAPI* | [**UpdateProxyHost**](docs/ProxyHostsAPI.md#updateproxyhost) | **Put** /nginx/proxy-hosts/{hostID} | Update a Proxy Host
*PublicAPI* | [**Health**](docs/PublicAPI.md#health) | **Get** / | Returns the API health status
*PublicAPI* | [**Schema**](docs/PublicAPI.md#schema) | **Get** /schema | Returns this swagger API schema
*RedirectionHostsAPI* | [**CreateRedirectionHost**](docs/RedirectionHostsAPI.md#createredirectionhost) | **Post** /nginx/redirection-hosts | Create a Redirection Host
*RedirectionHostsAPI* | [**DeleteRedirectionHost**](docs/RedirectionHostsAPI.md#deleteredirectionhost) | **Delete** /nginx/redirection-hosts/{hostID} | Delete a Redirection Host
*RedirectionHostsAPI* | [**DisableRedirectionHost**](docs/RedirectionHostsAPI.md#disableredirectionhost) | **Post** /nginx/redirection-hosts/{hostID}/disable | Disable a Redirection Host
*RedirectionHostsAPI* | [**EnableRedirectionHost**](docs/RedirectionHostsAPI.md#enableredirectionhost) | **Post** /nginx/redirection-hosts/{hostID}/enable | Enable a Redirection Host
*RedirectionHostsAPI* | [**GetRedirectionHost**](docs/RedirectionHostsAPI.md#getredirectionhost) | **Get** /nginx/redirection-hosts/{hostID} | Get a Redirection Host
*RedirectionHostsAPI* | [**GetRedirectionHosts**](docs/RedirectionHostsAPI.md#getredirectionhosts) | **Get** /nginx/redirection-hosts | Get all Redirection hosts
*RedirectionHostsAPI* | [**UpdateRedirectionHost**](docs/RedirectionHostsAPI.md#updateredirectionhost) | **Put** /nginx/redirection-hosts/{hostID} | Update a Redirection Host
*ReportsAPI* | [**ReportsHosts**](docs/ReportsAPI.md#reportshosts) | **Get** /reports/hosts | Report on Host Statistics
*SettingsAPI* | [**GetSetting**](docs/SettingsAPI.md#getsetting) | **Get** /settings/{settingID} | Get a setting
*SettingsAPI* | [**GetSettings**](docs/SettingsAPI.md#getsettings) | **Get** /settings | Get all settings
*SettingsAPI* | [**UpdateSetting**](docs/SettingsAPI.md#updatesetting) | **Put** /settings/{settingID} | Update a setting
*StreamsAPI* | [**CreateStream**](docs/StreamsAPI.md#createstream) | **Post** /nginx/streams | Create a Stream
*StreamsAPI* | [**DeleteStream**](docs/StreamsAPI.md#deletestream) | **Delete** /nginx/streams/{streamID} | Delete a Stream
*StreamsAPI* | [**DisableStream**](docs/StreamsAPI.md#disablestream) | **Post** /nginx/streams/{streamID}/disable | Disable a Stream
*StreamsAPI* | [**EnableStream**](docs/StreamsAPI.md#enablestream) | **Post** /nginx/streams/{streamID}/enable | Enable a Stream
*StreamsAPI* | [**GetStream**](docs/StreamsAPI.md#getstream) | **Get** /nginx/streams/{streamID} | Get a Stream
*StreamsAPI* | [**GetStreams**](docs/StreamsAPI.md#getstreams) | **Get** /nginx/streams | Get all streams
*StreamsAPI* | [**UpdateStream**](docs/StreamsAPI.md#updatestream) | **Put** /nginx/streams/{streamID} | Update a Stream
*TokensAPI* | [**RefreshToken**](docs/TokensAPI.md#refreshtoken) | **Get** /tokens | Refresh your access token
*TokensAPI* | [**RequestToken**](docs/TokensAPI.md#requesttoken) | **Post** /tokens | Request a new access token from credentials
*UsersAPI* | [**CreateUser**](docs/UsersAPI.md#createuser) | **Post** /users | Create a User
*UsersAPI* | [**DeleteUser**](docs/UsersAPI.md#deleteuser) | **Delete** /users/{userID} | Delete a User
*UsersAPI* | [**GetUser**](docs/UsersAPI.md#getuser) | **Get** /users/{userID} | Get a user
*UsersAPI* | [**GetUsers**](docs/UsersAPI.md#getusers) | **Get** /users | Get all users
*UsersAPI* | [**LoginAsUser**](docs/UsersAPI.md#loginasuser) | **Post** /users/{userID}/login | Login as this user
*UsersAPI* | [**UpdateUser**](docs/UsersAPI.md#updateuser) | **Put** /users/{userID} | Update a User
*UsersAPI* | [**UpdateUserAuth**](docs/UsersAPI.md#updateuserauth) | **Put** /users/{userID}/auth | Update a User&#39;s Authentication
*UsersAPI* | [**UpdateUserPermissions**](docs/UsersAPI.md#updateuserpermissions) | **Put** /users/{userID}/permissions | Update a User&#39;s Permissions


## Documentation For Models

 - [Create404HostRequest](docs/Create404HostRequest.md)
 - [CreateAccessList201Response](docs/CreateAccessList201Response.md)
 - [CreateAccessListRequest](docs/CreateAccessListRequest.md)
 - [CreateAccessListRequestClientsInner](docs/CreateAccessListRequestClientsInner.md)
 - [CreateAccessListRequestItemsInner](docs/CreateAccessListRequestItemsInner.md)
 - [CreateCertificate400Response](docs/CreateCertificate400Response.md)
 - [CreateCertificate400ResponseError](docs/CreateCertificate400ResponseError.md)
 - [CreateCertificateRequest](docs/CreateCertificateRequest.md)
 - [CreateProxyHostRequest](docs/CreateProxyHostRequest.md)
 - [CreateRedirectionHostRequest](docs/CreateRedirectionHostRequest.md)
 - [CreateStreamRequest](docs/CreateStreamRequest.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [GetAccessLists200ResponseInner](docs/GetAccessLists200ResponseInner.md)
 - [GetAccessLists200ResponseInnerClientsInner](docs/GetAccessLists200ResponseInnerClientsInner.md)
 - [GetAccessLists200ResponseInnerClientsInnerAddress](docs/GetAccessLists200ResponseInnerClientsInnerAddress.md)
 - [GetAccessLists200ResponseInnerItemsInner](docs/GetAccessLists200ResponseInnerItemsInner.md)
 - [GetAccessLists200ResponseInnerOwner](docs/GetAccessLists200ResponseInnerOwner.md)
 - [GetAccessLists200ResponseInnerOwnerPermissions](docs/GetAccessLists200ResponseInnerOwnerPermissions.md)
 - [GetAuditLog200Response](docs/GetAuditLog200Response.md)
 - [GetCertificates200ResponseInner](docs/GetCertificates200ResponseInner.md)
 - [GetCertificates200ResponseInnerMeta](docs/GetCertificates200ResponseInnerMeta.md)
 - [GetDeadHosts200ResponseInner](docs/GetDeadHosts200ResponseInner.md)
 - [GetProxyHosts200ResponseInner](docs/GetProxyHosts200ResponseInner.md)
 - [GetProxyHosts200ResponseInnerAccessList](docs/GetProxyHosts200ResponseInnerAccessList.md)
 - [GetProxyHosts200ResponseInnerCertificate](docs/GetProxyHosts200ResponseInnerCertificate.md)
 - [GetProxyHosts200ResponseInnerCertificateId](docs/GetProxyHosts200ResponseInnerCertificateId.md)
 - [GetProxyHosts200ResponseInnerLocationsInner](docs/GetProxyHosts200ResponseInnerLocationsInner.md)
 - [GetRedirectionHosts200ResponseInner](docs/GetRedirectionHosts200ResponseInner.md)
 - [GetSetting200Response](docs/GetSetting200Response.md)
 - [GetSetting200ResponseValue](docs/GetSetting200ResponseValue.md)
 - [GetSettings200ResponseInner](docs/GetSettings200ResponseInner.md)
 - [GetSettings200ResponseInnerValue](docs/GetSettings200ResponseInnerValue.md)
 - [GetStreams200ResponseInner](docs/GetStreams200ResponseInner.md)
 - [GetStreams200ResponseInnerForwardingHost](docs/GetStreams200ResponseInnerForwardingHost.md)
 - [GetUserUserIDParameter](docs/GetUserUserIDParameter.md)
 - [Health200Response](docs/Health200Response.md)
 - [Health200ResponseVersion](docs/Health200ResponseVersion.md)
 - [LoginAsUser200Response](docs/LoginAsUser200Response.md)
 - [RefreshToken200Response](docs/RefreshToken200Response.md)
 - [ReportsHosts200Response](docs/ReportsHosts200Response.md)
 - [RequestTokenRequest](docs/RequestTokenRequest.md)
 - [UpdateAccessListRequest](docs/UpdateAccessListRequest.md)
 - [UpdateDeadHostRequest](docs/UpdateDeadHostRequest.md)
 - [UpdateProxyHostRequest](docs/UpdateProxyHostRequest.md)
 - [UpdateRedirectionHostRequest](docs/UpdateRedirectionHostRequest.md)
 - [UpdateSettingRequest](docs/UpdateSettingRequest.md)
 - [UpdateSettingRequestMeta](docs/UpdateSettingRequestMeta.md)
 - [UpdateStreamRequest](docs/UpdateStreamRequest.md)
 - [UpdateUserAuthRequest](docs/UpdateUserAuthRequest.md)
 - [UpdateUserRequest](docs/UpdateUserRequest.md)
 - [UploadCertificate200Response](docs/UploadCertificate200Response.md)
 - [ValidateCertificates200Response](docs/ValidateCertificates200Response.md)
 - [ValidateCertificates200ResponseCertificate](docs/ValidateCertificates200ResponseCertificate.md)
 - [ValidateCertificates200ResponseCertificateDates](docs/ValidateCertificates200ResponseCertificateDates.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), nginxproxymanager.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



