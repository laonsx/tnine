package opts

import "context"

type Option func(o *Options)

type Options struct {
	Ctx context.Context
}

func NewOptions(opts ...Option) Options {
	opt := Options{}
	opt.Init(opts...)

	return opt
}

func (opt *Options) Init(opts ...Option) {
	for _, o := range opts {
		o(opt)
	}

}

func WithContext(ctx context.Context) Option {
	return func(o *Options) {
		o.Ctx = ctx
	}
}
