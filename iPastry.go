package abstractfactory 

type iPastry interface {
  setCategory(f string)
  getCategory() string

}

type pastry struct {
  category string
}

func (p *pastry) setCategory(s string) {
  p.category = s
}


func (p *pastry) getCategory() string {
  return p.category
}


