package utils

const FLENV_CONFIG_FILENAME string = ".flenv.json"

func NewConfig(stageName string) Flenv {
	return Flenv{
		Stages: []Stage{
			{StageName: stageName},
		},
	}
}
