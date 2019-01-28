package static

type Static struct {
	files map[string]*File
}

type File struct {
	Cache []byte
}

func New() *Static {
	return &Static{
		map[string]*File{},
	}
}
