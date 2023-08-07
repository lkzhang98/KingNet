package iface

type RouterI interface {
	// PreHandle defines a pre-handle function for the router
	PreHandle(request RequestI)
	// Handle defines a handle function for the router
	Handle(request RequestI)
	// PostHandle defines a post-handle function for the router
	PostHandle(request RequestI)
}
