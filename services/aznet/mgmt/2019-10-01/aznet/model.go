package aznet

type AzNetsClient struct {

}

// AzNet ...
type AzNet struct {
	// Name ...
	Name string
	// ID ...
	ID int
}

func (client AzNetsClient) Create(net AzNet) error {
	return nil
}

func (client AzNetsClient) Get() AzNet {
	return AzNet{
		Name: "Module-AzNet",
		ID:   20,
	}
}
