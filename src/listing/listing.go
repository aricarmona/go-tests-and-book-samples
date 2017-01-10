package listing

type Features struct {
	Barbecue bool
	Garden bool
	Pool bool
	// ...
}

type Address struct {
	State string
	City string
	Neighborhood string
	Street string
	Number string
	ZipCode string
}

type Property struct {
	Features Features
	Address Address
	TotalArea int
	UtilArea int
	Floor int
	Active bool
}

// Populate with fake values
func (p *Property) Populate() {
	p.TotalArea = 180
	p.UtilArea = 140
	p.Floor = 9999
	p.Active = true

	p.Features.Barbecue = true
	p.Features.Garden = true
	p.Features.Pool = false

	p.Address.State = "SP"
	p.Address.City = "São Paulo"
	p.Address.Neighborhood = "Limão"
	p.Address.Street = "Deputado Emílio Carlos"
	p.Address.Number = "539"
}