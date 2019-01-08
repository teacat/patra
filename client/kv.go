package client

// newKV
func newKV(client *Client) *KV {
	return &KV{
		client: client,
	}
}

// KV
type KV struct {
	// client
	client *Client
}

// Get
func (k *KV) Get() (string, error) {

}

// Delete
func (k *KV) Delete() (bool, error) {

}

// Put
func (k *KV) Put() error {
	return nil
}

// List
func (k *KV) List() (map[string]string, error) {

}

// Keys
func (k *KV) Keys() ([]string, error) {

}
