package goacoap

const PermissionPathDelim = "."

type Permission string

func NewPermission(permission string) Permission {
	return Permission(permission)
}

func (p Permission) String() string {
	return string(p)
}

func (p Permission) AppendPath(path string) Permission {
	return p + Permission(PermissionPathDelim+path)
}
