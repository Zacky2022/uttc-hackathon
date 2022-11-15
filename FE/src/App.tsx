import "./App.css";
import { useEffect, useState } from "react";
import Form from "./Form";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./Login";
import Register from "./Register";
import NotFound from "./NotFound";
import Mainpage from "./Mainpage";
import PointList from "./Pointlist";

// type typePOST = {
//   name: string
// }

// function App() {
//   const [name, setName] = useState("");
//   const [users, setUsers] = useState<typePOST[]>([])

//   useEffect(() => {
//       fetch('http://localhost:8080/user', {method: 'GET'})
//       .then((res) => res.json())
//       .then((data) => {
//         setUsers(data)
//       })
//   },[])

//   const onSubmit = async (name: string) => {
//     if (!name) {
//       alert("Please enter name");
//       return;
//     }

//     if (name.length > 50) {
//       alert("Please enter a name shorter than 50 characters");
//       return;
//     }
//     try {
//       const result = await fetch("http://localhost:8080/user", {
//         method: "POST",
//         body: JSON.stringify({
//           name: name
//         }),
//       });
//       if (!result.ok) {
//         throw Error(`Failed to create user: ${result.status}`);
//       }
//       setName("");
//       fetch('http://localhost:8080/user', {method: 'GET'})
//       .then((res) => res.json())
//       .then((data) => {
//         setUsers(data)
//       })
//     } catch (err) {
//       console.error(err);
//     }
//   };
   
//   return (
//     <div className="App">
//       <header className="App-header">
//         User Register
//       </header>
//       <main className="App-body">
//         <Form onSubmit={onSubmit} />
//         <div className="datacontainer">
//             <ul>
//                 {
//                     users.map((post) => 
//                     <p className="DBdata" key={post.name}>{post.name}</p>
//                     )
//                 }
//             </ul>
//         </div>
//         </main>
//     </div>
//   );
// }

// export default App;

const App = () => {
  const [id, setId] = useState<string|undefined>("");
  return (
    <BrowserRouter>
      <Routes>
        <Route path={`/`} element={<Login id={id} setId={setId}/>} />
        <Route path={`/register/`} element={<Register />} />
        <Route path={`/mainpage/`} element={<Mainpage id={id} />} />
        <Route path={`/pointlist/`} element={<PointList/>} />
        <Route path={`/*/`} element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;