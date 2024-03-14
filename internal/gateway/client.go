package gateway

type ClientGateway interface {
	Get(id string) (*Client, error)
	Save(client *Client) error
}
