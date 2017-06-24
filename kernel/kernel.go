package kernel

import "github.com/PumpkinSeed/refima/config"

type Kernel struct {
	Config config.Config
}

func NewKernel(c config.Config) *Kernel {
	k := new(Kernel)
	k.Config = c
	return k
}
