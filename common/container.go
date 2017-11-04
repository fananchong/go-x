package common

type IComponent interface {
	Init(params []interface{})
	Update()
}

type Container struct {
	components []IComponent
}

func (this *Container) Add(component IComponent) {
	this.components = append(this.components, component)
}

func (this *Container) Init(params []interface{}) {
	for _, component := range this.components {
		component.Init(params)
	}
}

func (this *Container) Update() {
	for _, component := range this.components {
		component.Update()
	}
}
