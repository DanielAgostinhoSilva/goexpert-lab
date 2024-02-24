package commons

import "time"

type Problem struct {
	Status      int       `json:"status"`
	Timestamp   time.Time `json:"timestamp"`
	Title       string    `json:"title"`
	Detail      string    `json:"detail"`
	UserMessage string    `json:"userMessage"`
}

func NewProblem(status int, timestamp time.Time, title string, detail string, userMessage string) *Problem {
	return &Problem{Status: status, Timestamp: timestamp, Title: title, Detail: detail, UserMessage: userMessage}
}

func NewError500(detail, userMessage string) *Problem {
	return &Problem{
		Status:      500,
		Timestamp:   time.Now(),
		Title:       SystemError,
		Detail:      detail,
		UserMessage: userMessage,
	}
}

func NewInternalServerError(detail string) *Problem {
	return &Problem{
		Status:      500,
		Timestamp:   time.Now(),
		Title:       SystemError,
		Detail:      detail,
		UserMessage: "An unexpected internal system error has occurred. Please try again and if the problem persists, contact the system administrator.",
	}
}

func NewError422(detail, userMessage string) *Problem {
	return &Problem{
		Status:      422,
		Timestamp:   time.Now(),
		Title:       InvalidData,
		Detail:      detail,
		UserMessage: userMessage,
	}
}

func NewError404(detail, userMessage string) *Problem {
	return &Problem{
		Status:      404,
		Timestamp:   time.Now(),
		Title:       ResourceNotFound,
		Detail:      detail,
		UserMessage: userMessage,
	}
}
