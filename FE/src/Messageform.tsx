import  { useState,useEffect } from "react";
import Select from "react-select";

type Props = {
  onSubmit: (to:string, point:number, message:string, setPoint:React.Dispatch<React.SetStateAction<number>>,setMessage:React.Dispatch<React.SetStateAction<string>>) => void;
};
type typeUsers = {
  id: string
  name: string
}

const Messageform = (props: Props) => {
  const [to,setTo] = useState("");
  const [point, setPoint] = useState<number>(0);
  const [message, setMessage] = useState("");
  
  const submit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    props.onSubmit(to, point, message, setPoint,setMessage)
  }

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
  
  return (
    <form  onSubmit={submit}>
      <h2>Sending contribution</h2>
      <Select 
          options={options}
          defaultValue={{label:'Send To', value:'default'}}
          onChange={(value)=>{
            setTo(String(value?.value));
          }}
          />
      <div style={{display: "flex", justifyContent: "center"}} >
        <label>Point: </label>
        <input
          type={"number"}
          style = {{ marginBottom: 15 }}
          value={point}
          onChange={(e) => {
            setPoint(Number(e.target.value))
          }}
        ></input>
        <label>Message: </label>
        <input
          type={"string"}
          style = {{ marginBottom: 15 }}
          value={message}
          onChange={(e) => {
            setMessage(String(e.target.value))
          }}
        ></input>
      </div>
      <button type={"submit"}>SEND</button>
    </form>
  );
};

export default Messageform;