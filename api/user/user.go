
// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package user

import (


	authV1 "github.com/sucold/starter/api/user/v1/auth"
)

type Interface interface { 
	V1() V1Interface
}


type V1Interface interface { 
	Auth() authV1.IAuth
}
