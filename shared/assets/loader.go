package assets

var LoaderInstance AssetLoader

type AssetLoader interface {
	Load(path string) ([]byte, error)
}

func GetLoader() AssetLoader {
	return LoaderInstance
}
