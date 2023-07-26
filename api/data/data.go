
// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package data

import (


	authV1 "github.com/sucold/starter/api/data/v1/auth"

	authV2 "github.com/sucold/starter/api/data/v2/auth"
	userV2 "github.com/sucold/starter/api/data/v2/user"
)

type Interface interface { 
	V1() V1Interface
	V2() V2Interface
}


type V1Interface interface { 
	Auth() authV1.IAuth
}
type V2Interface interface { 
	Auth() authV2.IAuth
	User() userV2.IUser
}
