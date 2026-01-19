package cache

type Option interface {
	GetValue() any
}

type expOption struct {
	minutes int
}

func WithExpiration(minutes int) Option {
	return expOption{minutes: minutes}
}

func (opt expOption) GetValue() any {
	return opt.minutes
}
