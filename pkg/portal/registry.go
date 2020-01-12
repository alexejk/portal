package portal

import (
	"fmt"

	"alexejk.io/portal/pkg/config"
)

type Registry struct {
	portals []Portal
}

func NewRegistry() *Registry {

	portals, _ := readPortalsFromConfig()

	r := &Registry{
		portals: portals,
	}

	return r
}

func (r *Registry) GetPortal(name string) (Portal, error) {

	for _, p := range r.portals {
		if p.Name() == name {
			return p, nil
		}
	}

	return nil, fmt.Errorf("unknown destination '%s'", name)
}

func (r *Registry) ListPortalNames() []string {

	var names []string
	for _, p := range r.portals {
		names = append(names, p.Name())
	}

	return names
}

func (r *Registry) ListPortals() []Portal {

	return r.portals
}

func readPortalsFromConfig() ([]Portal, error) {

	c, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	var portals []Portal

	if c.Portals != nil {
		for _, p := range c.Portals {

			var port Portal

			switch {
			case p.Aws != nil:
				port, err = newAwsPortal(*p.Name, p.Aws)

			case p.Raw != nil:
				port, err = newRawPortal(*p.Name, p.Raw)

			default:
				return nil, fmt.Errorf("empty portal definition: %s", *p.Name)
			}

			if err != nil {
				return nil, err
			}

			if p.Hint != nil {
				port.SetHint(*p.Hint)
			}

			portals = append(portals, port)
		}
	}

	return portals, nil
}
