package umldiagram

import (
	"time"

	"Dr.uml/backend/component"
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
