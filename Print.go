package organization

import (
	"fmt"
	"strconv"
)

func Print(_c Community) {
	for i := 0; i < _c.nuser; i++ {
		fmt.Println("User Name: " + _c.team[i].Uname + " IP Address: " + _c.team[i].Uaddress + " Leadership Status: " + strconv.FormatBool(_c.team[i].LStatus) + " Amount: " + fmt.Sprintf("%f", _c.team[i].Amount) + "")
	}
}
