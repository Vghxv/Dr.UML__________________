package umldiagram

import (
	"time"

	"Dr.uml/backend/component"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
)

type baseCommand struct {
	diagram *UMLDiagram
	before  time.Time
	after   time.Time
}

func (cmd *baseCommand) GetBefore() time.Time {
	return cmd.before
}

func (cmd *baseCommand) GetAfter() time.Time {
	return cmd.after
}

// add component
type addComponentCommand struct {
	baseCommand
	component component.Component
}

func (cmd *addComponentCommand) Execute() duerror.DUError {
	return cmd.diagram.addComponent(cmd.component)
}

func (cmd *addComponentCommand) Unexecute() duerror.DUError {
	return cmd.diagram.removeComponent(cmd.component)
}

// select
type removeSelectedComponentCommand struct {
	baseCommand
	components map[component.Component]bool
}

func (cmd *removeSelectedComponentCommand) Execute() duerror.DUError {
	return cmd.diagram.removeComponents(cmd.components)
}

func (cmd *removeSelectedComponentCommand) Unexecute() duerror.DUError {
	return cmd.diagram.addComponents(cmd.components)
}

type selectAllCommand struct {
	baseCommand
	components map[component.Component]bool
	newValue   bool
}

func (cmd *selectAllCommand) Execute() duerror.DUError {
	return cmd.diagram.selectAll(cmd.components, cmd.newValue)
}

func (cmd *selectAllCommand) Unexecute() duerror.DUError {
	return cmd.diagram.selectAll(cmd.components, !cmd.newValue)
}

// simple setters
type setterCommand struct {
	baseCommand
	component component.Component
	execute   func() duerror.DUError
	unexecute func() duerror.DUError
}

func (cmd *setterCommand) Execute() duerror.DUError {
	return cmd.execute()
}

func (cmd *setterCommand) Unexecute() duerror.DUError {
	return cmd.unexecute()
}

// move gadget
type moveGadgetCommand struct {
	baseCommand
	gadget   *component.Gadget
	newPoint utils.Point
	oldPoint utils.Point
}

func (cmd *moveGadgetCommand) Execute() duerror.DUError {
	return cmd.diagram.moveGadget(cmd.gadget, cmd.newPoint)
}

func (cmd *moveGadgetCommand) Unexecute() duerror.DUError {
	return cmd.diagram.moveGadget(cmd.gadget, cmd.oldPoint)
}
