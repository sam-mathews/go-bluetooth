// Code generated DO NOT EDIT

package mesh



import (
   "sync"
   "github.com/muka/go-bluetooth/bluez"
   "github.com/muka/go-bluetooth/util"
   "github.com/muka/go-bluetooth/props"
   "github.com/godbus/dbus/v5"
)

var Provisioner1Interface = "org.bluez.mesh.Provisioner1"


// NewProvisioner1 create a new instance of Provisioner1
//
// Args:
// - servicePath: unique name
// - objectPath: freely definable
func NewProvisioner1(servicePath string, objectPath dbus.ObjectPath) (*Provisioner1, error) {
	a := new(Provisioner1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  servicePath,
			Iface: Provisioner1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(Provisioner1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
Provisioner1 Mesh Provisioner Hierarchy

*/
type Provisioner1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*Provisioner1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// Provisioner1Properties contains the exposed properties of an interface
type Provisioner1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

}

//Lock access to properties
func (p *Provisioner1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *Provisioner1Properties) Unlock() {
	p.lock.Unlock()
}



// Close the connection
func (a *Provisioner1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return Provisioner1 object path
func (a *Provisioner1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return Provisioner1 dbus client
func (a *Provisioner1) Client() *bluez.Client {
	return a.client
}

// Interface return Provisioner1 interface
func (a *Provisioner1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *Provisioner1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}


// ToMap convert a Provisioner1Properties to map
func (a *Provisioner1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an Provisioner1Properties
func (a *Provisioner1Properties) FromMap(props map[string]interface{}) (*Provisioner1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an Provisioner1Properties
func (a *Provisioner1Properties) FromDBusMap(props map[string]dbus.Variant) (*Provisioner1Properties, error) {
	s := new(Provisioner1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *Provisioner1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *Provisioner1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *Provisioner1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *Provisioner1) GetProperties() (*Provisioner1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *Provisioner1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *Provisioner1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *Provisioner1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *Provisioner1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *Provisioner1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *Provisioner1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




/*
ScanResult 
		The method is called from the bluetooth-meshd daemon when a
		unique UUID has been seen during UnprovisionedScan() for
		unprovsioned devices.

		The rssi parameter is a signed, normalized measurement of the
		signal strength of the recieved unprovisioned beacon.

		The data parameter is a variable length byte array, that may
		have 1, 2 or 3 distinct fields contained in it including the 16
		byte remote device UUID (always), a 32 bit mask of OOB
		authentication flags (optional), and a 32 bit URI hash (if URI
		bit set in OOB mask). Whether these fields exist or not is a
		decision of the remote device.

		If a beacon with a UUID that has already been reported is
		recieved by the daemon, it will be silently discarded unless it
		was recieved at a higher rssi power level.



*/
func (a *Provisioner1) ScanResult(rssi int16, data []byte) error {
	
	return a.client.Call("ScanResult", 0, rssi, data).Store()
	
}

/*
RequestProvData 
		This method is implemented by a Provisioner capable application
		and is called when the remote device has been fully
		authenticated and confirmed.

		The count parameter is the number of consecutive unicast
		addresses the remote device is requesting.

		Return Parameters are from the Mesh Profile Spec:
		net_index - Subnet index of the net_key
		unicast - Primary Unicast address of the new node

		PossibleErrors:
			org.bluez.mesh.Error.Abort


*/
func (a *Provisioner1) RequestProvData(count uint8) (uint16 net_index, uint16 unicast, error) {
	
	var val0 uint16 net_index
  var val1 uint16 unicast
	err := a.client.Call("RequestProvData", 0, count).Store(&val0, &val1)
	return val0, val1, err	
}

/*
AddNodeComplete 
		This method is called when the node provisioning initiated
		by an AddNode() method call successfully completed.

		The unicast parameter is the primary address that has been
		assigned to the new node, and the address of it's config server.

		The count parameter is the number of unicast addresses assigned
		to the new node.

		The new node may now be sent messages using the credentials
		supplied by the RequestProvData method.


*/
func (a *Provisioner1) AddNodeComplete(uuid []byte, unicast uint16, count uint8) error {
	
	return a.client.Call("AddNodeComplete", 0, uuid, unicast, count).Store()
	
}

/*
AddNodeFailed 
		This method is called when the node provisioning initiated by
		AddNode() has failed. Depending on how far Provisioning
		proceeded before failing, some cleanup of cached data may be
		required.

		The reason parameter identifies the reason for provisioning
		failure. The defined values are: "aborted", "timeout",
		"bad-pdu", "confirmation-failed", "out-of-resources",
		"decryption-error", "unexpected-error",
		"cannot-assign-addresses".


*/
func (a *Provisioner1) AddNodeFailed(uuid []byte, reason string) error {
	
	return a.client.Call("AddNodeFailed", 0, uuid, reason).Store()
	
}

