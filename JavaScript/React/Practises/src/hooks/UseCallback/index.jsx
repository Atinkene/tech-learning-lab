import { useState } from "react"
function ReactUseCallback() {
  const [count, setCount] = useState(0)

  function handleClick() {
    setCount(count + 1)
  }

  return (
    <div className="flex flex-col items-center justify-center w-full h-screen space-y-4">
      <h1 className="text-2xl font-bold">Count: {count}</h1>
      <button
        className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
        onClick={handleClick}
      >
        Increment
      </button>
    </div>
  )
}
export default ReactUseCallback