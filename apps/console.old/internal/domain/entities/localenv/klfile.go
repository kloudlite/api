package localenv

type ResEnv struct {
	Name   string
	Key    string
	RefKey string
}

type Env struct {
	Key   string
	Value string
}

type Res struct {
	Name string
	Env  []ResEnv
}

type FileEntry struct {
	Path string
	Type string
	Name string
}

type Mount struct {
	MountBasePath string `yaml:"mountBasePath"`
	Mounts        []FileEntry
}

type KLFile struct {
	Mres      []Res
	Configs   []Res
	Secrets   []Res
	Env       []Env
	FileMount Mount
}