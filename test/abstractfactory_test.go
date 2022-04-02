package test

import (
	"go-design-pattern/createpattern/abstractfactory"
	"testing"
)

func getMainAndDetail(factory abstractfactory.DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func TestRdbFactory(t *testing.T) {
	var factory abstractfactory.DAOFactory
	factory = &abstractfactory.RDBDAOFactory{}
	getMainAndDetail(factory)
	// Output:
	// rdb main save
	// rdb detail save
}

func TestXmlFactory(t *testing.T) {
	var factory abstractfactory.DAOFactory
	factory = &abstractfactory.XMLDAOFactory{}
	getMainAndDetail(factory)
	// Output:
	// xml main save
	// xml detail save
}
