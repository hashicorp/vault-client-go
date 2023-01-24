# Identity

Method | HTTP request | Description
------------- | ------------- | -------------
[**AliasDeleteByID**](Identity.md#AliasDeleteByID) | **Delete** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**AliasListByID**](Identity.md#AliasListByID) | **Get** /identity/alias/id | List all the alias IDs.
[**AliasReadByID**](Identity.md#AliasReadByID) | **Get** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**AliasWrite**](Identity.md#AliasWrite) | **Post** /identity/alias | Create a new alias.
[**AliasWriteByID**](Identity.md#AliasWriteByID) | **Post** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**EntityBatchDelete**](Identity.md#EntityBatchDelete) | **Post** /identity/entity/batch-delete | Delete all of the entities provided
[**EntityDeleteAliasByID**](Identity.md#EntityDeleteAliasByID) | **Delete** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**EntityDeleteByID**](Identity.md#EntityDeleteByID) | **Delete** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**EntityDeleteByName**](Identity.md#EntityDeleteByName) | **Delete** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**EntityListAliasesByID**](Identity.md#EntityListAliasesByID) | **Get** /identity/entity-alias/id | List all the alias IDs.
[**EntityListByID**](Identity.md#EntityListByID) | **Get** /identity/entity/id | List all the entity IDs
[**EntityListByName**](Identity.md#EntityListByName) | **Get** /identity/entity/name | List all the entity names
[**EntityLookup**](Identity.md#EntityLookup) | **Post** /identity/lookup/entity | Query entities based on various properties.
[**EntityMerge**](Identity.md#EntityMerge) | **Post** /identity/entity/merge | Merge two or more entities together
[**EntityReadAliasByID**](Identity.md#EntityReadAliasByID) | **Get** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**EntityReadByID**](Identity.md#EntityReadByID) | **Get** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**EntityReadByName**](Identity.md#EntityReadByName) | **Get** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**EntityWrite**](Identity.md#EntityWrite) | **Post** /identity/entity | Create a new entity
[**EntityWriteAlias**](Identity.md#EntityWriteAlias) | **Post** /identity/entity-alias | Create a new alias.
[**EntityWriteAliasByID**](Identity.md#EntityWriteAliasByID) | **Post** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**EntityWriteByID**](Identity.md#EntityWriteByID) | **Post** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**EntityWriteByName**](Identity.md#EntityWriteByName) | **Post** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**GroupDeleteAliasByID**](Identity.md#GroupDeleteAliasByID) | **Delete** /identity/group-alias/id/{id} | 
[**GroupDeleteByID**](Identity.md#GroupDeleteByID) | **Delete** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**GroupDeleteByName**](Identity.md#GroupDeleteByName) | **Delete** /identity/group/name/{name} | 
[**GroupListAliasesByID**](Identity.md#GroupListAliasesByID) | **Get** /identity/group-alias/id | List all the group alias IDs.
[**GroupListByID**](Identity.md#GroupListByID) | **Get** /identity/group/id | List all the group IDs.
[**GroupListByName**](Identity.md#GroupListByName) | **Get** /identity/group/name | 
[**GroupLookup**](Identity.md#GroupLookup) | **Post** /identity/lookup/group | Query groups based on various properties.
[**GroupReadAliasByID**](Identity.md#GroupReadAliasByID) | **Get** /identity/group-alias/id/{id} | 
[**GroupReadByID**](Identity.md#GroupReadByID) | **Get** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**GroupReadByName**](Identity.md#GroupReadByName) | **Get** /identity/group/name/{name} | 
[**GroupWrite**](Identity.md#GroupWrite) | **Post** /identity/group | Create a new group.
[**GroupWriteAlias**](Identity.md#GroupWriteAlias) | **Post** /identity/group-alias | Creates a new group alias, or updates an existing one.
[**GroupWriteAliasByID**](Identity.md#GroupWriteAliasByID) | **Post** /identity/group-alias/id/{id} | 
[**GroupWriteByID**](Identity.md#GroupWriteByID) | **Post** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**GroupWriteByName**](Identity.md#GroupWriteByName) | **Post** /identity/group/name/{name} | 
[**MFADeleteLoginEnforcement**](Identity.md#MFADeleteLoginEnforcement) | **Delete** /identity/mfa/login-enforcement/{name} | Delete a login enforcement
[**MFAListLoginEnforcements**](Identity.md#MFAListLoginEnforcements) | **Get** /identity/mfa/login-enforcement | List login enforcements
[**MFAMethodAdminDestroyTOTP**](Identity.md#MFAMethodAdminDestroyTOTP) | **Post** /identity/mfa/method/totp/admin-destroy | Destroys a TOTP secret for the given MFA method ID on the given entity
[**MFAMethodAdminGenerateTOTP**](Identity.md#MFAMethodAdminGenerateTOTP) | **Post** /identity/mfa/method/totp/admin-generate | Update or create TOTP secret for the given method ID on the given entity.
[**MFAMethodDeleteDuo**](Identity.md#MFAMethodDeleteDuo) | **Delete** /identity/mfa/method/duo/{method_id} | Delete a configuration for the given MFA method
[**MFAMethodDeleteOkta**](Identity.md#MFAMethodDeleteOkta) | **Delete** /identity/mfa/method/okta/{method_id} | Delete a configuration for the given MFA method
[**MFAMethodDeletePingID**](Identity.md#MFAMethodDeletePingID) | **Delete** /identity/mfa/method/pingid/{method_id} | Delete a configuration for the given MFA method
[**MFAMethodDeleteTOTP**](Identity.md#MFAMethodDeleteTOTP) | **Delete** /identity/mfa/method/totp/{method_id} | Delete a configuration for the given MFA method
[**MFAMethodGenerateTOTP**](Identity.md#MFAMethodGenerateTOTP) | **Post** /identity/mfa/method/totp/generate | Update or create TOTP secret for the given method ID on the given entity.
[**MFAMethodList**](Identity.md#MFAMethodList) | **Get** /identity/mfa/method | List MFA method configurations for all MFA methods
[**MFAMethodListDuo**](Identity.md#MFAMethodListDuo) | **Get** /identity/mfa/method/duo | List MFA method configurations for the given MFA method
[**MFAMethodListOkta**](Identity.md#MFAMethodListOkta) | **Get** /identity/mfa/method/okta | List MFA method configurations for the given MFA method
[**MFAMethodListPingID**](Identity.md#MFAMethodListPingID) | **Get** /identity/mfa/method/pingid | List MFA method configurations for the given MFA method
[**MFAMethodListTOTP**](Identity.md#MFAMethodListTOTP) | **Get** /identity/mfa/method/totp | List MFA method configurations for the given MFA method
[**MFAMethodRead**](Identity.md#MFAMethodRead) | **Get** /identity/mfa/method/{method_id} | Read the current configuration for the given ID regardless of the MFA method type
[**MFAMethodReadDuo**](Identity.md#MFAMethodReadDuo) | **Get** /identity/mfa/method/duo/{method_id} | Read the current configuration for the given MFA method
[**MFAMethodReadOkta**](Identity.md#MFAMethodReadOkta) | **Get** /identity/mfa/method/okta/{method_id} | Read the current configuration for the given MFA method
[**MFAMethodReadPingID**](Identity.md#MFAMethodReadPingID) | **Get** /identity/mfa/method/pingid/{method_id} | Read the current configuration for the given MFA method
[**MFAMethodReadTOTP**](Identity.md#MFAMethodReadTOTP) | **Get** /identity/mfa/method/totp/{method_id} | Read the current configuration for the given MFA method
[**MFAMethodWriteDuo**](Identity.md#MFAMethodWriteDuo) | **Post** /identity/mfa/method/duo/{method_id} | Update or create a configuration for the given MFA method
[**MFAMethodWriteOkta**](Identity.md#MFAMethodWriteOkta) | **Post** /identity/mfa/method/okta/{method_id} | Update or create a configuration for the given MFA method
[**MFAMethodWritePingID**](Identity.md#MFAMethodWritePingID) | **Post** /identity/mfa/method/pingid/{method_id} | Update or create a configuration for the given MFA method
[**MFAMethodWriteTOTP**](Identity.md#MFAMethodWriteTOTP) | **Post** /identity/mfa/method/totp/{method_id} | Update or create a configuration for the given MFA method
[**MFAReadLoginEnforcement**](Identity.md#MFAReadLoginEnforcement) | **Get** /identity/mfa/login-enforcement/{name} | Read the current login enforcement
[**MFAWriteLoginEnforcement**](Identity.md#MFAWriteLoginEnforcement) | **Post** /identity/mfa/login-enforcement/{name} | Create or update a login enforcement
[**OIDCDeleteAssignment**](Identity.md#OIDCDeleteAssignment) | **Delete** /identity/oidc/assignment/{name} | 
[**OIDCDeleteClient**](Identity.md#OIDCDeleteClient) | **Delete** /identity/oidc/client/{name} | 
[**OIDCDeleteKey**](Identity.md#OIDCDeleteKey) | **Delete** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**OIDCDeleteProvider**](Identity.md#OIDCDeleteProvider) | **Delete** /identity/oidc/provider/{name} | 
[**OIDCDeleteRole**](Identity.md#OIDCDeleteRole) | **Delete** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**OIDCDeleteScope**](Identity.md#OIDCDeleteScope) | **Delete** /identity/oidc/scope/{name} | 
[**OIDCIntrospect**](Identity.md#OIDCIntrospect) | **Post** /identity/oidc/introspect | Verify the authenticity of an OIDC token
[**OIDCListAssignments**](Identity.md#OIDCListAssignments) | **Get** /identity/oidc/assignment | 
[**OIDCListClients**](Identity.md#OIDCListClients) | **Get** /identity/oidc/client | 
[**OIDCListKeys**](Identity.md#OIDCListKeys) | **Get** /identity/oidc/key | List OIDC keys
[**OIDCListProviders**](Identity.md#OIDCListProviders) | **Get** /identity/oidc/provider | 
[**OIDCListRoles**](Identity.md#OIDCListRoles) | **Get** /identity/oidc/role | List configured OIDC roles
[**OIDCListScopes**](Identity.md#OIDCListScopes) | **Get** /identity/oidc/scope | 
[**OIDCReadAssignment**](Identity.md#OIDCReadAssignment) | **Get** /identity/oidc/assignment/{name} | 
[**OIDCReadClient**](Identity.md#OIDCReadClient) | **Get** /identity/oidc/client/{name} | 
[**OIDCReadConfig**](Identity.md#OIDCReadConfig) | **Get** /identity/oidc/config | OIDC configuration
[**OIDCReadKey**](Identity.md#OIDCReadKey) | **Get** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**OIDCReadProvider**](Identity.md#OIDCReadProvider) | **Get** /identity/oidc/provider/{name} | 
[**OIDCReadProviderAuthorize**](Identity.md#OIDCReadProviderAuthorize) | **Get** /identity/oidc/provider/{name}/authorize | 
[**OIDCReadProviderUserInfo**](Identity.md#OIDCReadProviderUserInfo) | **Get** /identity/oidc/provider/{name}/userinfo | 
[**OIDCReadProviderWellKnownKeys**](Identity.md#OIDCReadProviderWellKnownKeys) | **Get** /identity/oidc/provider/{name}/.well-known/keys | 
[**OIDCReadProviderWellKnownOpenIDConfiguration**](Identity.md#OIDCReadProviderWellKnownOpenIDConfiguration) | **Get** /identity/oidc/provider/{name}/.well-known/openid-configuration | 
[**OIDCReadRole**](Identity.md#OIDCReadRole) | **Get** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**OIDCReadScope**](Identity.md#OIDCReadScope) | **Get** /identity/oidc/scope/{name} | 
[**OIDCReadToken**](Identity.md#OIDCReadToken) | **Get** /identity/oidc/token/{name} | Generate an OIDC token
[**OIDCReadWellKnownKeys**](Identity.md#OIDCReadWellKnownKeys) | **Get** /identity/oidc/.well-known/keys | Retrieve public keys
[**OIDCReadWellKnownOpenIDConfiguration**](Identity.md#OIDCReadWellKnownOpenIDConfiguration) | **Get** /identity/oidc/.well-known/openid-configuration | Query OIDC configurations
[**OIDCRotateKey**](Identity.md#OIDCRotateKey) | **Post** /identity/oidc/key/{name}/rotate | Rotate a named OIDC key.
[**OIDCWriteAssignment**](Identity.md#OIDCWriteAssignment) | **Post** /identity/oidc/assignment/{name} | 
[**OIDCWriteClient**](Identity.md#OIDCWriteClient) | **Post** /identity/oidc/client/{name} | 
[**OIDCWriteConfig**](Identity.md#OIDCWriteConfig) | **Post** /identity/oidc/config | OIDC configuration
[**OIDCWriteKey**](Identity.md#OIDCWriteKey) | **Post** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**OIDCWriteProvider**](Identity.md#OIDCWriteProvider) | **Post** /identity/oidc/provider/{name} | 
[**OIDCWriteProviderAuthorize**](Identity.md#OIDCWriteProviderAuthorize) | **Post** /identity/oidc/provider/{name}/authorize | 
[**OIDCWriteProviderToken**](Identity.md#OIDCWriteProviderToken) | **Post** /identity/oidc/provider/{name}/token | 
[**OIDCWriteProviderUserInfo**](Identity.md#OIDCWriteProviderUserInfo) | **Post** /identity/oidc/provider/{name}/userinfo | 
[**OIDCWriteRole**](Identity.md#OIDCWriteRole) | **Post** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**OIDCWriteScope**](Identity.md#OIDCWriteScope) | **Post** /identity/oidc/scope/{name} | 
[**PersonaIDDeleteByID**](Identity.md#PersonaIDDeleteByID) | **Delete** /identity/persona/id/{id} | Update, read or delete an alias ID.
[**PersonaIDReadByID**](Identity.md#PersonaIDReadByID) | **Get** /identity/persona/id/{id} | Update, read or delete an alias ID.
[**PersonaIDWriteByID**](Identity.md#PersonaIDWriteByID) | **Post** /identity/persona/id/{id} | Update, read or delete an alias ID.
[**PersonaListByID**](Identity.md#PersonaListByID) | **Get** /identity/persona/id | List all the alias IDs.
[**PersonaWrite**](Identity.md#PersonaWrite) | **Post** /identity/persona | Create a new alias.



## AliasDeleteByID

> AliasDeleteByID(ctx, id).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> AliasListByID(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> AliasReadByID(ctx, id).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> AliasWrite(ctx).AliasWriteRequest(aliasWriteRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> AliasWriteByID(ctx, id).AliasWriteByIDRequest(aliasWriteByIDRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityBatchDelete(ctx).EntityBatchDeleteRequest(entityBatchDeleteRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityDeleteAliasByID(ctx, id).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityDeleteByID(ctx, id).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityDeleteByName(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityListAliasesByID(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityListByID(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityListByName(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityLookup(ctx).EntityLookupRequest(entityLookupRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityMerge(ctx).EntityMergeRequest(entityMergeRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityReadAliasByID(ctx, id).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityReadByID(ctx, id).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityReadByName(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityWrite(ctx).EntityWriteRequest(entityWriteRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityWriteAlias(ctx).EntityWriteAliasRequest(entityWriteAliasRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityWriteAliasByID(ctx, id).EntityWriteAliasByIDRequest(entityWriteAliasByIDRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityWriteByID(ctx, id).EntityWriteByIDRequest(entityWriteByIDRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> EntityWriteByName(ctx, name).EntityWriteByNameRequest(entityWriteByNameRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupDeleteAliasByID(ctx, id).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupDeleteByID(ctx, id).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupDeleteByName(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupListAliasesByID(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupListByID(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupListByName(ctx).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupLookup(ctx).GroupLookupRequest(groupLookupRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupReadAliasByID(ctx, id).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupReadByID(ctx, id).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupReadByName(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupWrite(ctx).GroupWriteRequest(groupWriteRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupWriteAlias(ctx).GroupWriteAliasRequest(groupWriteAliasRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupWriteAliasByID(ctx, id).GroupWriteAliasByIDRequest(groupWriteAliasByIDRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupWriteByID(ctx, id).GroupWriteByIDRequest(groupWriteByIDRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> GroupWriteByName(ctx, name).GroupWriteByNameRequest(groupWriteByNameRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFADeleteLoginEnforcement(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAListLoginEnforcements(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodAdminDestroyTOTP(ctx).MFAMethodAdminDestroyTOTPRequest(mFAMethodAdminDestroyTOTPRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodAdminGenerateTOTP(ctx).MFAMethodAdminGenerateTOTPRequest(mFAMethodAdminGenerateTOTPRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodDeleteDuo(ctx, methodId).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodDeleteOkta(ctx, methodId).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodDeletePingID(ctx, methodId).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodDeleteTOTP(ctx, methodId).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodGenerateTOTP(ctx).MFAMethodGenerateTOTPRequest(mFAMethodGenerateTOTPRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodList(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodListDuo(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodListOkta(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodListPingID(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodListTOTP(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodRead(ctx, methodId).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodReadDuo(ctx, methodId).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodReadOkta(ctx, methodId).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodReadPingID(ctx, methodId).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodReadTOTP(ctx, methodId).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodWriteDuo(ctx, methodId).MFAMethodWriteDuoRequest(mFAMethodWriteDuoRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodWriteOkta(ctx, methodId).MFAMethodWriteOktaRequest(mFAMethodWriteOktaRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodWritePingID(ctx, methodId).MFAMethodWritePingIDRequest(mFAMethodWritePingIDRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAMethodWriteTOTP(ctx, methodId).MFAMethodWriteTOTPRequest(mFAMethodWriteTOTPRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAReadLoginEnforcement(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> MFAWriteLoginEnforcement(ctx, name).MFAWriteLoginEnforcementRequest(mFAWriteLoginEnforcementRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCDeleteAssignment(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCDeleteClient(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCDeleteKey(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCDeleteProvider(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCDeleteRole(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCDeleteScope(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCIntrospect(ctx).OIDCIntrospectRequest(oIDCIntrospectRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCListAssignments(ctx).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCListClients(ctx).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCListKeys(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCListProviders(ctx).List(list).AllowedClientId(allowedClientId).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCListRoles(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCListScopes(ctx).List(list).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadAssignment(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadClient(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadConfig(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadKey(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadProvider(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadProviderAuthorize(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadProviderUserInfo(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadProviderWellKnownKeys(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadProviderWellKnownOpenIDConfiguration(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadRole(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadScope(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadToken(ctx, name).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadWellKnownKeys(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCReadWellKnownOpenIDConfiguration(ctx).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCRotateKey(ctx, name).OIDCRotateKeyRequest(oIDCRotateKeyRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCWriteAssignment(ctx, name).OIDCWriteAssignmentRequest(oIDCWriteAssignmentRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCWriteClient(ctx, name).OIDCWriteClientRequest(oIDCWriteClientRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCWriteConfig(ctx).OIDCWriteConfigRequest(oIDCWriteConfigRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCWriteKey(ctx, name).OIDCWriteKeyRequest(oIDCWriteKeyRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCWriteProvider(ctx, name).OIDCWriteProviderRequest(oIDCWriteProviderRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCWriteProviderAuthorize(ctx, name).OIDCWriteProviderAuthorizeRequest(oIDCWriteProviderAuthorizeRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCWriteProviderToken(ctx, name).OIDCWriteProviderTokenRequest(oIDCWriteProviderTokenRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCWriteProviderUserInfo(ctx, name).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCWriteRole(ctx, name).OIDCWriteRoleRequest(oIDCWriteRoleRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> OIDCWriteScope(ctx, name).OIDCWriteScopeRequest(oIDCWriteScopeRequest).Execute()



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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> PersonaIDDeleteByID(ctx, id).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> PersonaIDReadByID(ctx, id).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> PersonaIDWriteByID(ctx, id).PersonaIDWriteByIDRequest(personaIDWriteByIDRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> PersonaListByID(ctx).List(list).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

> PersonaWrite(ctx).PersonaWriteRequest(personaWriteRequest).Execute()

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
		vault.WithBaseAddress("http://127.0.0.1:8200"),
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

