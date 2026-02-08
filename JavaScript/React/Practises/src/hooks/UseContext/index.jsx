import { useState, createContext, useContext, useEffect } from "react"

/** 
 * useContext allows you to share state across components without having to pass props down manually at every level.
 * It is often used in combination with createContext to create a context object that can be accessed by any component within the provider.
 */
const MyContext = createContext()

function ReactUseContext() {
  function ChildComponent() {
    const { value, setValue } = useContext(MyContext)

    useEffect(() => {
        setValue("Nicolas")
    }, [])

    return (
      <div>
        <h1 className="text-2xl font-bold">Context Value: {value}</h1>
        <ChildComponent1 />
      </div>
    )
  }

  function ChildComponent1() {
    const contextValue = useContext(MyContext)
    return <ChildComponent2 />
  }

  function ChildComponent2() {
    const { value, setValue } = useContext(MyContext)
    
    /**
     * To update the context value, you can call the setValue function provided by the context. In this example, we use useEffect to update the context value after 2 seconds. When the context value is updated, all components that consume the context will re-render with the new value.
     * In this case, after 2 seconds, the context value will change to "BASSENE", and ChildComponent2 will display this updated value.
     * Note that the useEffect hook is used to perform side effects in functional components, and it runs after the component has rendered. The empty dependency array [] ensures that the effect runs only once when the component mounts.
     * When the context value is updated, all components that consume the context will re-render with the new value. In this case, after 2 seconds, the context value will change to "BASSENE", and ChildComponent2 will display this updated value.
     * crearTime allows you to set a timer that executes a function after a specified delay. In this example, we use setTimeout to update the context value after 2 seconds. When the timer expires, the setValue function is called to update the context value to "BASSENE". The return statement in useEffect is used to clean up the timer when the component unmounts, preventing memory leaks. 
     */
    useEffect(() => {
      const timer = setTimeout(() => {
        setValue("BASSENE")
      }, 2000) 
      
      return () => clearTimeout(timer)
    }, [])

    return (
      <h1 className="text-2xl font-bold">Context Value: {value}</h1>
    )
  }

  function MainComponent() {
    const [value, setValue] = useState("Massina")
    return (
      <MyContext.Provider value={{ value, setValue }}>
        <ChildComponent />
      </MyContext.Provider>
    )
  }

  return (
    <div className="flex flex-col items-center justify-center w-full h-screen space-y-4">
      <MainComponent />
    </div>
  )
}

export default ReactUseContext