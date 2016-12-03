// Copyright 2016 Marapongo, Inc. All rights reserved.

package compiler

import (
	"github.com/marapongo/mu/pkg/ast"
	"github.com/marapongo/mu/pkg/compiler/backends"
	"github.com/marapongo/mu/pkg/util"
)

// Context holds all state available to any templates or code evaluated at compile-time.
type Context struct {
	Cluster    *ast.Cluster    // the cluster that we will deploy to.
	Arch       backends.Arch   // the target cloud architecture.
	Properties ast.PropertyBag // properties supplied at stack construction time.
}

// NewContext returns a new, empty context.
func NewContext() *Context {
	return &Context{
		Properties: make(ast.PropertyBag),
	}
}

// WithClusterArch returns a clone of this Context with the given cluster and architecture attached to it.
func (c *Context) WithClusterArch(cl *ast.Cluster, a backends.Arch) *Context {
	util.Assert(cl != nil)
	return &Context{
		Cluster:    cl,
		Arch:       a,
		Properties: c.Properties,
	}
}

// WithProps returns a clone of this Context with the given properties attached to it.
func (c *Context) WithProps(props ast.PropertyBag) *Context {
	if props == nil {
		props = make(ast.PropertyBag)
	}
	return &Context{
		Cluster:    c.Cluster,
		Arch:       c.Arch,
		Properties: props,
	}
}