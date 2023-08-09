// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package knet

import "KingNet/internal/kingnet/iface"

// BaseRouter 实现router时，先嵌入这个基类，然后根据需要对这个基类的方法进行重写
// 适用于只有部分方法的类，可以不用重写所有的方法.
type BaseRouter struct{}

func (br *BaseRouter) PreHandle(req iface.RequestI)  {}
func (br *BaseRouter) Handle(req iface.RequestI)     {}
func (br *BaseRouter) PostHandle(req iface.RequestI) {}
