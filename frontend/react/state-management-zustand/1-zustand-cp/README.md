# Simple Todo App With Zustand

## Setup
```shell
npm install
```

## Running the project
```shell
npm start
```

## Testing manually
```shell
npm test
```

## Challenge
### Expected Outcome


https://user-images.githubusercontent.com/33412865/169703336-1c0dfa1d-9aaa-4047-9618-7fb63c0f18c5.mov



### 1. Create store for saving the todos(in array) state
__file:__ `src/store/todoStore.js`
__state name:__ `todos`
__array of object structure example:__ 
```TS
todos = [
    {
        id: 9013851123
        text: 'makan'
        isDone: true
    },
    {
        id: 9013851125
        text: 'minum'
        isDone: false
    },
]
```

### 2. Add function to the store
__file:__ `src/store/todoStore.js`

- `addTodo({id, text, isDone})` : function for adding new todo to array
- `removeTodo(id)` : function for removing todo (using `id` for parameter)
- `toggleTodo(id)`: unction for toggling todo is done or not (using `id` for parameter)

### 3. Use state and function from store

#### Add Todo Component
<<<<<<< HEAD
Add new todo \
__file:__ `src/components/AddTodo.js` \
__used items:__ `addTodo`

#### List Todos Component
List all todos \
__file:__ `src/components/ListTodos.js` \
__used items:__ `todos` 

#### Summary Todos Component
Summary of todos \
__file:__ `src/components/SummaryTodos.js` \
__used items:__ `todos` 

#### Todo Item Component
Item component for showing the todo text, deleting todo, toggling todo(done/not) \

__file:__ `src/components/TodoItem.js` \
=======
Add new todo
__file:__ `src/components/AddTodo.js`
__used items:__ `addTodo`

#### List Todos Component
List all todos
__file:__ `src/components/ListTodos.js`
__used items:__ `todos`

#### Summary Todos Component
Summary of todos

__file:__ `src/components/SummaryTodos.js`
__used items:__ `todos`

#### Todo Item Component
Item component for showing the todo text, deleting todo, toggling todo(done/not)

__file:__ `src/components/TodoItem.js`
>>>>>>> 4eb2273 (add new zustand demo and checkpoint (part-1))
__used items:__ `removeTodo`, `toggleTodo`
