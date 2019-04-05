package group

type Repository interface {
	FindAll() ([]*Group, error)
	Save(user *Group)
}
