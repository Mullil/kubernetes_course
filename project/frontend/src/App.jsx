import { useState, useEffect } from 'react'
import axios from 'axios'
const baseUrl = '/todos'

function App() {
  const [todos, setTodos] = useState([])

  useEffect(() => {
    const getTodos = async () => {
      const response = await axios.get(baseUrl)
      setTodos(Array.isArray(response.data) ? response.data : [])
    }
    getTodos()
  }, [])

  const handleDone = async (id) => {
    try {
      const response = await axios.put(`${baseUrl}/${id}`)
      setTodos(todos.map(t => t.id !== id ? t : response.data))
    } catch (error) {
        console.error("Failed to mark todo as done", error)
    }
  }

  const todoList = () => {
    return !todos ? null : (
      <div>
        <h3>Todo</h3>
        <ul>
        {todos.filter(t => t.done === false).map((todo) =>
          <li key={todo.id}>
            {todo.content} <button onClick={() => handleDone(todo.id)}>Mark as done</button>
          </li>
        )}
        </ul>
        <h3>Done</h3>
        <ul>
        {todos.filter(t => t.done === true).map((todo) =>
          <li key={todo.id}>{todo.content}</li>
        )}
        </ul>
      </div>
    )
  }

  const handleSubmit = async (event) => {
    event.preventDefault()
    const todo = event.target.todo.value
    event.target.todo.value = ''
    if (todo.length <= 140) {
      const response = await axios.post(baseUrl, { content: todo })
      setTodos(todos.concat(response.data))
    }
  }

  return (
    <div>
      <h2>The project App</h2>
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
