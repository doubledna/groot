## gctl
groot command tools, used to list, create, update or delete tasks

### create task
`gctl create -f task.yaml`
`gctl create task --x y`

### get task list or one task
`gctl get task`
`gctl get task taskName`

### update task
`gctl apply -f task.yaml`
`gctl edit task taskName`

### delete task
`gctl delete task taskName`
