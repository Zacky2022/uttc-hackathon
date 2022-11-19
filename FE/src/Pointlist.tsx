import { Link } from "react-router-dom";
import { useState, useEffect } from "react";

type typeUser = {
  id: string
  name: string
  point: number
}

function List() {
  const [users, setUsers] = useState<typeUser[]>([])
  useEffect(() => {
      fetch('https://hackathon-2-sk7fvtjuea-uc.a.run.app:8080/user', {method: 'GET'})
      .then((res) => res.json())
      .then((data) => {
        setUsers(data)
      })
  },[])
     
  return (
    <div className="App">
      <header className="App-header">
        Point List
      </header>
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
//   return (
//     <>
//       <h1>ポイント一覧</h1>
//       <div>
//         <Link to={`/`}>Login画面に戻る</Link>
//       </div>
//     </>
//   );
// };

export default List;