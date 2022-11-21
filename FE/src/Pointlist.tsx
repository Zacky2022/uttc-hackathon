import { Link } from "react-router-dom";
import { useState, useEffect } from "react";

const BELink = "https://hackathon-2-sk7fvtjuea-uc.a.run.app";

type typeUser = {
  id: string
  name: string
  point: number
}

function List() {
  const [users, setUsers] = useState<typeUser[]>([])
  useEffect(() => {
      fetch(`${BELink}/user`, {method: 'GET'})
      .then((res) => res.json())
      .then((data) => {
        setUsers(data)
      })
  },[])
     
  return (
    <div className="App">
      <header className="App-header">Point List</header>
      <main className="App-body">
        <div className="datacontainer">
            <ul>
                {
                    users.map((post) => 
                    <p className="DBdata" key={post.name}>{post.name}:{post.point}</p>
                    )
                }
            </ul>
        </div>
        <div>
//        <Link to={`/`}>Go back to Login Form</Link>
//      </div>
        </main>
    </div>
  );
}

export default List;