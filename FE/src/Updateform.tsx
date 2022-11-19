import  { useState,useEffect } from "react";
import { stringify } from "querystring";
import Select from "react-select";

type Consftype = {
  msid: string
  sentpoint:number
  message:string
  name:string
}

type Props = {
  onSubmit: (targ:string, point:number, message:string, setPoint:React.Dispatch<React.SetStateAction<number>>,setMessage:React.Dispatch<React.SetStateAction<string>>) => void;
  consf: Consftype[]
};

type typeCons = {
  msid: string
  sentpoint: number
  message: string
  name: string
}

const Messageform = (props: Props) => {
  const [targ,setTarg] = useState("");
  const [point, setPoint] = useState<number>(0);
  const [message, setMessage] = useState("");
  
  const submit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    props.onSubmit(targ, point, message, setPoint,setMessage)
  }

  let options = []
  for (let i=0; i<props.consf.length; i++) {
    options[i] = {value:props.consf[i]['msid'], label:props.consf[i]['message']}
  }
  
  return (
    <form  onSubmit={submit}>
      <h2>Updating contribution</h2>
      <Select 
          options={options}
          defaultValue={{label:'Contribution to modify', value:'default'}}
          onChange={(value)=>{
            setTarg(String(value?.value));
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
      <button type={"submit"}>UPDATE</button>
    </form>
  );
};

export default Messageform;