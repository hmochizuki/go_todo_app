package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hmochizuki/go_todo_app/entity"
	"github.com/hmochizuki/go_todo_app/store"
)

type AddTask struct {
	Store     *store.TaskStore
	validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var b struct {
		Title string `json:"title" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJson(ctx, w, ErrResponse{Message: "failed to decode request body"}, http.StatusInternalServerError)
		return
	}

	if err := at.validator.Struct(b); err != nil {
		RespondJson(ctx, w, &ErrResponse{Message: "validation failed", Details: []string{err.Error()}}, http.StatusBadRequest)
		return
	}

	t := &entity.Task{
		Title:     b.Title,
		Status:    "todo",
		CreatedAt: time.Now(),
	}
	id, err := at.Store.Add(t)

	if err != nil {
		RespondJson(ctx, w, ErrResponse{Message: "failed to add task"}, http.StatusInternalServerError)
		return
	}

	rsp := struct {
		ID entity.TaskID `json:"id"`
	}{ID: id}

	RespondJson(ctx, w, rsp, http.StatusOK)
}
