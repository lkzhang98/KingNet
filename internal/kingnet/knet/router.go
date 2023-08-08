package knet

import "go.mod/internal/kingnet/iface"

// BaseRouter 实现router时，先嵌入这个基类，然后根据需要对这个基类的方法进行重写
// 适用于只有部分方法的类，可以不用重写所有的方法.
type BaseRouter struct{}

func (br *BaseRouter) PreHandle(req iface.RequestI)  {}
func (br *BaseRouter) Handle(req iface.RequestI)     {}
func (br *BaseRouter) PostHandle(req iface.RequestI) {}
