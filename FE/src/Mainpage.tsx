import { Link } from "react-router-dom";
import { useEffect, useState } from "react";

type Props = {
  id:string|undefined
}

const Mainpage = (props:Props) => {
  const [users, setUsers] = useState("");

  useEffect(() => {
    fetch('http://localhost:8080/user', {method: 'GET'})
    .then((res) => res.json())
    .then((data) => {
      setUsers(data)
    })
},[])
  return (
    <>
      <h1>{props.id}</h1>
      <div>
        <Link to={`/`}>ホームに戻る</Link>
      </div>
    </>
  );
};

export default Mainpage;