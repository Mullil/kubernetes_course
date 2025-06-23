import { useState } from "react";

function App() {
  const [todos, setTodo] = useState([
    "Learn Javascript",
    "Learn React",
    "Build a project"
  ])


  const todoList = () => {
    return (
      <ul>
      {todos.map((todo, idx) =>
        <li key={idx}>{todo}</li>
      )}
      </ul>
    )
  }

  const handleSubmit = (event) => {
    event.preventDefault()
    const todo = event.target.todo.value
    event.target.todo.value = ''
    if (todo.length <= 140) setTodo(todos.concat(todo))
  }

  return (
    <div>
      The project App
      <p><img style={{ height: 300, width: 300 }} src='/files/image' alt='Random'/></p>
      <form onSubmit={ handleSubmit }>
        <div><input name="todo"></input></div>
        <button type="submit">Create todo</button>
      </form>
      {todoList()}

      DevOps with Kubernetes 2025
    </div>
    
  )
}
export default App;
