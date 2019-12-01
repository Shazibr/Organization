package organization

func AddUser(_c Community, _user User) Community {
	if _c.nuser < 10 {
		_c.team = append(_c.team, _user)
	}
	_c.nuser = _c.nuser + 1
	return _c
}
