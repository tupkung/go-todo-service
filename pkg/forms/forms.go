package forms

type NewTask struct {
	Title    string
	Failures map[string]string
}

type EditTask struct {
	UID      string
	Title    string
	Complete bool
	Failures map[string]string
}

func (t *NewTask) Valid() bool {
	t.Failures = make(map[string]string)

	return len(t.Failures) == 0
}

func (t *EditTask) Valid() bool {
	t.Failures = make(map[string]string)

	return len(t.Failures) == 0
}
