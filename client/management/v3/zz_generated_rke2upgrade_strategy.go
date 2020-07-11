package client

const (
	Rke2UpgradeStrategyType                   = "rke2UpgradeStrategy"
	Rke2UpgradeStrategyFieldDrainServerNodes  = "drainServerNodes"
	Rke2UpgradeStrategyFieldDrainWorkerNodes  = "drainWorkerNodes"
	Rke2UpgradeStrategyFieldServerConcurrency = "serverConcurrency"
	Rke2UpgradeStrategyFieldWorkerConcurrency = "workerConcurrency"
)

type Rke2UpgradeStrategy struct {
	DrainServerNodes  bool  `json:"drainServerNodes,omitempty" yaml:"drainServerNodes,omitempty"`
	DrainWorkerNodes  bool  `json:"drainWorkerNodes,omitempty" yaml:"drainWorkerNodes,omitempty"`
	ServerConcurrency int64 `json:"serverConcurrency,omitempty" yaml:"serverConcurrency,omitempty"`
	WorkerConcurrency int64 `json:"workerConcurrency,omitempty" yaml:"workerConcurrency,omitempty"`
}
