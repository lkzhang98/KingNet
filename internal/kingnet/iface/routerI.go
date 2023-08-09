// Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/lkzhang98/kingnet.

package iface

type RouterI interface {
	// PreHandle defines a pre-handle function for the router
	PreHandle(request RequestI)
	// Handle defines a handle function for the router
	Handle(request RequestI)
	// PostHandle defines a post-handle function for the router
	PostHandle(request RequestI)
}
