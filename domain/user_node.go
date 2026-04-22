package domain

type UserNode struct {
	IdUser      int
	NamaUser    string
	DateCreated string
	DateUpdated string
	Role        string
}

type LL_User struct {
	DataUser UserNode
	Linked   *LL_User
}
