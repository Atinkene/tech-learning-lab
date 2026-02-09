import { useState, useCallback} from "react"


function ReactUseCallback() {

  const [count1, setCount1] = useState(0)
  const [count2, setCount2] = useState(0)

  function Button({ onClick, children }) {
    console.log(`Rendering button: ${children}`)
    return (
      <button
        className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 cursor-pointer"
        onClick={onClick}
      >
        {children}
      </button>
    )
  }

  const handleClick1 = useCallback(() => {
    console.log("Button 1 clicked")
    setCount1(count1 + 1)
  }, [count1])

  const handleClick2 = useCallback(() => {
    setCount2(count2 + 1)
  }, [count2])    

  return (
    <div className="flex flex-col items-center justify-center w-full h-screen space-y-4">
      <Button onClick={handleClick1}>Count 1: {count1}</Button>
      <Button onClick={handleClick2}>Count 2: {count2}</Button>
    </div>
  )
}
export default ReactUseCallback