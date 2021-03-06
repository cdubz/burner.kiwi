package data

import "errors"

// ErrMessageDoesntExist is returned by GetMessagesByID when it cant find that specific message
var ErrMessageDoesntExist = errors.New("message doesn't exist")

// Database lists methods needed to implement a db
type Database interface {
	SaveNewInbox(Inbox) error
	GetInboxByID(string) (Inbox, error)
	EmailAddressExists(string) (bool, error)
	SetInboxCreated(Inbox) error
	SetInboxFailed(Inbox) error
	SaveNewMessage(Message) error
	GetMessagesByInboxID(string) ([]Message, error)
	GetMessageByID(string, string) (Message, error)
}
