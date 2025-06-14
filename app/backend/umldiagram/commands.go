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

// set parent association
type setParentStartCommand struct {
	baseCommand
	association *component.Association
	stNew       *component.Gadget
	stOld       *component.Gadget
	stRatioNew  [2]float64
	stRatioOld  [2]float64
}

func (cmd *setParentStartCommand) Execute() duerror.DUError {
	return cmd.diagram.setParentStartAssociation(
		cmd.association,
		cmd.stNew,
		cmd.stRatioNew,
	)
}

func (cmd *setParentStartCommand) Unexecute() duerror.DUError {
	return cmd.diagram.setParentStartAssociation(
		cmd.association,
		cmd.stOld,
		cmd.stRatioOld,
	)
}

type setParentEndCommand struct {
	baseCommand
	association *component.Association
	enNew       *component.Gadget
	enOld       *component.Gadget
	enRatioNew  [2]float64
	enRatioOld  [2]float64
}

func (cmd *setParentEndCommand) Execute() duerror.DUError {
	return cmd.diagram.setParentEndAssociation(
		cmd.association,
		cmd.enNew,
		cmd.enRatioNew,
	)
}

func (cmd *setParentEndCommand) Unexecute() duerror.DUError {
	return cmd.diagram.setParentEndAssociation(
		cmd.association,
		cmd.enOld,
		cmd.enRatioOld,
	)
}

// add/remove attribute
type addAttributeGadgetCommand struct {
	baseCommand
	gadget  *component.Gadget
	content string
	section int
	index   int
}

func (cmd *addAttributeGadgetCommand) Execute() duerror.DUError {
	return cmd.diagram.addAttributeGadget(
		cmd.gadget,
		cmd.section,
		cmd.index,
		cmd.content,
	)
}

func (cmd *addAttributeGadgetCommand) Unexecute() duerror.DUError {
	return cmd.diagram.removeAttributeGadget(
		cmd.gadget,
		cmd.section,
		cmd.index,
	)
}

type removeAttributeGadgetCommand struct {
	baseCommand
	gadget  *component.Gadget
	content string
	section int
	index   int
}

func (cmd *removeAttributeGadgetCommand) Execute() duerror.DUError {
	return cmd.diagram.removeAttributeGadget(
		cmd.gadget,
		cmd.section,
		cmd.index,
	)
}

func (cmd *removeAttributeGadgetCommand) Unexecute() duerror.DUError {
	return cmd.diagram.addAttributeGadget(
		cmd.gadget,
		cmd.section,
		cmd.index,
		cmd.content,
	)
}

type addAttributeAssociationCommand struct {
	baseCommand
	association *component.Association
	content     string
	ratio       float64
	index       int
}

func (cmd *addAttributeAssociationCommand) Execute() duerror.DUError {
	return cmd.diagram.addAttributeAssociation(
		cmd.association,
		cmd.index,
		cmd.ratio,
		cmd.content,
	)
}

func (cmd *addAttributeAssociationCommand) Unexecute() duerror.DUError {
	return cmd.diagram.removeAttributeAssociation(
		cmd.association,
		cmd.index,
	)
}

type removeAttributeAssociationCommand struct {
	baseCommand
	association *component.Association
	content     string
	ratio       float64
	index       int
}

func (cmd *removeAttributeAssociationCommand) Execute() duerror.DUError {
	return cmd.diagram.removeAttributeAssociation(
		cmd.association,
		cmd.index,
	)
}

func (cmd *removeAttributeAssociationCommand) Unexecute() duerror.DUError {
	return cmd.diagram.addAttributeAssociation(
		cmd.association,
		cmd.index,
		cmd.ratio,
		cmd.content,
	)
}
