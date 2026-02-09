import { useEffect, useRef, useState } from "react"

function UseRef() {
    const [inputValue, setInputValue] = useState("")
    const inputRef = useRef("")

    useEffect(() => {
        inputRef.current = inputValue;
    }, [inputValue])

    return (
        <div className="flex flex-col items-center justify-center w-full h-screen space-y-4">
            <input 
                ref={inputRef} 
                type="text" 
                placeholder="Focus me!" 
                className="border border-gray-300 rounded px-4 py-2" 
                onChange={(e) => setInputValue(e.target.value)} 
                value={inputValue} 
            />

            <p>current value : {inputValue}</p>
            <p>Previous value : {inputRef.current}</p>
        </div>
    )
}

export default UseRef