package client

const (
	ClusterConfigType                        = "clusterConfig"
	ClusterConfigFieldClusterUpgradeStrategy = "k3supgradeStrategy"
	ClusterConfigFieldVersion                = "kubernetesVersion"
)

type ClusterConfig struct {
	ClusterUpgradeStrategy *ClusterUpgradeStrategy `json:"k3supgradeStrategy,omitempty" yaml:"k3supgradeStrategy,omitempty"`
	Version                string                  `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
}
