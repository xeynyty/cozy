package cozy

const defaultPath = "config.cozy"

func Init(path *string) error {
	core := GetInstance()

	var _path string
	if path == nil {
		_path = defaultPath
	} else {
		_path = *path
	}

	err := core.Init(_path)
	return err
}

func Get(key string, out *any) error {
	core := GetInstance()
	return core.Get(key, out)
}

func Print() {
	core := GetInstance()
	core.Print()
}
