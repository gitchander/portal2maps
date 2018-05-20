package portal

type Player struct {
	areaId int
}

type Element interface {
	Name() string
}

type Portal struct {
}

type CubeDropper struct {
}

func (CubeDropper) Name() string {
	return "CubeDropper"
}

type Cube struct {
	ParentId int
}

func (Cube) Name() string {
	return "Cube"
}

type Fizzler struct {
}

func (Fizzler) Name() string {
	return "Fizzler"
}

type LaserField struct {
}

func (LaserField) Name() string {
	return "LaserField"
}

type Stairs struct {
}

type PedestalButton struct {
}

func (PedestalButton) Name() string {
	return "PedestalButton"
}
