package portal

// Portal is a common interface for all destinations to implement
type Portal interface {
	// Name of the portal
	Name() string
	// Connect to this portal
	Connect() error

	Hint() string
	SetHint(string)
}

type portalShared struct {
	name string
	hint string
}

func (p *portalShared) Name() string {
	return p.name
}

func (p *portalShared) SetHint(h string) {
	p.hint = h
}

func (p *portalShared) Hint() string {
	return p.hint
}
