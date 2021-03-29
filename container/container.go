package container

type Builder func(con *Container) interface{}

type Definition struct {
	Name    string
	Builder Builder
}

type Container struct {
	store map[string]Builder
}

func NewContainer() *Container {
	return &Container{
		store: map[string]Builder{},
	}
}

func (c *Container) Register(d *Definition) {
	c.store[d.Name] = d.Builder
}

func (c *Container) Get(key string) interface{} {
	builder := c.store[key]
	instance := builder(c)
	return instance
}
