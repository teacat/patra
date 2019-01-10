package client

// Config
type Config struct {
	// Address
	Address string
}

// QueryOptions
type QueryOptions struct {
	// Name
	Name string
	// Version
	Version string
	// Tag
	Tag string
}

// New
func New(config *Config) (*Client, error) {
	return &Client{}, nil
}

// Client
type Client struct {
	// config
	config *Config
	//

}

// Register
func (c *Client) Register(config *RegisterConfig) error {
	return nil
}

// Service
func (c *Client) Service(q *QueryOptions) (*ServiceInfo, error) {

}

// Services
func (c *Client) Services(q *QueryOptions) ([]*ServiceInfo, error) {
	return nil, nil
}

// Deregister
func (c *Client) Deregister(id string) error {
	return nil
}

// KV
func (c *Client) KV() *KV {
	return newKV(c)
}

//
type RegisterConfig struct {
	// ID
	ID string
	// Name
	Name string
	// Port
	Port int
	// Address
	Address string
	// Tags
	Tags []string
	// Metadata
	Metadata map[string]string
	// Version
	Version string
}

// ServiceInfo
type ServiceInfo struct {
	// ID
	ID string
	// Name
	Name string
	// Port
	Port int
	// Address
	Address string
	// Tags
	Tags []string
	// Metadata
	Metadata map[string]string
	// Version
	Version string
}
