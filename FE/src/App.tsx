import "./App.css";
import { useEffect, useState } from "react";
import Form from "./Form";

type typePOST = {
  name: string
  age : number
}

function App() {
  const [name, setName] = useState("");
  const [age, setAge] = useState(0);
  const [users, setUsers] = useState<typePOST[]>([])

  useEffect(() => {
      fetch('http://localhost:8080/user', {method: 'GET'})
      .then((res) => res.json())
      .then((data) => {
        setUsers(data)
      })
  },[])

  const onSubmit = async (name: string, age: number) => {
    if (!name) {
      alert("Please enter name");
      return;
    }

    if (name.length > 50) {
      alert("Please enter a name shorter than 50 characters");
      return;
    }

    if (age < 20 || age > 80) {
      alert("Please enter age between 20 and 80");
      return;
    }

    try {
      const result = await fetch("http://localhost:8080/user", {
        method: "POST",
        body: JSON.stringify({
          name: name,
          age: age,
        }),
      });
      if (!result.ok) {
        throw Error(`Failed to create user: ${result.status}`);
      }
      setName("");
      setAge(0);
      fetch('http://localhost:8080/user', {method: 'GET'})
      .then((res) => res.json())
      .then((data) => {
        setUsers(data)
      })
    } catch (err) {
      console.error(err);
    }
  };
   
  return (
    <div className="App">
      <header className="App-header">
        User Register
      </header>
      <main className="App-body">
        <Form onSubmit={onSubmit} />
        <div className="datacontainer">
            <ul>
                {
                    users.map((post) => 
                    <p className="DBdata" key={post.name}>{post.name}:{post.age}</p>
                    )
                }
            </ul>
        </div>
        </main>
    </div>
  );
}

export default App;
