package limiter

type Limiter struct {
	mask string
}

func NewLimiter(mask string) *Limiter {
	return &Limiter{
		mask: mask,
	}
}
