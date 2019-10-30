package aznetapi

import "github.com/ArcturusZhang/module-demo-test/services/aznet/mgmt/2019-10-01/aznet"

type AzNetApi interface {
	Create(net aznet.AzNet) error
	Get() aznet.AzNet
}
