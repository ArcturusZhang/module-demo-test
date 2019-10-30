package bznet

type BzNetClient struct {

}

type Bznet struct {
	Nmae string
	Age int
}

func (client BzNetClient) Create(net Bznet) error {
	return nil
}

func (client BzNetClient) Get() Bznet {
	return Bznet{
		Nmae: "Module-BzNet",
		Age:  10,
	}
}
