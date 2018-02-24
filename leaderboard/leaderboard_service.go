package leaderboard

// Service leaderboard
type Service interface {
	GetLeaderboard()
}

type service struct {
	printer Printer
}

func (s *service) GetLeaderboard() {
	s.printer.Print()
}

// NewService creates a service
func NewService(printer Printer) Service {
	return &service{
		printer,
	}
}
