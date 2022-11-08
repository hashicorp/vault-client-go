package vault

type ClientOption func(*Configuration) error

func New(opts ...ClientOption) (*Client, error) {

	configuration := DefaultConfiguration()
	for _, opt := range opts {
		if err := opt(&configuration); err != nil {
			return nil, err
		}
	}

	return NewClient(configuration)
}

func FromEnv(configuration *Configuration) error {
	return configuration.LoadEnvironment()
}

func WithBaseAddress(address string) ClientOption {
	return func(configuration *Configuration) error {
		configuration.BaseAddress = address
		return nil
	}
}
