// WARNING: generated code, do not edit!
// Copyright © 2019 luca capra
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package thermometer



import (
  "sync"
  "github.com/muka/go-bluetooth/bluez"
  "reflect"
  "github.com/fatih/structs"
  "github.com/muka/go-bluetooth/util"
  "github.com/godbus/dbus"
)

var ThermometerManager1Interface = "org.bluez.ThermometerManager1"


// NewThermometerManager1 create a new instance of ThermometerManager1
//
// Args:
// 	objectPath: [variable prefix]/{hci0,hci1,...}
func NewThermometerManager1(objectPath dbus.ObjectPath) (*ThermometerManager1, error) {
	a := new(ThermometerManager1)
	a.propertiesSignal = make(chan *dbus.Signal)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: ThermometerManager1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(ThermometerManager1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


// ThermometerManager1 Health Thermometer Manager hierarchy

type ThermometerManager1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	Properties 				*ThermometerManager1Properties
}

// ThermometerManager1Properties contains the exposed properties of an interface
type ThermometerManager1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

}

func (p *ThermometerManager1Properties) Lock() {
	p.lock.Lock()
}

func (p *ThermometerManager1Properties) Unlock() {
	p.lock.Unlock()
}

// Close the connection
func (a *ThermometerManager1) Close() {
	
	a.unregisterSignal()
	
	a.client.Disconnect()
}

// Path return ThermometerManager1 object path
func (a *ThermometerManager1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Interface return ThermometerManager1 interface
func (a *ThermometerManager1) Interface() string {
	return a.client.Config.Iface
}


// ToMap convert a ThermometerManager1Properties to map
func (a *ThermometerManager1Properties) ToMap() (map[string]interface{}, error) {
	return structs.Map(a), nil
}

// FromMap convert a map to an ThermometerManager1Properties
func (a *ThermometerManager1Properties) FromMap(props map[string]interface{}) (*ThermometerManager1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an ThermometerManager1Properties
func (a *ThermometerManager1Properties) FromDBusMap(props map[string]dbus.Variant) (*ThermometerManager1Properties, error) {
	s := new(ThermometerManager1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// GetProperties load all available properties
func (a *ThermometerManager1) GetProperties() (*ThermometerManager1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *ThermometerManager1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *ThermometerManager1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *ThermometerManager1) GetPropertiesSignal() (chan *dbus.Signal, error) {

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
func (a *ThermometerManager1) unregisterSignal() {
	if a.propertiesSignal == nil {
		a.propertiesSignal <- nil
	}
}

// WatchProperties updates on property changes
func (a *ThermometerManager1) WatchProperties() (chan *bluez.PropertyChanged, error) {

	channel, err := a.client.Register(a.Path(), a.Interface())
	if err != nil {
		return nil, err
	}

	ch := make(chan *bluez.PropertyChanged)

	go (func() {
		for {

			if channel == nil {
				break
			}

			sig := <-channel

			if sig == nil {
				return
			}

			if sig.Name != bluez.PropertiesChanged {
				continue
			}
			if sig.Path != a.Path() {
				continue
			}

			iface := sig.Body[0].(string)
			changes := sig.Body[1].(map[string]dbus.Variant)

			for field, val := range changes {

				// updates [*]Properties struct
				props := a.Properties

				s := reflect.ValueOf(props).Elem()
				// exported field
				f := s.FieldByName(field)
				if f.IsValid() {
					// A Value can be changed only if it is
					// addressable and was not obtained by
					// the use of unexported struct fields.
					if f.CanSet() {
						x := reflect.ValueOf(val.Value())
						props.Lock()
						f.Set(x)
						props.Unlock()
					}
				}

				propChanged := &bluez.PropertyChanged{
					Interface: iface,
					Name:      field,
					Value:     val.Value(),
				}
				ch <- propChanged
			}

		}
	})()

	return ch, nil
}

func (a *ThermometerManager1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	ch <- nil
	close(ch)
	return nil
}





//RegisterWatcher Registers a watcher to monitor scanned measurements.
// This agent will be notified about final temperature
// measurements.
func (a *ThermometerManager1) RegisterWatcher(agent dbus.ObjectPath) error {
	
	return a.client.Call("RegisterWatcher", 0, agent).Store()
	
}

//UnregisterWatcher 
func (a *ThermometerManager1) UnregisterWatcher(agent dbus.ObjectPath) error {
	
	return a.client.Call("UnregisterWatcher", 0, agent).Store()
	
}

//EnableIntermediateMeasurement Enables intermediate measurement notifications
// for this agent. Intermediate measurements will
// be enabled only for thermometers which support it.
func (a *ThermometerManager1) EnableIntermediateMeasurement(agent dbus.ObjectPath) error {
	
	return a.client.Call("EnableIntermediateMeasurement", 0, agent).Store()
	
}

//DisableIntermediateMeasurement Disables intermediate measurement notifications
// for this agent. It will disable notifications in
// thermometers when the last agent removes the
// watcher for intermediate measurements.
// Possible Errors: org.bluez.Error.InvalidArguments
// org.bluez.Error.NotFound
func (a *ThermometerManager1) DisableIntermediateMeasurement(agent dbus.ObjectPath) error {
	
	return a.client.Call("DisableIntermediateMeasurement", 0, agent).Store()
	
}
