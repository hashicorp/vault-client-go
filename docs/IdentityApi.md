# Identity

Method | HTTP request | Description
------------- | ------------- | -------------
[**AliasDeleteByID**](IdentityApi.md#AliasDeleteByID) | **Delete** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**AliasListByID**](IdentityApi.md#AliasListByID) | **Get** /identity/alias/id | List all the alias IDs.
[**AliasReadByID**](IdentityApi.md#AliasReadByID) | **Get** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**AliasWrite**](IdentityApi.md#AliasWrite) | **Post** /identity/alias | Create a new alias.
[**AliasWriteByID**](IdentityApi.md#AliasWriteByID) | **Post** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**EntityBatchDelete**](IdentityApi.md#EntityBatchDelete) | **Post** /identity/entity/batch-delete | Delete all of the entities provided
[**EntityDeleteAliasByID**](IdentityApi.md#EntityDeleteAliasByID) | **Delete** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**EntityDeleteByID**](IdentityApi.md#EntityDeleteByID) | **Delete** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**EntityDeleteByName**](IdentityApi.md#EntityDeleteByName) | **Delete** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**EntityListAliasesByID**](IdentityApi.md#EntityListAliasesByID) | **Get** /identity/entity-alias/id | List all the alias IDs.
[**EntityListByID**](IdentityApi.md#EntityListByID) | **Get** /identity/entity/id | List all the entity IDs
[**EntityListByName**](IdentityApi.md#EntityListByName) | **Get** /identity/entity/name | List all the entity names
[**EntityLookup**](IdentityApi.md#EntityLookup) | **Post** /identity/lookup/entity | Query entities based on various properties.
[**EntityMerge**](IdentityApi.md#EntityMerge) | **Post** /identity/entity/merge | Merge two or more entities together
[**EntityReadAliasByID**](IdentityApi.md#EntityReadAliasByID) | **Get** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**EntityReadByID**](IdentityApi.md#EntityReadByID) | **Get** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**EntityReadByName**](IdentityApi.md#EntityReadByName) | **Get** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**EntityWrite**](IdentityApi.md#EntityWrite) | **Post** /identity/entity | Create a new entity
[**EntityWriteAlias**](IdentityApi.md#EntityWriteAlias) | **Post** /identity/entity-alias | Create a new alias.
[**EntityWriteAliasByID**](IdentityApi.md#EntityWriteAliasByID) | **Post** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**EntityWriteByID**](IdentityApi.md#EntityWriteByID) | **Post** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**EntityWriteByName**](IdentityApi.md#EntityWriteByName) | **Post** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**GroupDeleteAliasByID**](IdentityApi.md#GroupDeleteAliasByID) | **Delete** /identity/group-alias/id/{id} | 
[**GroupDeleteByID**](IdentityApi.md#GroupDeleteByID) | **Delete** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**GroupDeleteByName**](IdentityApi.md#GroupDeleteByName) | **Delete** /identity/group/name/{name} | 
[**GroupListAliasesByID**](IdentityApi.md#GroupListAliasesByID) | **Get** /identity/group-alias/id | List all the group alias IDs.
[**GroupListByID**](IdentityApi.md#GroupListByID) | **Get** /identity/group/id | List all the group IDs.
[**GroupListByName**](IdentityApi.md#GroupListByName) | **Get** /identity/group/name | 
[**GroupLookup**](IdentityApi.md#GroupLookup) | **Post** /identity/lookup/group | Query groups based on various properties.
[**GroupReadAliasByID**](IdentityApi.md#GroupReadAliasByID) | **Get** /identity/group-alias/id/{id} | 
[**GroupReadByID**](IdentityApi.md#GroupReadByID) | **Get** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**GroupReadByName**](IdentityApi.md#GroupReadByName) | **Get** /identity/group/name/{name} | 
[**GroupWrite**](IdentityApi.md#GroupWrite) | **Post** /identity/group | Create a new group.
[**GroupWriteAlias**](IdentityApi.md#GroupWriteAlias) | **Post** /identity/group-alias | Creates a new group alias, or updates an existing one.
[**GroupWriteAliasByID**](IdentityApi.md#GroupWriteAliasByID) | **Post** /identity/group-alias/id/{id} | 
[**GroupWriteByID**](IdentityApi.md#GroupWriteByID) | **Post** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**GroupWriteByName**](IdentityApi.md#GroupWriteByName) | **Post** /identity/group/name/{name} | 
[**MFADeleteLoginEnforcement**](IdentityApi.md#MFADeleteLoginEnforcement) | **Delete** /identity/mfa/login-enforcement/{name} | Delete a login enforcement
[**MFAListLoginEnforcements**](IdentityApi.md#MFAListLoginEnforcements) | **Get** /identity/mfa/login-enforcement | List login enforcements
[**MFAMethodAdminDestroyTOTP**](IdentityApi.md#MFAMethodAdminDestroyTOTP) | **Post** /identity/mfa/method/totp/admin-destroy | Destroys a TOTP secret for the given MFA method ID on the given entity
[**MFAMethodAdminGenerateTOTP**](IdentityApi.md#MFAMethodAdminGenerateTOTP) | **Post** /identity/mfa/method/totp/admin-generate | Update or create TOTP secret for the given method ID on the given entity.
[**MFAMethodDeleteDuo**](IdentityApi.md#MFAMethodDeleteDuo) | **Delete** /identity/mfa/method/duo/{method_id} | Delete a configuration for the given MFA method
[**MFAMethodDeleteOkta**](IdentityApi.md#MFAMethodDeleteOkta) | **Delete** /identity/mfa/method/okta/{method_id} | Delete a configuration for the given MFA method
[**MFAMethodDeletePingID**](IdentityApi.md#MFAMethodDeletePingID) | **Delete** /identity/mfa/method/pingid/{method_id} | Delete a configuration for the given MFA method
[**MFAMethodDeleteTOTP**](IdentityApi.md#MFAMethodDeleteTOTP) | **Delete** /identity/mfa/method/totp/{method_id} | Delete a configuration for the given MFA method
[**MFAMethodGenerateTOTP**](IdentityApi.md#MFAMethodGenerateTOTP) | **Post** /identity/mfa/method/totp/generate | Update or create TOTP secret for the given method ID on the given entity.
[**MFAMethodList**](IdentityApi.md#MFAMethodList) | **Get** /identity/mfa/method | List MFA method configurations for all MFA methods
[**MFAMethodListDuo**](IdentityApi.md#MFAMethodListDuo) | **Get** /identity/mfa/method/duo | List MFA method configurations for the given MFA method
[**MFAMethodListOkta**](IdentityApi.md#MFAMethodListOkta) | **Get** /identity/mfa/method/okta | List MFA method configurations for the given MFA method
[**MFAMethodListPingID**](IdentityApi.md#MFAMethodListPingID) | **Get** /identity/mfa/method/pingid | List MFA method configurations for the given MFA method
[**MFAMethodListTOTP**](IdentityApi.md#MFAMethodListTOTP) | **Get** /identity/mfa/method/totp | List MFA method configurations for the given MFA method
[**MFAMethodRead**](IdentityApi.md#MFAMethodRead) | **Get** /identity/mfa/method/{method_id} | Read the current configuration for the given ID regardless of the MFA method type
[**MFAMethodReadDuo**](IdentityApi.md#MFAMethodReadDuo) | **Get** /identity/mfa/method/duo/{method_id} | Read the current configuration for the given MFA method
[**MFAMethodReadOkta**](IdentityApi.md#MFAMethodReadOkta) | **Get** /identity/mfa/method/okta/{method_id} | Read the current configuration for the given MFA method
[**MFAMethodReadPingID**](IdentityApi.md#MFAMethodReadPingID) | **Get** /identity/mfa/method/pingid/{method_id} | Read the current configuration for the given MFA method
[**MFAMethodReadTOTP**](IdentityApi.md#MFAMethodReadTOTP) | **Get** /identity/mfa/method/totp/{method_id} | Read the current configuration for the given MFA method
[**MFAMethodWriteDuo**](IdentityApi.md#MFAMethodWriteDuo) | **Post** /identity/mfa/method/duo/{method_id} | Update or create a configuration for the given MFA method
[**MFAMethodWriteOkta**](IdentityApi.md#MFAMethodWriteOkta) | **Post** /identity/mfa/method/okta/{method_id} | Update or create a configuration for the given MFA method
[**MFAMethodWritePingID**](IdentityApi.md#MFAMethodWritePingID) | **Post** /identity/mfa/method/pingid/{method_id} | Update or create a configuration for the given MFA method
[**MFAMethodWriteTOTP**](IdentityApi.md#MFAMethodWriteTOTP) | **Post** /identity/mfa/method/totp/{method_id} | Update or create a configuration for the given MFA method
[**MFAReadLoginEnforcement**](IdentityApi.md#MFAReadLoginEnforcement) | **Get** /identity/mfa/login-enforcement/{name} | Read the current login enforcement
[**MFAWriteLoginEnforcement**](IdentityApi.md#MFAWriteLoginEnforcement) | **Post** /identity/mfa/login-enforcement/{name} | Create or update a login enforcement
[**OIDCDeleteAssignment**](IdentityApi.md#OIDCDeleteAssignment) | **Delete** /identity/oidc/assignment/{name} | 
[**OIDCDeleteClient**](IdentityApi.md#OIDCDeleteClient) | **Delete** /identity/oidc/client/{name} | 
[**OIDCDeleteKey**](IdentityApi.md#OIDCDeleteKey) | **Delete** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**OIDCDeleteProvider**](IdentityApi.md#OIDCDeleteProvider) | **Delete** /identity/oidc/provider/{name} | 
[**OIDCDeleteRole**](IdentityApi.md#OIDCDeleteRole) | **Delete** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**OIDCDeleteScope**](IdentityApi.md#OIDCDeleteScope) | **Delete** /identity/oidc/scope/{name} | 
[**OIDCIntrospect**](IdentityApi.md#OIDCIntrospect) | **Post** /identity/oidc/introspect | Verify the authenticity of an OIDC token
[**OIDCListAssignments**](IdentityApi.md#OIDCListAssignments) | **Get** /identity/oidc/assignment | 
[**OIDCListClients**](IdentityApi.md#OIDCListClients) | **Get** /identity/oidc/client | 
[**OIDCListKeys**](IdentityApi.md#OIDCListKeys) | **Get** /identity/oidc/key | List OIDC keys
[**OIDCListProviders**](IdentityApi.md#OIDCListProviders) | **Get** /identity/oidc/provider | 
[**OIDCListRoles**](IdentityApi.md#OIDCListRoles) | **Get** /identity/oidc/role | List configured OIDC roles
[**OIDCListScopes**](IdentityApi.md#OIDCListScopes) | **Get** /identity/oidc/scope | 
[**OIDCReadAssignment**](IdentityApi.md#OIDCReadAssignment) | **Get** /identity/oidc/assignment/{name} | 
[**OIDCReadClient**](IdentityApi.md#OIDCReadClient) | **Get** /identity/oidc/client/{name} | 
[**OIDCReadConfig**](IdentityApi.md#OIDCReadConfig) | **Get** /identity/oidc/config | OIDC configuration
[**OIDCReadKey**](IdentityApi.md#OIDCReadKey) | **Get** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**OIDCReadProvider**](IdentityApi.md#OIDCReadProvider) | **Get** /identity/oidc/provider/{name} | 
[**OIDCReadProviderAuthorize**](IdentityApi.md#OIDCReadProviderAuthorize) | **Get** /identity/oidc/provider/{name}/authorize | 
[**OIDCReadProviderUserInfo**](IdentityApi.md#OIDCReadProviderUserInfo) | **Get** /identity/oidc/provider/{name}/userinfo | 
[**OIDCReadProviderWellKnownKeys**](IdentityApi.md#OIDCReadProviderWellKnownKeys) | **Get** /identity/oidc/provider/{name}/.well-known/keys | 
[**OIDCReadProviderWellKnownOpenIDConfiguration**](IdentityApi.md#OIDCReadProviderWellKnownOpenIDConfiguration) | **Get** /identity/oidc/provider/{name}/.well-known/openid-configuration | 
[**OIDCReadRole**](IdentityApi.md#OIDCReadRole) | **Get** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**OIDCReadScope**](IdentityApi.md#OIDCReadScope) | **Get** /identity/oidc/scope/{name} | 
[**OIDCReadToken**](IdentityApi.md#OIDCReadToken) | **Get** /identity/oidc/token/{name} | Generate an OIDC token
[**OIDCReadWellKnownKeys**](IdentityApi.md#OIDCReadWellKnownKeys) | **Get** /identity/oidc/.well-known/keys | Retrieve public keys
[**OIDCReadWellKnownOpenIDConfiguration**](IdentityApi.md#OIDCReadWellKnownOpenIDConfiguration) | **Get** /identity/oidc/.well-known/openid-configuration | Query OIDC configurations
[**OIDCRotateKey**](IdentityApi.md#OIDCRotateKey) | **Post** /identity/oidc/key/{name}/rotate | Rotate a named OIDC key.
[**OIDCWriteAssignment**](IdentityApi.md#OIDCWriteAssignment) | **Post** /identity/oidc/assignment/{name} | 
[**OIDCWriteClient**](IdentityApi.md#OIDCWriteClient) | **Post** /identity/oidc/client/{name} | 
[**OIDCWriteConfig**](IdentityApi.md#OIDCWriteConfig) | **Post** /identity/oidc/config | OIDC configuration
[**OIDCWriteKey**](IdentityApi.md#OIDCWriteKey) | **Post** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**OIDCWriteProvider**](IdentityApi.md#OIDCWriteProvider) | **Post** /identity/oidc/provider/{name} | 
[**OIDCWriteProviderAuthorize**](IdentityApi.md#OIDCWriteProviderAuthorize) | **Post** /identity/oidc/provider/{name}/authorize | 
[**OIDCWriteProviderToken**](IdentityApi.md#OIDCWriteProviderToken) | **Post** /identity/oidc/provider/{name}/token | 
[**OIDCWriteProviderUserInfo**](IdentityApi.md#OIDCWriteProviderUserInfo) | **Post** /identity/oidc/provider/{name}/userinfo | 
[**OIDCWriteRole**](IdentityApi.md#OIDCWriteRole) | **Post** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**OIDCWriteScope**](IdentityApi.md#OIDCWriteScope) | **Post** /identity/oidc/scope/{name} | 
[**PersonaIDDeleteByID**](IdentityApi.md#PersonaIDDeleteByID) | **Delete** /identity/persona/id/{id} | Update, read or delete an alias ID.
[**PersonaIDReadByID**](IdentityApi.md#PersonaIDReadByID) | **Get** /identity/persona/id/{id} | Update, read or delete an alias ID.
[**PersonaIDWriteByID**](IdentityApi.md#PersonaIDWriteByID) | **Post** /identity/persona/id/{id} | Update, read or delete an alias ID.
[**PersonaListByID**](IdentityApi.md#PersonaListByID) | **Get** /identity/persona/id | List all the alias IDs.
[**PersonaWrite**](IdentityApi.md#PersonaWrite) | **Post** /identity/persona | Create a new alias.





## AliasDeleteByID

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the alias
	resp, err := client.Identity.AliasDeleteByID(
		context.Background(),
		id,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the alias | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliasListByID

List all the alias IDs.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.AliasListByID(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliasReadByID

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the alias
	resp, err := client.Identity.AliasReadByID(
		context.Background(),
		id,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the alias | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliasWrite

Create a new alias.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewAliasWriteRequestWithDefaults()
	resp, err := client.Identity.AliasWrite(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aliasWriteRequest** | [**AliasWriteRequest**](AliasWriteRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## AliasWriteByID

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the alias
	request := schema.NewAliasWriteByIDRequestWithDefaults()
	resp, err := client.Identity.AliasWriteByID(
		context.Background(),
		id,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the alias | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aliasWriteByIDRequest** | [**AliasWriteByIDRequest**](AliasWriteByIDRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityBatchDelete

Delete all of the entities provided

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewEntityBatchDeleteRequestWithDefaults()
	resp, err := client.Identity.EntityBatchDelete(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityBatchDeleteRequest** | [**EntityBatchDeleteRequest**](EntityBatchDeleteRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityDeleteAliasByID

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the alias
	resp, err := client.Identity.EntityDeleteAliasByID(
		context.Background(),
		id,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the alias | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityDeleteByID

Update, read or delete an entity using entity ID

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the entity. If set, updates the corresponding existing entity.
	resp, err := client.Identity.EntityDeleteByID(
		context.Background(),
		id,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the entity. If set, updates the corresponding existing entity. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityDeleteByName

Update, read or delete an entity using entity name

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the entity
	resp, err := client.Identity.EntityDeleteByName(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the entity | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityListAliasesByID

List all the alias IDs.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.EntityListAliasesByID(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityListByID

List all the entity IDs

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.EntityListByID(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityListByName

List all the entity names

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.EntityListByName(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityLookup

Query entities based on various properties.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewEntityLookupRequestWithDefaults()
	resp, err := client.Identity.EntityLookup(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityLookupRequest** | [**EntityLookupRequest**](EntityLookupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityMerge

Merge two or more entities together

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewEntityMergeRequestWithDefaults()
	resp, err := client.Identity.EntityMerge(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityMergeRequest** | [**EntityMergeRequest**](EntityMergeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityReadAliasByID

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the alias
	resp, err := client.Identity.EntityReadAliasByID(
		context.Background(),
		id,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the alias | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityReadByID

Update, read or delete an entity using entity ID

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the entity. If set, updates the corresponding existing entity.
	resp, err := client.Identity.EntityReadByID(
		context.Background(),
		id,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the entity. If set, updates the corresponding existing entity. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityReadByName

Update, read or delete an entity using entity name

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the entity
	resp, err := client.Identity.EntityReadByName(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the entity | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityWrite

Create a new entity

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewEntityWriteRequestWithDefaults()
	resp, err := client.Identity.EntityWrite(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityWriteRequest** | [**EntityWriteRequest**](EntityWriteRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityWriteAlias

Create a new alias.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewEntityWriteAliasRequestWithDefaults()
	resp, err := client.Identity.EntityWriteAlias(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityWriteAliasRequest** | [**EntityWriteAliasRequest**](EntityWriteAliasRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityWriteAliasByID

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the alias
	request := schema.NewEntityWriteAliasByIDRequestWithDefaults()
	resp, err := client.Identity.EntityWriteAliasByID(
		context.Background(),
		id,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the alias | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityWriteAliasByIDRequest** | [**EntityWriteAliasByIDRequest**](EntityWriteAliasByIDRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityWriteByID

Update, read or delete an entity using entity ID

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the entity. If set, updates the corresponding existing entity.
	request := schema.NewEntityWriteByIDRequestWithDefaults()
	resp, err := client.Identity.EntityWriteByID(
		context.Background(),
		id,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the entity. If set, updates the corresponding existing entity. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityWriteByIDRequest** | [**EntityWriteByIDRequest**](EntityWriteByIDRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## EntityWriteByName

Update, read or delete an entity using entity name

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the entity
	request := schema.NewEntityWriteByNameRequestWithDefaults()
	resp, err := client.Identity.EntityWriteByName(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the entity | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityWriteByNameRequest** | [**EntityWriteByNameRequest**](EntityWriteByNameRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupDeleteAliasByID



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the group alias.
	resp, err := client.Identity.GroupDeleteAliasByID(
		context.Background(),
		id,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the group alias. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupDeleteByID

Update or delete an existing group using its ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the group. If set, updates the corresponding existing group.
	resp, err := client.Identity.GroupDeleteByID(
		context.Background(),
		id,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the group. If set, updates the corresponding existing group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupDeleteByName



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the group.
	resp, err := client.Identity.GroupDeleteByName(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupListAliasesByID

List all the group alias IDs.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.GroupListAliasesByID(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupListByID

List all the group IDs.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.GroupListByID(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupListByName



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.GroupListByName(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupLookup

Query groups based on various properties.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewGroupLookupRequestWithDefaults()
	resp, err := client.Identity.GroupLookup(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupLookupRequest** | [**GroupLookupRequest**](GroupLookupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupReadAliasByID



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the group alias.
	resp, err := client.Identity.GroupReadAliasByID(
		context.Background(),
		id,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the group alias. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupReadByID

Update or delete an existing group using its ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the group. If set, updates the corresponding existing group.
	resp, err := client.Identity.GroupReadByID(
		context.Background(),
		id,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the group. If set, updates the corresponding existing group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupReadByName



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the group.
	resp, err := client.Identity.GroupReadByName(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupWrite

Create a new group.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewGroupWriteRequestWithDefaults()
	resp, err := client.Identity.GroupWrite(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupWriteRequest** | [**GroupWriteRequest**](GroupWriteRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupWriteAlias

Creates a new group alias, or updates an existing one.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewGroupWriteAliasRequestWithDefaults()
	resp, err := client.Identity.GroupWriteAlias(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupWriteAliasRequest** | [**GroupWriteAliasRequest**](GroupWriteAliasRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupWriteAliasByID



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the group alias.
	request := schema.NewGroupWriteAliasByIDRequestWithDefaults()
	resp, err := client.Identity.GroupWriteAliasByID(
		context.Background(),
		id,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the group alias. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupWriteAliasByIDRequest** | [**GroupWriteAliasByIDRequest**](GroupWriteAliasByIDRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupWriteByID

Update or delete an existing group using its ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the group. If set, updates the corresponding existing group.
	request := schema.NewGroupWriteByIDRequestWithDefaults()
	resp, err := client.Identity.GroupWriteByID(
		context.Background(),
		id,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the group. If set, updates the corresponding existing group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupWriteByIDRequest** | [**GroupWriteByIDRequest**](GroupWriteByIDRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## GroupWriteByName



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the group.
	request := schema.NewGroupWriteByNameRequestWithDefaults()
	resp, err := client.Identity.GroupWriteByName(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupWriteByNameRequest** | [**GroupWriteByNameRequest**](GroupWriteByNameRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFADeleteLoginEnforcement

Delete a login enforcement

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name for this login enforcement configuration
	resp, err := client.Identity.MFADeleteLoginEnforcement(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name for this login enforcement configuration | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAListLoginEnforcements

List login enforcements

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.MFAListLoginEnforcements(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodAdminDestroyTOTP

Destroys a TOTP secret for the given MFA method ID on the given entity

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewMFAMethodAdminDestroyTOTPRequestWithDefaults()
	resp, err := client.Identity.MFAMethodAdminDestroyTOTP(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mFAMethodAdminDestroyTOTPRequest** | [**MFAMethodAdminDestroyTOTPRequest**](MFAMethodAdminDestroyTOTPRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodAdminGenerateTOTP

Update or create TOTP secret for the given method ID on the given entity.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewMFAMethodAdminGenerateTOTPRequestWithDefaults()
	resp, err := client.Identity.MFAMethodAdminGenerateTOTP(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mFAMethodAdminGenerateTOTPRequest** | [**MFAMethodAdminGenerateTOTPRequest**](MFAMethodAdminGenerateTOTPRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodDeleteDuo

Delete a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MFAMethodDeleteDuo(
		context.Background(),
		methodId,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodDeleteOkta

Delete a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MFAMethodDeleteOkta(
		context.Background(),
		methodId,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodDeletePingID

Delete a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MFAMethodDeletePingID(
		context.Background(),
		methodId,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodDeleteTOTP

Delete a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MFAMethodDeleteTOTP(
		context.Background(),
		methodId,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodGenerateTOTP

Update or create TOTP secret for the given method ID on the given entity.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewMFAMethodGenerateTOTPRequestWithDefaults()
	resp, err := client.Identity.MFAMethodGenerateTOTP(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mFAMethodGenerateTOTPRequest** | [**MFAMethodGenerateTOTPRequest**](MFAMethodGenerateTOTPRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodList

List MFA method configurations for all MFA methods

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.MFAMethodList(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodListDuo

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.MFAMethodListDuo(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodListOkta

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.MFAMethodListOkta(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodListPingID

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.MFAMethodListPingID(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodListTOTP

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.MFAMethodListTOTP(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodRead

Read the current configuration for the given ID regardless of the MFA method type

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MFAMethodRead(
		context.Background(),
		methodId,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodReadDuo

Read the current configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MFAMethodReadDuo(
		context.Background(),
		methodId,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodReadOkta

Read the current configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MFAMethodReadOkta(
		context.Background(),
		methodId,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodReadPingID

Read the current configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MFAMethodReadPingID(
		context.Background(),
		methodId,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodReadTOTP

Read the current configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MFAMethodReadTOTP(
		context.Background(),
		methodId,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodWriteDuo

Update or create a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	request := schema.NewMFAMethodWriteDuoRequestWithDefaults()
	resp, err := client.Identity.MFAMethodWriteDuo(
		context.Background(),
		methodId,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mFAMethodWriteDuoRequest** | [**MFAMethodWriteDuoRequest**](MFAMethodWriteDuoRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodWriteOkta

Update or create a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	request := schema.NewMFAMethodWriteOktaRequestWithDefaults()
	resp, err := client.Identity.MFAMethodWriteOkta(
		context.Background(),
		methodId,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mFAMethodWriteOktaRequest** | [**MFAMethodWriteOktaRequest**](MFAMethodWriteOktaRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodWritePingID

Update or create a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	request := schema.NewMFAMethodWritePingIDRequestWithDefaults()
	resp, err := client.Identity.MFAMethodWritePingID(
		context.Background(),
		methodId,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mFAMethodWritePingIDRequest** | [**MFAMethodWritePingIDRequest**](MFAMethodWritePingIDRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAMethodWriteTOTP

Update or create a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	request := schema.NewMFAMethodWriteTOTPRequestWithDefaults()
	resp, err := client.Identity.MFAMethodWriteTOTP(
		context.Background(),
		methodId,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mFAMethodWriteTOTPRequest** | [**MFAMethodWriteTOTPRequest**](MFAMethodWriteTOTPRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAReadLoginEnforcement

Read the current login enforcement

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name for this login enforcement configuration
	resp, err := client.Identity.MFAReadLoginEnforcement(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name for this login enforcement configuration | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## MFAWriteLoginEnforcement

Create or update a login enforcement

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name for this login enforcement configuration
	request := schema.NewMFAWriteLoginEnforcementRequestWithDefaults()
	resp, err := client.Identity.MFAWriteLoginEnforcement(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name for this login enforcement configuration | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mFAWriteLoginEnforcementRequest** | [**MFAWriteLoginEnforcementRequest**](MFAWriteLoginEnforcementRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCDeleteAssignment



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the assignment
	resp, err := client.Identity.OIDCDeleteAssignment(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the assignment | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCDeleteClient



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the client.
	resp, err := client.Identity.OIDCDeleteClient(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the client. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCDeleteKey

CRUD operations for OIDC keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	resp, err := client.Identity.OIDCDeleteKey(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCDeleteProvider



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OIDCDeleteProvider(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCDeleteRole

CRUD operations on OIDC Roles

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	resp, err := client.Identity.OIDCDeleteRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCDeleteScope



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the scope
	resp, err := client.Identity.OIDCDeleteScope(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the scope | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCIntrospect

Verify the authenticity of an OIDC token

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewOIDCIntrospectRequestWithDefaults()
	resp, err := client.Identity.OIDCIntrospect(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oIDCIntrospectRequest** | [**OIDCIntrospectRequest**](OIDCIntrospectRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCListAssignments



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.OIDCListAssignments(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCListClients



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.OIDCListClients(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCListKeys

List OIDC keys

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.OIDCListKeys(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCListProviders



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	allowedClientId := "allowedClientId_example" // string | Filters the list of OIDC providers to those that allow the given client ID in their set of allowed_client_ids. (defaults to "")
	resp, err := client.Identity.OIDCListProviders(
		context.Background(),
		allowedClientId,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 
 **allowedClientId** | **string** | Filters the list of OIDC providers to those that allow the given client ID in their set of allowed_client_ids. | [default to &quot;&quot;]

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCListRoles

List configured OIDC roles

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.OIDCListRoles(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCListScopes



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.OIDCListScopes(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadAssignment



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the assignment
	resp, err := client.Identity.OIDCReadAssignment(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the assignment | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadClient



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the client.
	resp, err := client.Identity.OIDCReadClient(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the client. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadConfig

OIDC configuration

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.OIDCReadConfig(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadKey

CRUD operations for OIDC keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	resp, err := client.Identity.OIDCReadKey(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadProvider



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OIDCReadProvider(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadProviderAuthorize



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OIDCReadProviderAuthorize(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadProviderUserInfo



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OIDCReadProviderUserInfo(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadProviderWellKnownKeys



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OIDCReadProviderWellKnownKeys(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadProviderWellKnownOpenIDConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OIDCReadProviderWellKnownOpenIDConfiguration(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadRole

CRUD operations on OIDC Roles

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	resp, err := client.Identity.OIDCReadRole(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadScope



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the scope
	resp, err := client.Identity.OIDCReadScope(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the scope | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadToken

Generate an OIDC token

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	resp, err := client.Identity.OIDCReadToken(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadWellKnownKeys

Retrieve public keys

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.OIDCReadWellKnownKeys(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCReadWellKnownOpenIDConfiguration

Query OIDC configurations

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.OIDCReadWellKnownOpenIDConfiguration(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCRotateKey

Rotate a named OIDC key.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	request := schema.NewOIDCRotateKeyRequestWithDefaults()
	resp, err := client.Identity.OIDCRotateKey(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oIDCRotateKeyRequest** | [**OIDCRotateKeyRequest**](OIDCRotateKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteAssignment



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the assignment
	request := schema.NewOIDCWriteAssignmentRequestWithDefaults()
	resp, err := client.Identity.OIDCWriteAssignment(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the assignment | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oIDCWriteAssignmentRequest** | [**OIDCWriteAssignmentRequest**](OIDCWriteAssignmentRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteClient



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the client.
	request := schema.NewOIDCWriteClientRequestWithDefaults()
	resp, err := client.Identity.OIDCWriteClient(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the client. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oIDCWriteClientRequest** | [**OIDCWriteClientRequest**](OIDCWriteClientRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteConfig

OIDC configuration

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewOIDCWriteConfigRequestWithDefaults()
	resp, err := client.Identity.OIDCWriteConfig(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oIDCWriteConfigRequest** | [**OIDCWriteConfigRequest**](OIDCWriteConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteKey

CRUD operations for OIDC keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the key
	request := schema.NewOIDCWriteKeyRequestWithDefaults()
	resp, err := client.Identity.OIDCWriteKey(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oIDCWriteKeyRequest** | [**OIDCWriteKeyRequest**](OIDCWriteKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteProvider



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the provider
	request := schema.NewOIDCWriteProviderRequestWithDefaults()
	resp, err := client.Identity.OIDCWriteProvider(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oIDCWriteProviderRequest** | [**OIDCWriteProviderRequest**](OIDCWriteProviderRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteProviderAuthorize



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the provider
	request := schema.NewOIDCWriteProviderAuthorizeRequestWithDefaults()
	resp, err := client.Identity.OIDCWriteProviderAuthorize(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oIDCWriteProviderAuthorizeRequest** | [**OIDCWriteProviderAuthorizeRequest**](OIDCWriteProviderAuthorizeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteProviderToken



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the provider
	request := schema.NewOIDCWriteProviderTokenRequestWithDefaults()
	resp, err := client.Identity.OIDCWriteProviderToken(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oIDCWriteProviderTokenRequest** | [**OIDCWriteProviderTokenRequest**](OIDCWriteProviderTokenRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteProviderUserInfo



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OIDCWriteProviderUserInfo(
		context.Background(),
		name,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteRole

CRUD operations on OIDC Roles

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the role
	request := schema.NewOIDCWriteRoleRequestWithDefaults()
	resp, err := client.Identity.OIDCWriteRole(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oIDCWriteRoleRequest** | [**OIDCWriteRoleRequest**](OIDCWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## OIDCWriteScope



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	name := "name_example" // string | Name of the scope
	request := schema.NewOIDCWriteScopeRequestWithDefaults()
	resp, err := client.Identity.OIDCWriteScope(
		context.Background(),
		name,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the scope | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oIDCWriteScopeRequest** | [**OIDCWriteScopeRequest**](OIDCWriteScopeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PersonaIDDeleteByID

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the persona
	resp, err := client.Identity.PersonaIDDeleteByID(
		context.Background(),
		id,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the persona | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PersonaIDReadByID

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the persona
	resp, err := client.Identity.PersonaIDReadByID(
		context.Background(),
		id,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the persona | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PersonaIDWriteByID

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	id := "id_example" // string | ID of the persona
	request := schema.NewPersonaIDWriteByIDRequestWithDefaults()
	resp, err := client.Identity.PersonaIDWriteByID(
		context.Background(),
		id,
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the persona | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **personaIDWriteByIDRequest** | [**PersonaIDWriteByIDRequest**](PersonaIDWriteByIDRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PersonaListByID

List all the alias IDs.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Identity.PersonaListByID(
		context.Background(),
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



## PersonaWrite

Create a new alias.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

	request := schema.NewPersonaWriteRequestWithDefaults()
	resp, err := client.Identity.PersonaWrite(
		context.Background(),
		request,
		vault.WithToken("my-token"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personaWriteRequest** | [**PersonaWriteRequest**](PersonaWriteRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)



