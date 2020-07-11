package v3

type Rke2Config struct {
	// rke2 Kubernetes version, unset the value indicates an unmanaged cluster
	Version             string `yaml:"kubernetes_version" json:"kubernetesVersion,omitempty"`
	Rke2UpgradeStrategy `yaml:"rke2_upgrade_strategy,omitempty" json:"rke2upgradeStrategy,omitempty"`
}

//rke2UpgradeStrategy provides configuration to the downstream system-upgrade-controller
type Rke2UpgradeStrategy struct {
	// How many controlplane nodes should be upgrade at time, defaults to 1
	ServerConcurrency int `yaml:"server_concurrency" json:"serverConcurrency,omitempty" norman:"min=1"`
	// How many workers should be upgraded at a time
	WorkerConcurrency int `yaml:"worker_concurrency" json:"workerConcurrency,omitempty" norman:"min=1"`
	// Whether controlplane nodes should be drained
	DrainServerNodes bool `yaml:"drain_server_nodes" json:"drainServerNodes,omitempty"`
	// Whether worker nodes should be drained
	DrainWorkerNodes bool `yaml:"drain_worker_nodes" json:"drainWorkerNodes,omitempty"`
}
