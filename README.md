coa
===============

coa = **Co**mposable **A**ction

Overview
--------------

Package coa provides the action executor and interfaces for composable action.

Example
--------------

```go
package main

import (
	"fmt"

	"github.com/kazukgw/coa"
)

type Context struct {
	actionGroup coa.ActionGroup
}

func (c *Context) ActionGroup() coa.ActionGroup {
	return c.actionGroup
}

type ActionPrint1 struct {
}

func (t1 ActionPrint1) Do(ctx coa.Context) error {
	fmt.Println(1)
	return nil
}

type ActionPrint2 struct {
}

func (t2 ActionPrint2) Do(ctx coa.Context) error {
	fmt.Println(2)
	return nil
}

type ActionGroup struct {
	ActionPrint1
	coa.DoSelf
	ActionPrint2
}

func (ag *ActionGroup) Do(ctx coa.Context) error {
	fmt.Println("do")
	return nil
}

func (ag *ActionGroup) HandleError(ctx coa.Context, err error) error {
	return nil
}

func main() {
	ag := &ActionGroup{}
	ctx := &Context{ag}
	coa.Exec(ctx)
	// Output:
	// 1
	// do
	// 2
}
```
