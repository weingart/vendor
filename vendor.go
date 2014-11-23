package vendor

import (
	"expvar"
	"sync"
)

type Info map[string]string

type vendorLockedArray struct {
	sync.RWMutex
	info []Info
}

var vendor *vendorLockedArray = &vendorLockedArray{
	info: make([]Info, 0),
}

func GetInfo() []Info {
	defer vendor.RUnlock()
	vendor.RLock()

	info := make([]Info, len(vendor.info))
	copy(info, vendor.info)

	return info
}

func Add(v *Info) {
	defer vendor.Unlock()
	vendor.Lock()

	vendor.info = append(vendor.info, *v)
}

func init() {
	defer vendor.RUnlock()
	vendor.RLock()

	expvar.Publish("vendor", expvar.Func(func() interface{} {
		return vendor.info
	}))
}
