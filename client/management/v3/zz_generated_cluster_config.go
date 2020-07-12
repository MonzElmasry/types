package client

const (
	ClusterConfigType                        = "clusterConfig"
	ClusterConfigFieldClusterUpgradeStrategy = "clusterupgradeStrategy"
	ClusterConfigFieldVersion                = "kubernetesVersion"
)

type ClusterConfig struct {
	ClusterUpgradeStrategy *ClusterUpgradeStrategy `json:"clusterupgradeStrategy,omitempty" yaml:"clusterupgradeStrategy,omitempty"`
	Version                string                  `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
}
