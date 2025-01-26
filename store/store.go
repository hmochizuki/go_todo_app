package store

import (
	"errors"

	"github.com/hmochizuki/go_todo_app/entity"
)

var (
	Tasks       = &TaskStore{Tasks: map[entity.TaskID]*entity.Task{}}
	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	// TODO: 動作確認用にexportしている
	LastId entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastId++
	t.ID = ts.LastId
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) All() entity.Tasks {
	tasks := make([]*entity.Task, len(ts.Tasks))

	for i, t := range ts.Tasks {
		tasks[i-1] = t
	}
	return tasks
}
