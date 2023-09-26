package abstractfactory 

type iCroissant interface {
  setFlavour(f string)
  getFlavour() string
  setSize(s string)
  getSize() string

}

type croissant struct {
  flavour string
  size string
}

func (c *croissant) setFlavour(f string) {
  c.flavour = f
}


func (c *croissant) getFlavour() string {
  return c.flavour 
}


func (c *croissant) setSize(s string) {
  c.size= s
}


func (c *croissant) getSize() string {
  return c.size
}
