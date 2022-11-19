import  {useState} from "react";
import { stringify } from "querystring";

type Props = {
  onSubmit: (name:string, setName:React.Dispatch<React.SetStateAction<string>>) => void;
};

const Registerform = (props: Props) => {
  const [name, setName] = useState("");
  
  const submit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    props.onSubmit(name,setName)
  }
  
  return (
    <form  onSubmit={submit}>
      <div style={{display: "flex", justifyContent: "center"}} >
      <label>Name: </label>
      <input
        type={"text"}
        style = {{ marginBottom: 15 }}
        value={name}
        onChange={(e) => {
          setName(e.target.value)
        }}
      ></input>
      </div>
      <button type={"submit"}>CREATE</button>
    </form>
  );
};

export default Registerform;