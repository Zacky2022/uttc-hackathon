import "./App.css";
import { Link, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import Form from "./Form";
import Select from "react-select"
import { convertTypeAcquisitionFromJson, idText } from "typescript";
import { stringify } from "querystring";

// const Home = () => {
//   return (
//     <>
//       <h1>ホーム</h1>
//       <div>
//         新規登録は<Link to={`/register/`}>こちら</Link>
//       </div>
//     </>
//   );
// };

// export {};
// export default Home;
type typeUsers = {
  id: string
  name: string
}

type Props = {
  id:string|undefined
  setId: React.Dispatch<React.SetStateAction<string|undefined>>
};

const Login = (props:Props) => {
  const [name, setName] = useState("");
  const [users, setUsers] = useState<typeUsers[]>([])
  useEffect(() => {
      fetch('http://localhost:8080/user', {method: 'GET'})
      .then((res) => res.json())
      .then((data) => {
        setUsers(data)
      })
  },[])
  let options = []
  for (let i=0; i<users.length; i++) {
    options[i] = {value:users[i]['id'], label:users[i]['name']}
  }

  
  // const onSubmit = async (id: string, name: string) => {
  //   if (!name) {
  //     alert("Please enter name");
  //     return;
  //   }

  //   if (name.length > 50) {
  //     alert("Please enter a name shorter than 50 characters");
  //     return;
  //   }
  //   if (users.includes(name) != true) {
  //     alert("Sorry, couldn't find such user name")
  //   }
  //   try {
  //     const result = await fetch("http://localhost:8080/user", {
  //       method: "GET",
  //       body: JSON.stringify({
  //         id : id,
  //         name : name,
  //       }),
  //     });
  //     if (!result.ok) {
  //       throw Error(`Failed to create user: ${result.status}`);
  //     }
  //     setName("");
  //     fetch('http://localhost:8080/user', {method: 'GET'})
  //     .then((res) => res.json())
  //     .then((data) => {
  //       setUsers(data)
  //     })
  //   } catch (err) {
  //     console.error(err);
  //   }
  // };


// const onChange = (value:string) => {
//   useMainpage;
//   console.log(value);
// };

  const navigate = useNavigate();
  return (
    <div className="App">
      <header className="App-header">
        User Login
      </header>
      <main className="App-body">
        <Select 
        options={options}
        defaultValue={{label:'LoginName', value:'default'}}
        // onChange={(value)=>{
        //   props.setId(value['value'])
        // }{tomainpage}}
        onChange={(value)=>{
          navigate(`/mainpage/`);
          props.setId(value?.value);
        }}
        />
        <div className="datacontainer">
        </div>
        <div>
        <Link to={`/register/`}>Create New Account</Link>
        </div>
        <div>
        <Link to={`/pointlist/`}>View Point List</Link>
        </div>
        </main>
    </div>
  );
}

export {};
export default Login;