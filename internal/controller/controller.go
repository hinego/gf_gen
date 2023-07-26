package controller

import (
	userFace "github.com/sucold/starter/api/user"
	userAuthV1Api "github.com/sucold/starter/api/user/v1/auth"
	dataFace "github.com/sucold/starter/api/data"
	dataAuthV1Api "github.com/sucold/starter/api/data/v1/auth"
	dataAuthV2Api "github.com/sucold/starter/api/data/v2/auth"
	dataUserV2Api "github.com/sucold/starter/api/data/v2/user"
)

var (
	_userAuthV1 userAuthV1Api.IAuth
	_dataAuthV1 dataAuthV1Api.IAuth
	_dataAuthV2 dataAuthV2Api.IAuth
	_dataUserV2 dataUserV2Api.IUser
)

func Register(name string, data any) {
	var ok bool
	switch name {
	case "userAuthV1":
		if _userAuthV1,ok = data.(userAuthV1Api.IAuth);!ok {
			panic("userAuthV1 register error")
		}
		break
	case "dataAuthV1":
		if _dataAuthV1,ok = data.(dataAuthV1Api.IAuth);!ok {
			panic("dataAuthV1 register error")
		}
		break
	case "dataAuthV2":
		if _dataAuthV2,ok = data.(dataAuthV2Api.IAuth);!ok {
			panic("dataAuthV2 register error")
		}
		break
	case "dataUserV2":
		if _dataUserV2,ok = data.(dataUserV2Api.IUser);!ok {
			panic("dataUserV2 register error")
		}
		break
	}
}
type user struct{}
func User() userFace.Interface {
	return &user{}
}
type data struct{}
func Data() dataFace.Interface {
	return &data{}
}
func (r *user) V1() userFace.V1Interface {
	return &userV1{}
}
func (r *data) V1() dataFace.V1Interface {
	return &dataV1{}
}
func (r *data) V2() dataFace.V2Interface {
	return &dataV2{}
}
type userV1 struct{}
func (d *userV1) Auth() userAuthV1Api.IAuth {
	return _userAuthV1
}
type dataV1 struct{}
func (d *dataV1) Auth() dataAuthV1Api.IAuth {
	return _dataAuthV1
}
type dataV2 struct{}
func (d *dataV2) Auth() dataAuthV2Api.IAuth {
	return _dataAuthV2
}
func (d *dataV2) User() dataUserV2Api.IUser {
	return _dataUserV2
}
