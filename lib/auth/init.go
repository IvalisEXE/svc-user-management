package auth

type Opts struct {
	TTL          int
	SignatureKey string
}

type Module struct {
	ttl          int
	signatureKey string
}

func New(o *Opts) *Module {
	return &Module{
		ttl:          o.TTL,
		signatureKey: o.SignatureKey,
	}
}
