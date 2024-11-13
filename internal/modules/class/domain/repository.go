package classDomain

type ClassRepository interface {
	Create(userID, name, subject, grade string) (*ClassEntity, error)
	Join(userID, code string) (*ClassEntity, error)
	List(userID string, page, limit int32) ([]*ClassEntity, error)
}
