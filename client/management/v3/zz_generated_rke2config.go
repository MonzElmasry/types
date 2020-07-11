package client

const (
	Rke2ConfigType                     = "rke2Config"
	Rke2ConfigFieldRke2UpgradeStrategy = "rke2upgradeStrategy"
	Rke2ConfigFieldVersion             = "kubernetesVersion"
)

type Rke2Config struct {
	Rke2UpgradeStrategy *Rke2UpgradeStrategy `json:"rke2upgradeStrategy,omitempty" yaml:"rke2upgradeStrategy,omitempty"`
	Version             string               `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
}
