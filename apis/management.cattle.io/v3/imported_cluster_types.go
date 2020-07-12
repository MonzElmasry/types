package v3

//K3sConfig provides desired configuration for k3s clusters
type ClusterConfig struct {
	// k3s Kubernetes version, unset the value indicates an unmanaged cluster
	Version                string `yaml:"kubernetes_version" json:"kubernetesVersion,omitempty"`
	ClusterUpgradeStrategy `yaml:"cluster_upgrade_strategy,omitempty" json:"clusterupgradeStrategy,omitempty"`
}

//K3sUpgradeStrategy provides configuration to the downstream system-upgrade-controller
type ClusterUpgradeStrategy struct {
	// How many controlplane nodes should be upgrade at time, defaults to 1
	ServerConcurrency int `yaml:"server_concurrency" json:"serverConcurrency,omitempty" norman:"min=1"`
	// How many workers should be upgraded at a time
	WorkerConcurrency int `yaml:"worker_concurrency" json:"workerConcurrency,omitempty" norman:"min=1"`
	// Whether controlplane nodes should be drained
	DrainServerNodes bool `yaml:"drain_server_nodes" json:"drainServerNodes,omitempty"`
	// Whether worker nodes should be drained
	DrainWorkerNodes bool `yaml:"drain_worker_nodes" json:"drainWorkerNodes,omitempty"`
}
