package portal

import (
	"encoding/json"
	"fmt"
	"log"
)

/*

area
domain
region

*/

type Areas []Area

type Area struct {
	Id               int
	HasPortalSurface bool // Has surface portalability.
	Items            []Element
}

// Переход
type Transition struct {
}

type Chamber struct {
	Areas []Area
	Items []Element
}

func ExampleFillAreas() {
	chamber := Chamber{
		Areas: []Area{
			Area{
				Id:               1,
				HasPortalSurface: true,
				Items: []Element{
					Cube{},
				},
			},
		},
		Items: []Element{
			Cube{
				ParentId: 1,
			},
		},
	}

	data, err := json.MarshalIndent(chamber, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
