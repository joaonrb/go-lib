package async

type Status int

const (
	StatusNotStarted Status = iota
	StatusWorking
	StatusFinished
	StatusError
)

func (status Status) String() string {
	switch status {
	case StatusNotStarted:
		return "NotStarted"
	case StatusWorking:
		return "Working"
	case StatusFinished:
		return "Finished"
	case StatusError:
		return "Error"
	default:
		return "Unknown"
	}
}
