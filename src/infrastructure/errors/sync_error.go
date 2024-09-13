package errors

type SyncError struct {
	Message string
}

func NewSyncError(message string) SyncError {
	return SyncError{
		Message: message,
	}
}

func (s SyncError) Error() string {
	return s.Message
}
