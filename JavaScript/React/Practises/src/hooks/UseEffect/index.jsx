import { use, useState, useEffect } from "react"

function ReactUseEffect() {
  const [count, setCount] = useState(0)
  const [data, setData] = useState(null)


  useEffect(() => {
    fetch("http://192.168.1.35:8000/Students")
      .then(response => {
        if (!response.ok) {
          throw new Error('Fichier non trouvé')
        }
        return response.json()
      })
      .then(data => setData(data))
      .catch(error => {
        console.error('Erreur:', error)
        setError(error.message)
      })
  }, [])

  return (
    <div className="flex flex-col items-center justify-center w-full h-screen space-y-4">
      <h1 className="text-2xl font-bold">Student's info</h1>
        {data ? (
            <div className="space-x-2 flex">
                {data.students.map((student, index) => (
                    <div key={index} className="p-4 border rounded">
                        <h2 className="text-lg font-semibold">{student.firstname}</h2>
                        <p>Age: {student.age}</p>
                        <p>Email: {student.email}</p>
                        <p>Address: {student.address.street}, {student.address.city}, {student.address.state} {student.address.zip}</p>
                        <p>Courses:</p>
                        <ul className="list-disc list-inside">
                            {student.courses.map((course, courseIndex) => (
                                <li key={courseIndex}>{course.name}: {course.grade}</li>
                            ))}
                        </ul>
                    </div>
                ))}
            </div>
        ) : (
            <p>Loading...</p>
        )}
      
    </div>
  )
}
export default ReactUseEffect