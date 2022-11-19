import "./App.css";
import { Link, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import Select from "react-select"

type typeUsers = {
  id: string
  name: string
}

type Props = {
  id:string
  setId: React.Dispatch<React.SetStateAction<string>>
  name:string
  setName: React.Dispatch<React.SetStateAction<string>>
};

const Login = (props:Props) => {
  const [name, setName] = useState("");
  const [users, setUsers] = useState<typeUsers[]>([])
  useEffect(() => {
      fetch('https://hackathon-2-sk7fvtjuea-uc.a.run.app/user', {method: 'GET'})
      .then((res) => res.json())
      .then((data) => {
        setUsers(data)
      })
  },[])
  let options = []
  for (let i=0; i<users.length; i++) {
    options[i] = {value:users[i]['id'], label:users[i]['name']}
  }

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
        onChange={(value)=>{
          navigate(`/mainpage/`);
          props.setId(String(value?.value));
          props.setName(String(value?.label))
        }}
        />
        <div className="datacontainer">
        </div>
        <div>
        <Link to={`/register/`}>Create New Account</Link>
        </div>
        <div>
        <Link to={`/pointlist/`}>Show Point List</Link>
        </div>
        </main>
    </div>
  );
}

export {};
export default Login;