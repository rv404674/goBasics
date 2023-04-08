package dependencyInjection

// https://medium.com/avenue-tech/dependency-injection-in-go-35293ef7b6
/**
	NOTE: Rule 1 always inject interfaces not structs. So that later on they can be mocked
Also the process is to make an interface (define methods here), and then make a struct a concrete impelmentation of it.
NOTE: Constructor injection + dependency injection
NOTE: can use empty interface as well
*/

// Approach 1 - creating your dependencies inside the implementation
// FIXME - shittiest of all
//type MyRepository struct{}
//type MyLogger struct{}
//type MyMessageBroker struct{}
//
//type MyService struct{}
//
//func (mS *MyService) compute(id string) error {
//	logger := &MyLogger{}
//	repository := &MyRepository{}
//
//	_ = logger
//	_ = repository
//	return nil
//}

// NOTE: Approach 2 - SOLID (D of SOLID)
type MyRepo interface{}
type CustomRepo struct {
}

type MyLogger interface{}
type MessageBroker interface{}

type MyService struct {
	logger MyLogger
	repo   MyRepo
	broker MessageBroker
}

func (ns2 *MyService) compute(id string) error {
}

// NOTE: constructor injection
// for this all the dependencies need to be ready, else you will get an error
func NewService(logger MyLogger, repo MyRepo, broker MessageBroker) *MyService {
	return &MyService{
		logger: logger,
		repo:   repo,
		broker: broker,
	}
}

// Method injection
func (mS2 *MyService) SetLogger(logger MyLogger) {
	mS2.logger = logger
}

// for now
func main() {
	logger := MyLogger(nil)
	broker := MessageBroker(nil)
	repo := CustomRepo{}

	service := NewService(logger, broker, repo)
	_ = service

}
