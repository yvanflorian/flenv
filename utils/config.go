package utils

const (
	FLENV_CONFIG_FILENAME string = ".flenv.json"
	FLENV_CONFIG_PREFIX   string = "FLENV"
)

func NewConfig(stageName string) Flenv {
	return Flenv{
		Stages: []Stage{
			{StageName: stageName},
		},
	}
}
