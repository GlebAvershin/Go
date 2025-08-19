package todo

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	list := List{
		tasks: make(map[string]Task),
	}

	return &list
}

func (l *List) AddTask(task Task) {
	l.tasks[task.Title] = task
}

func (l *List) ListTasks() map[string]Task {
	return l.tasks
}

func (l *List) DoneTask(title string) string {
	task, ok := l.tasks[title]
	if !ok {
		return TaskNotFound
	}

	task.Done()
	l.tasks[title] = task

	return ""

}

func (l *List) DeleteTask(title string) string {
	_, ok := l.tasks[title]
	if !ok {
		return TaskNotFound
	}

	delete(l.tasks, title)

	return ""
}
