package requests

type General struct {
	WorkCenter             	int     `json:"WorkCenter"`
	IsMarkedForDeletion		*bool   `json:"IsMarkedForDeletion"`
}
