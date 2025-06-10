package command

import (
	"time"

	"Dr.uml/backend/utils/duerror"
)

const CMD_LIMIT = 20

type Command interface {
	Execute() duerror.DUError
	Unexecute() duerror.DUError
	GetBefore() time.Time
	GetAfter() time.Time
}

type Manager struct {
	undoStack    []Command
	redoStack    []Command
	lastModified time.Time
	limit        int
}

func NewManager(lastModified time.Time) *Manager {
	return &Manager{
		undoStack:    make([]Command, 0, CMD_LIMIT),
		redoStack:    make([]Command, 0, CMD_LIMIT),
		lastModified: lastModified,
		limit:        CMD_LIMIT,
	}
}

func (m *Manager) GetLastModified() time.Time {
	return m.lastModified
}

func (m *Manager) Execute(cmd Command) duerror.DUError {
	if cmd == nil {
		return duerror.NewInvalidArgumentError("command is nil")
	}
	if err := cmd.Execute(); err != nil {
		return err
	}
	if len(m.undoStack) == m.limit {
		m.undoStack = m.undoStack[1:]
	}
	m.undoStack = append(m.undoStack, cmd)
	m.redoStack = nil
	m.lastModified = cmd.GetAfter()
	return nil
}

func (m *Manager) Undo() duerror.DUError {
	if len(m.undoStack) == 0 {
		return duerror.NewInvalidArgumentError("no undo command")
	}
	cmd := m.undoStack[len(m.undoStack)-1]
	m.undoStack = m.undoStack[:len(m.undoStack)-1]
	if err := cmd.Unexecute(); err != nil {
		return err
	}
	m.redoStack = append(m.redoStack, cmd)
	m.lastModified = cmd.GetBefore()
	return nil
}

func (m *Manager) Redo() duerror.DUError {
	if len(m.redoStack) == 0 {
		return duerror.NewInvalidArgumentError("no redo command")
	}
	cmd := m.redoStack[len(m.redoStack)-1]
	m.redoStack = m.redoStack[:len(m.redoStack)-1]
	if err := cmd.Execute(); err != nil {
		return err
	}
	m.undoStack = append(m.undoStack, cmd)
	m.lastModified = cmd.GetAfter()
	return nil
}
