package organization

func New(_Name string, _IPaddress string, _User User) Community {
	c := Community{_Name, _IPaddress, make([]User, 0), 0, 0}
	c.team = append(c.team, _User)
	c.nuser = c.nuser + 1
	//var community C
	//C.Name = _Name
	//C.IPaddress = _IPaddress
	//C.team = make([]User, 0)
	//c.team[0] := User
	return c
}
