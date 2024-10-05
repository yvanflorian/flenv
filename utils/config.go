package utils

const (
	FLENV_CONFIG_FILENAME     string = ".flenv.json"
	FLENV_KEYRING_SERVICENAME string = "GOFLENVCLI"
	FLENV_KEYRING_SERVICEUSER string = "GOFLENVCLICLIENT"
)

func NewConfig(stageName string) Flenv {
	return Flenv{
		Stages: []Stage{
			{StageName: stageName},
		},
	}
}
