package goacoap

const PathDelimiter = "."

type Path string

func NewPath(path string) Path {
	return Path(path)
}

func (p Path) String() string {
	return string(p)
}

func (p Path) Append(name string) Path {
	return p + PathDelimiter + Path(name)
}
