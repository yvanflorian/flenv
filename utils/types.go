package utils

// main structs to use

type Flenv struct {
	Stages []Stage `json:"flenv"`
}

type Stage struct {
	StageName string   `json:"stage"`
	Configs   []Config `json:"configs"`
}

type Config struct {
	Name      string     `json:"name"`
	Variables []Variable `json:"variables"`
}

type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
