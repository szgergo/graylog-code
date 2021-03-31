package api

type InputConfigurationApi struct {
	RecvBufferSize int `json:"recv_buffer_size"`
	Port int `json:"port"`
	NumberWorkerThreads int `json:"number_worker_threads"`
	OverrideSource string `json:"override_source,omitempty"`
	BindAddress string `json:"bind_address"`
	DecompressSizeLimit int `json:"decompress_size_limit"`
}

type InputApi struct {
	Title string                         `json:"title"`
	Global bool                          `json:"global"`
	Type string                          `json:"type"`
	Configuration *InputConfigurationApi `json:"configuration"`
	Node string                          `json:"node"`
}
