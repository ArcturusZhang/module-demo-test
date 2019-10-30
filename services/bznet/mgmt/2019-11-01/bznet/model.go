package bznet

type BzNetsClient struct {

}

type Bznet struct {
	Nmae string
	Age int
}

func (client BzNetsClient) Create(net Bznet) error {
	return nil
}

func (client BzNetsClient) Get() Bznet {
	return Bznet{
		Nmae: "Module-BzNet",
		Age:  10,
	}
}
