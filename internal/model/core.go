package model

import (
	"fmt"

	"gopkg.in/validator.v2"
)

type UserID string

func (u UserID) String() string {
	return string(u)
}

type MessageID string

func (m MessageID) String() string {
	return string(m)
}

type DialogID string

func (d DialogID) String() string {
	return string(d)
}

type Message struct {
	ID       MessageID `validator:"nonzero"`
	DialogID DialogID  `validator:"nonzero"`
	SenderID UserID    `validator:"nonzero"`
	Text     string    `validator:"nonzero"`
}

func (m *Message) Validate() error {
	if err := validator.Validate(m); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

type Dialog struct {
	ID              DialogID       `validator:"nonzero"`
	ParticipantsIDs []*Participant `validator:"nonzero"`
}

func (d *Dialog) Validate() error {
	if err := validator.Validate(d); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

type ParticipantID string

func (p ParticipantID) String() string {
	return string(p)
}

type Participant struct {
	ID     ParticipantID `validator:"nonzero"`
	UserID UserID        `validator:"nonzero"`
}

func (p *Participant) Validate() error {
	if err := validator.Validate(p); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}
