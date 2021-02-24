package env

type EnvType struct {
}


//本地环境
func (env EnvType) Local() string {
	return "local"
}

//测试环境
func (env EnvType) Dev() string {
	return "dev"
}

//生产环境
func (env EnvType) Product() string {
	return "product"
}
