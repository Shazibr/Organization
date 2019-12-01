package organization

import "fmt"

func Election(_c Community) Community {
	var maxvotes int
	var teampos int
	maxvotes = -1
	_c.team[_c.Currentleader].LStatus = false
	if _c.nuser > 2 {
		for i := 0; i < _c.nuser; i++ {
			_c = Vote(_c, _c.team[i])
		}
		for i := 0; i < _c.nuser; i++ {
			if maxvotes < _c.team[i].Votes {
				maxvotes = _c.team[i].Votes
				teampos = i
			}
		}
		_c.team[teampos].LStatus = true
	} else {
		fmt.Println("Minimum User for election is 3")
	}
	return _c
}
