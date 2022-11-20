import { Link } from "react-router-dom";
import { useEffect, useState } from "react";
import Messageform from "./Messageform";
import Updateform from "./Updateform";
import Deleteform from "./Deleteform";
import Sidebar from "./components/Sidebar"

const BELink = "https://hackathon-2-sk7fvtjuea-uc.a.run.app";

type Props = {
  id:string
  name: string
}

type Consftype = {
  msid: string
  sentpoint:number
  message:string
  name:string
}

type Constotype = {
  sentpoint:number
  message:string
  name:string
}

const Mainpage = (props:Props) => {
  const [consf, setConsf] = useState<Consftype[]>([]);
  const [consto, setConsto] = useState<Constotype[]>([]);
  const Id = props.id
  const Name = props.name

  const Getfunction = () => {
    useEffect(() => {
      fetch(`${BELink}/con-list?user_id=${Id}&ft=from`, {method: 'GET'})
      .then((res) => res.json())
      .then((data) => {
        setConsf(data)
      })
    },[])
    useEffect(() => {
      fetch(`${BELink}/con-list?user_id=${Id}&ft=to`, {method: 'GET'})
      .then((res) => res.json())
      .then((data) => {
        setConsto(data)
      })
    },[])
  };
  Getfunction();
  
  const onSubmit = async (to:string|undefined, point:number, message:string, setPoint:React.Dispatch<React.SetStateAction<number>>,setMessage:React.Dispatch<React.SetStateAction<string>>) => {
        if (!to) {
          alert("Please enter name");
          return;
        }
        if (to.length > 50) {
          alert("Please enter a name shorter than 50 characters");
          return;
        }
        if (point<=0) {
          alert("You cannot send point less than 1");
          return;
        }
        try {
          const result = await fetch(`${BELink}/con-list?user_id=${Id}`, {
            method: "POST",
            body: JSON.stringify({
              from:Id,
              to: to,
              point: point,
              message: message,
            }),
          });
          if (!result.ok) {
            throw Error(`Failed to create user: ${result.status}`);
          }
          setPoint(0);
          setMessage("");
            fetch(`${BELink}/con-list?user_id=${Id}&ft=from`, {method: 'GET'})
            .then((res) => res.json())
            .then((data) => {
              setConsf(data)
            })
            fetch(`${BELink}/con-list?user_id=${Id}&ft=to`, {method: 'GET'})
            .then((res) => res.json())
            .then((data) => {
              setConsto(data)
            })
          } catch (err) {
          console.error(err);
        }
      };

  const onSubmit2 = async (targ:string, point:number, message:string, setPoint:React.Dispatch<React.SetStateAction<number>>,setMessage:React.Dispatch<React.SetStateAction<string>>) => {
    if (!targ) {
      alert("Please choose one message to update");
      return;
    }
    try {
      const result = await fetch(`${BELink}/update`, {
        method: "POST",
          body: JSON.stringify({
          targ:targ,
          point: point,
          message: message,
        }),
      });
      if (!result.ok) {
        throw Error(`Failed to create user: ${result.status}`);
      }
      setPoint(0);
      setMessage("");
        fetch(`${BELink}/con-list?user_id=${Id}&ft=from`, {method: 'GET'})
        .then((res) => res.json())
        .then((data) => {
          setConsf(data)
        })
        fetch(`${BELink}/con-list?user_id=${Id}&ft=to`, {method: 'GET'})
        .then((res) => res.json())
        .then((data) => {
          setConsto(data)
        })
      } catch (err) {
      console.error(err);
    }
  };

  const onSubmit3 = async (targ:string, point:number, message:string, setPoint:React.Dispatch<React.SetStateAction<number>>,setMessage:React.Dispatch<React.SetStateAction<string>>) => {
    if (!targ) {
      alert("Please choose one message to delete");
      return;
    }
    try {
      const result = await fetch(`${BELink}/delete?targ_id=${targ}`, {method: 'GET'})
      if (!result.ok) {
        throw Error(`Failed to create user: ${result.status}`);
      }
      setPoint(0);
      setMessage("");
        fetch(`${BELink}/con-list?user_id=${Id}&ft=from`, {method: 'GET'})
        .then((res) => res.json())
        .then((data) => {
          setConsf(data)
        })
        fetch(`${BELink}/con-list?user_id=${Id}&ft=to`, {method: 'GET'})
        .then((res) => res.json())
        .then((data) => {
          setConsto(data)
        })
      } catch (err) {
      console.error(err);
    }
  };

  return (
    <>
      <h1>Hello, {Name} san !</h1>
      <div className="Sidebar">
        <Sidebar/>
      </div>
      <h2>contributions you sent</h2>
      <ul>
                {
                    consf.map((post) => 
                    <p className="DBdata" key={post.message}>You sent {post.sentpoint} points to {post.name}. Message:{post.message} </p>
                    )
                }
      </ul>
      <h2>contributions you got</h2>
      <ul>
                {
                    consto.map((post) => 
                    <p className="DBdata" key={post.message}>You got {post.sentpoint} points from {post.name}. Message:{post.message} </p>
                    )
                }
      </ul>
      <Messageform onSubmit={onSubmit}/>
      <Updateform onSubmit={onSubmit2} consf={consf}/>
      <Deleteform onSubmit={onSubmit3} consf={consf}/>
      <div>
        <Link to={`/`}>Go back to Lorgin Form</Link>
      </div>
      <div>
        <Link to={`/pointlist/`}>Show Point List</Link>
      </div>
    </>
  );
};

export default Mainpage;