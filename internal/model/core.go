package model

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
	ID          MessageID
	DialogID    DialogID
	SenderID    UserID
	RecipientID UserID
	Text        string
}

func (m *Message) Validate() error {
	return nil
}
