package data

type Database interface {
	Create(todo Todo) error

	GetById(id int) (Todo, error)

	GetAll() ([]Todo, error)

	Update(todo Todo) error
}

type SliceDatabase struct {
	todos []Todo
}

func NewSliceDatabase() *SliceDatabase {
	return &SliceDatabase{
		todos: []Todo{
			{
				Id:    0,
				Name:  "llllllll",
				State: true,
			},
		},
	}
}

func (db *SliceDatabase) Create(todo Todo) error {
	todo.Id = len(db.todos) - 1
	db.todos = append(db.todos, todo)
	return nil
}

func (db *SliceDatabase) GetById(id int) (Todo, error) {
	return db.todos[id], nil
}

func (db *SliceDatabase) GetAll() ([]Todo, error) {
	return db.todos, nil
}

func (db *SliceDatabase) Update(todo Todo) error {
	db.todos[todo.Id] = todo
	return nil
}
