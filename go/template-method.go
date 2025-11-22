package main

type Component struct{}

func (b *Component) render() {
	b.setup()

	b.beforeCreate()
	b.initOptionsApi()
	b.created()

	b.breforeMount()
	b.createAndInsertDomNodes()
	b.mounted()

	b.beforeUpdate()
	b.updated()

	b.beforeUnmount()
	b.unmounted()
}

func (b *Component) initOptionsApi()          {}
func (b *Component) createAndInsertDomNodes() {}

func (b *Component) setup()         {}
func (b *Component) beforeCreate()  {}
func (b *Component) created()       {}
func (b *Component) breforeMount()  {}
func (b *Component) mounted()       {}
func (b *Component) beforeUpdate()  {}
func (b *Component) updated()       {}
func (b *Component) beforeUnmount() {}
func (b *Component) unmounted()     {}

type Avatar struct {
	Component
}

func (a *Avatar) mounted()       {}
func (a *Avatar) beforeUnmount() {}

func init() {
	avatar := Avatar{
		Component{},
	}

	avatar.render()
}
