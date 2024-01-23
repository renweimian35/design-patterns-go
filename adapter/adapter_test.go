package adapter

import "testing"

func TestAdapter(t *testing.T) {
	cigaretteLighter := &CigaretteLighter{
		ConnectedDevice: "USB Charger",
	}
	adapter := &Adapter{
		CigaretteLighter: cigaretteLighter,
	}
	adapter.Charge()
}
