package airbyte

import "fmt"

type AirbyteClient struct {
}

func (a *AirbyteClient) GetSources() {
	fmt.Println("not implemented")
}

func (a *AirbyteClient) CreateSource() {
	fmt.Println("not implemented")

}

func (a *AirbyteClient) DeleteSource() {
	fmt.Println("not implemented")
}

func (a *AirbyteClient) UpdateSource() {
	fmt.Println("not implemented")
}

func NewClient() *AirbyteClient {
	return &AirbyteClient{}
}
