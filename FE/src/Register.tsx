import { Link, useNavigate } from "react-router-dom";
import Registerbox from "./Registerform";
import { useEffect, useState } from "react";

function Register() {
  const navigate = useNavigate();
  const onSubmit = async (name: string, setName:React.Dispatch<React.SetStateAction<string>>) => {
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
          name: name,
        }),
      });
    if (!result.ok) {
      throw Error(`Failed to create user: ${result.status}`);
    }
    setName("");
    navigate(`/`)
  } catch (err) {
      console.error(err);
  }
};
     
    return (
      <div className="App">
        <header className="App-header">
          Create New Account
        </header>
        <main className="App-body">
          <Registerbox onSubmit={onSubmit} />
          <div className="datacontainer">
          </div>
          <div>
//        <Link to={`/`}>Go back to Login Form</Link>
//      </div>
        </main>
      </div>
    );
  }
  
  export default Register;
  