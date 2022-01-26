package abstractfactory

type CrudFactory interface {
	GetListById(id int)
}
