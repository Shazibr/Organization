package organization

import (
	"fmt"
	"strconv"
)

func Vote(_c Community, _user User) Community {
	var option int
	if _c.nuser > 3 {
		fmt.Println("Voting Candidates")
		for i := 0; i < _c.nuser; i++ {
			fmt.Println(strconv.Itoa(i) + ") User Name: " + _c.team[i].Uname + "")
		}
		for true {
			fmt.Scan(&option)
			if _c.team[option].Uname != _user.Uname {
				_c.team[option].Votes = _c.team[option].Votes + 1
				break
			} else {
				fmt.Println("You cant vote for yourself, Please enter again")
			}
		}
	}
	return _c
}
