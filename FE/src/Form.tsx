import { useEffect, useState } from "react";
import Select from 'react-select';
import { stringify } from "querystring";

type Props = {
    onSubmit: (id: string, name:string) => void;
    options:string[]
};

const Form = (props: Props) => {
  const [name, setName] = useState("");
  const [id, setId] = useState("");
  const submit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    props.onSubmit(id, name)
  }

  return (
     <form  onSubmit={submit}>
      <div style={{display: "flex", justifyContent: "center"}} >
      <label>Name: </label>
      <input
        type={"text"}
        style = {{ marginBottom: 15 }}
        value={name}
        onChange={(e) => setName(e.target.value)}
      ></input>
      </div>
      <div style={{display: "flex", justifyContent: "center"}}>
      </div>
      {/* <button type={"submit"}>LOGIN</button> */}
      <Select options={props.options} />
    </form>
  );
};

export default Form;