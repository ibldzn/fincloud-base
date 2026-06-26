package fincloud

type AuthorizationModel struct {
	Locations []AuthLabel `json:"locationid"`
	Roles     []AuthLabel `json:"roleid"`
}

type AuthLabel struct {
	ID          string `json:"id"`
	Description string `json:"descr"`
}
