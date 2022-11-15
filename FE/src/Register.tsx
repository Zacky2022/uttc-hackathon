import { Link } from "react-router-dom";
import Form from "./Form";
import { useEffect, useState } from "react";

// const Register = () => {
//   return (
//     <>
//       <h1>新規登録ページ</h1>
//       <div>
//         ログインは<Link to={`/login/`}>こちら</Link>
//       </div>
//       <div>
//         <Link to={`/`}>ホームに戻る</Link>
//       </div>
//     </>
//   );
// };

// export default Register;
type typePOST = {
  name: string
}


function Register() {
    const [name, setName] = useState("");
    const [users, setUsers] = useState<typePOST[]>([])
  
    useEffect(() => {
        fetch('http://localhost:8080/user', {method: 'GET'})
        .then((res) => res.json())
        .then((data) => {
          setUsers(data)
        })
    },[])
  
    const onSubmit = async (name: string) => {
      if (!name) {
        alert("Please enter name");
        return;
      }
  
      if (name.length > 50) {
        alert("Please enter a name shorter than 50 characters");
        return;
      }
      try {
        const result = await fetch("http://localhost:8080/user", {
          method: "POST",
          body: JSON.stringify({
            name: name
          }),
        });
        if (!result.ok) {
          throw Error(`Failed to create user: ${result.status}`);
        }
        setName("");
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
          User Registers
        </header>
        <main className="App-body">
          {/* <Form onSubmit={onSubmit} /> */}
          <div className="datacontainer">
              <ul>
                  {
                      users.map((post) => 
                      <p className="DBdata" key={post.name}>{post.name}</p>
                      )
                  }
              </ul>
          </div>
          </main>
      </div>
    );
  }
  
  export default Register;