package client

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindInterfaceWireguardPeer_onNonExistantInterfacePeer(t *testing.T) {
	SkipIfRouterOSV6OrEarlier(t, sysResources)
	c := NewClient(GetConfigFromEnv())

	interfaceName := "Interface peer does not exist"
	_, err := c.FindInterfaceWireguardPeer(interfaceName)

	require.Truef(t, IsNotFoundError(err),
		"Expecting to receive NotFound error for Interface peer `%q`, instead error was nil.", interfaceName)
}

func TestInterfaceWireguardPeer_Crud(t *testing.T) {
	SkipIfRouterOSV6OrEarlier(t, sysResources)
	c := NewClient(GetConfigFromEnv())

	name := "new_interface_wireguard"
	interfaceWireguard := &InterfaceWireguard{
		Name:       name,
		Disabled:   false,
		ListenPort: 10000,
		Mtu:        10001,
		PrivateKey: "YOi0P0lTTiN8hAQvuRET23Srb+U7C52iOZokj0CCSkM=",
		Comment:    "new interface from test",
	}

	createdInterface, err := c.Add(interfaceWireguard)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}
	defer func() {
		err = c.Delete(interfaceWireguard)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	}()

	interfaceWireguardPeer := &InterfaceWireguardPeer{
		Interface: createdInterface.(*InterfaceWireguard).Name,
		Disabled:  false,
		Comment:   "new interface from test",
	}

	created, err := c.Add(interfaceWireguardPeer)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}
	defer func() {
		err = c.Delete(interfaceWireguardPeer)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	}()
	findInterface := &InterfaceWireguardPeer{}
	findInterface.Interface = createdInterface.(*InterfaceWireguard).Name
	found, err := c.Find(findInterface)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	if _, ok := found.(Resource); !ok {
		t.Error("expected found resource to implement Resource interface, but it doesn't")
		return
	}
	if !reflect.DeepEqual(created, found) {
		t.Error("expected created and found resources to be equal, but they aren't")
	}
}
