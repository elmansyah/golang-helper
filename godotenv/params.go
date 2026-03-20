package godotenv

type Params struct{}

type Godotenv interface {
	Load() string
}
