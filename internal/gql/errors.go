package gql

import "errors"

var ErrTargetIsNotPointer = errors.New("Client request target is not pointer")
var ErrResponseWithErrors = errors.New("Client response return with errors")
