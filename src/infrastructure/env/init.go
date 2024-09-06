package env

type EnvType string

const (
	ENV_TYPE_DEVELOPMENT EnvType = "development"
	ENV_TYPE_TEST        EnvType = "test"
	ENV_TYPE_PRODUCTION  EnvType = "production"
)

var (
	_env EnvType = ENV_TYPE_DEVELOPMENT
)

func GetEnv() EnvType {
	return _env
}

func SetEnv(env EnvType) {
	_env = env
}

func IsEnvDevelopment() bool {
	return _env == ENV_TYPE_DEVELOPMENT
}

func IsEnvTest() bool {
	return _env == ENV_TYPE_TEST
}

func IsEnvProduction() bool {
	return _env == ENV_TYPE_PRODUCTION
}
