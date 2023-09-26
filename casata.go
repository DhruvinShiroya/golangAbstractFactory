package github.com/DhruvinShiroya/golangAbstractFactory

type casata struct {
  
}

func (c *casata) makeCroissant() iCroissant {
  return &casataCroissant{
    croissant: croissant{
      flavour: "cinemon",
      size: "midium",
    },
  }
}

func (c *casata) makePastry() iPastry {

  return &casataPastry{
    pastry : pastry{
      category: "shortcrust",
    },
  }

}


