package planet

type COGCProgram struct {
	Type         string `json:"Type"`
	StartEpochMs int64  `json:"StartEpocMs"`
	EndEpocMs    int64  `json:"EndEpocMs"`
}
