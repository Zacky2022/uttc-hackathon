import { useEffect, useState } from "react";

type Props = {
    onSubmit: (name:string, age:number) => void;
};

const Form = (props: Props) => {
  const [name, setName] = useState("");
  const [age, setAge] = useState(0);
  
  const submit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    props.onSubmit(name, age)
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
      <label>Age: </label>
      <input
        type={"number"}
        style={{ marginBottom: 20 }}
        value={age}
        onChange={(e) => setAge(Number(e.target.value))}
      ></input>
      </div>
      <button type={"submit"}>POST</button>
    </form>
  );
};

export default Form;