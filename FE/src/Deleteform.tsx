import  { useState } from "react";
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

const Deleteform = (props: Props) => {
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
    <form  onSubmit={submit} className="DeleteForm">
      <h2>Deleting contribution</h2>
      <Select 
          options={options}
          defaultValue={{label:'Contribution to delete', value:'default'}}
          onChange={(value)=>{
            setTarg(String(value?.value));
          }}
          />
      <button type={"submit"}>DELETE</button>
    </form>
  );
};

export default Deleteform;