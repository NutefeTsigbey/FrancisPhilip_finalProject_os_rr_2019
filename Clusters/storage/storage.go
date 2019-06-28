
package storage

import (
	"errors"
)

var (
	ErrNoSuchNode            = errors.New("No such node in storage")
	ErrNoSuchContainer       = errors.New("No such container in storage")
	ErrNoSuchImage           = errors.New("No such image in storage")
	ErrNoSuchExec            = errors.New("No such exec in storage")
	ErrDuplicatedNodeAddress = errors.New("Node address shouldn't repeat")
)
