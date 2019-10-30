package bznetapi

import "github.com/ArcturusZhang/module-demo-test/services/bznet/mgmt/2019-11-01/bznet"

type BznetApi interface {
	Create(net bznet.Bznet) error
	Get() bznet.Bznet
}
