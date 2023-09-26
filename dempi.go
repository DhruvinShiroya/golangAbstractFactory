package github.com/DhruvinShiroya/golangAbstractFactory

type dempi struct {
}

func (d *dempi) makeCroissant() iCroissant {
	return &dempiCroissant{
		croissant: croissant{
			flavour: "chocolate",
			size:    "midium",
		},
	}
}

func (d *dempi) makePastry() iPastry {

	return &dempiPastry{
		pastry: pastry{
			category: "shortcrust",
		},
	}

}
