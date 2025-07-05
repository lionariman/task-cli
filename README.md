# task-cli
a CLI app to track your tasks and manage your to‑do list

---

## Adding a new task
```bash
task-cli add "Buy groceries"
```
**Output:** Task added successfully (ID: 1)

---

## Updating and deleting tasks
```bash
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1
```

---

## Marking a task as in progress or done
```bash
task-cli mark-in-progress 1
task-cli mark-done 1
```

---

## Listing tasks
```bash
task-cli list
```

---

## Listing tasks by status
```bash
task-cli list todo
task-cli list in-progress
task-cli list done
```
