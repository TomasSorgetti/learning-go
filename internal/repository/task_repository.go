package repository

import (
    "database/sql"
    "go-app/internal/models"
)

type TaskRepository struct {
    db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
    return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *models.Task) error {
    query := `INSERT INTO tasks (title, description, done) VALUES (?, ?, ?)`
    result, err := r.db.Exec(query, task.Title, task.Description, task.Done)
    if err != nil {
        return err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    task.ID = int(id)
    return nil
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
    query := `SELECT id, title, description, done FROM tasks`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var task models.Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}

func (r *TaskRepository) GetByID(id int) (*models.Task, error) {
    query := `SELECT id, title, description, done FROM tasks WHERE id = ?`
    var task models.Task
    err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Done)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return &task, nil
}

func (r *TaskRepository) Update(task *models.Task) error {
    query := `UPDATE tasks SET title = ?, description = ?, done = ? WHERE id = ?`
    _, err := r.db.Exec(query, task.Title, task.Description, task.Done, task.ID)
    return err
}

func (r *TaskRepository) Delete(id int) error {
    query := `DELETE FROM tasks WHERE id = ?`
    _, err := r.db.Exec(query, id)
    return err
}