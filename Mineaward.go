package organization

func Mineaward(_c Community, award float32) Community {
	var Leaderaward float32
	var Useraward float32
	Leaderaward = ((award / float32(_c.nuser)) + (1 / (float32(_c.nuser) - 1)))
	Useraward = (award - Leaderaward) / (float32(_c.nuser) - 1)
	for i := 0; i < _c.nuser; i++ {
		if _c.team[i].LStatus == false {
			_c.team[i].Amount += Useraward
		} else if _c.team[i].LStatus == true {
			_c.team[i].Amount += Leaderaward
		}
	}
	return _c
}
