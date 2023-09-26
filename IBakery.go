package golangAbstractFactory

import "fmt"

type iBakery interface {
  makeCroissant() iCroissant 
  makePastry() iPastry
}

func getBakery(brand string) (iBakery , error){
  if brand == "casata"{
    return &casata{} , nil
  }
  if brand == "dempi"{
    return &dempi{} , nil
  }
  return nil , fmt.Errorf("No brand exist with name provided")
}
