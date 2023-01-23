import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";

import Navbar from "../components/Navbar";

const Main = () => {

  const [auth, setAuth] = useState(false);
  const [user, setUser] = useState({});

  useEffect(() => {
    fetch("api/me")
    .then(response=>{
      if (response.status === 200) {
        setAuth(true);
      }
      return response
    })
    .then(response=>response.json())
    .then(response=>setUser(response["msg"]))
  },[])

  return (
    <div>
    <Navbar auth={auth} userEmail={user["Email"]}/>
    <main role="main">
      <div className="container">
        <h1>Navbar example</h1>
        <p className="lead">This example is a quick exercise to illustrate how fixed to top navbar works. As you scroll, it will remain fixed to the top of your browser's viewport.</p>
        <a className="btn btn-lg btn-primary" href="../../components/navbar/" role="button">View navbar docs &raquo;</a>
      </div>
    </main>
    </div>
  );
};

export default Main;