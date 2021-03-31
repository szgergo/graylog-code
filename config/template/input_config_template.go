package template

type InputConfigurationTemplate struct {
	RecvBufferSize 		int 	`yaml:"recv_buffer_size"`
	Port 				int 	`yaml:"port"`
	NumberWorkerThreads int 	`yaml:"number_worker_threads"`
	OverrideSource 		string  `yaml:"override_source,omitempty"`
	BindAddress 		string  `yaml:"bind_address"`
	DecompressSizeLimit int 	`yaml:"decompress_size_limit"`
}

func (cc *InputConfigurationTemplate) Equals(other InputConfigurationTemplate) bool {
	return cc.RecvBufferSize == other.RecvBufferSize &&
		cc.Port == other.Port &&
		cc.NumberWorkerThreads == other.NumberWorkerThreads &&
		cc.OverrideSource == other.OverrideSource &&
		cc.BindAddress == other.BindAddress &&
		cc.DecompressSizeLimit == other.DecompressSizeLimit
}

func inputConfigurationPointerEquals(a *InputConfigurationTemplate, b *InputConfigurationTemplate) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Equals(*b)
}

