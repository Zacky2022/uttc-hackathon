import "./App.css";
import { useState } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./Login";
import Register from "./Register";
import NotFound from "./NotFound";
import Mainpage from "./Mainpage";
import PointList from "./Pointlist";

const App = () => {
  const [id, setId] = useState("");
  const [name, setName] = useState("");

  return (
    <BrowserRouter>
      <Routes>
        <Route path={`/`} element={<Login id={id} setId={setId} name={name} setName={setName}/>} />
        <Route path={`/register/`} element={<Register />} />
        <Route path={`/mainpage/`} element={<Mainpage id={id} name={name} />} />
        <Route path={`/pointlist/`} element={<PointList/>} />
        <Route path={`/*/`} element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;