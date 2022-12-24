package oap

import "fmt"

const PermissionPathDelim = "."

var _ fmt.Stringer = Permission("")

type Permission string

func NewPermission(permission string) Permission { return Permission(permission) }

func (p Permission) String() string                        { return string(p) }
func (p Permission) AppendPath(path Permission) Permission { return p + PermissionPathDelim + path }
