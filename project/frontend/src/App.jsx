import { useState, useEffect } from 'react'
import axios from 'axios'
const baseUrl = '/todos'

function App() {
  const [todos, setTodos] = useState([])

  useEffect(() => {
    const getTodos = async () => {
      const response = await axios.get(baseUrl)
      setTodos(response.data)
    }
    getTodos()
  }, [])

  const todoList = () => {
    return !todos ? null : (
      <ul>
      {todos.map((todo) =>
        <li key={todo.id}>{todo.content}</li>
      )}
      </ul>
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
