package dependencyInjection

import "testing"

/***
NOTE: Goals
 1. Initialization - should be initizliaed only once.
2. Testing - How to write independent test cases for modules using an external dependency
3. State - Easy way of exposing the state.
*/

// NOTE: https://parserdigital.com/dependency-injection-and-unit-testing-in-golang/
// Approach 1

type Service1 struct {
}

// ISSUES: maintainability, testability, how to mock.
func (s *Service) doTheThing() error {
	// initializde the db connection
	// store on the db.
	// initialize the email service
	// send email
}

// NOTE: Approach 2
// 3 components, one that handles DB logic
// one that handles email logic
// one that handles the main logic.

type Db struct {
}

func (db *Db) connect() error {
	// opens the connection
}

func (db *Db) store() error {
	// store the db
}

type EmailSender struct {
}

func (ex *EmailSender) connect() error {
	// connects with the email service
}

func (es *EmailSender) sendEmail() error {
	// send email
}

type Service2 struct {
	// Need a DB
	Db *Db
	// Need an EmailSender
	EmailSender *EmailSender
}

func (s *Service2) doTheThing() error {
	// store on the db
	s.Db
	// send email
	s.EmailSender
}

// constructor for dependency injection
// NOTE: issues here cant still send the mocks here.
func NewServiceX(db *Db, emailSender *EmailSender) *Service2 {
	return &Service2{
		Db:          db,
		EmailSender: emailSender,
	}
}

// Approach 3 - final approach
type Service3 struct {
	Db interface {
		Store(s something) error
	}

	EmailSender interface {
		Send(email Email) error
	}
}

// New initializes our service object with dep
func New(db *Db, sender *EmailSender) *Service3 {
	return &Service3{
		Db:          db,
		EmailSender: sender,
	}
}

func (s *Service3) doTheThing() error {
	// store on the db
	s.Db
	// send email
	s.EmailSender
}

// NOTE: UTS
type EmailSenderMock struct {
}

func (es *EmailSenderMock) Send(email Email) error {

}

type DbMock struct {
}

func (db *DbMock) Store(s something) error {

}

func TestService_doThing(t *testing.T) {
	dbMock := DbMock{}
	emailSenderMock := EmailSenderMock{}

	service := New(dbMock, emailSenderMock)
	err := service.doTheThing()
}
