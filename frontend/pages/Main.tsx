import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";


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
    <div>Main
    <p>User:<b> {auth? <span>{user["Email"]}</span>: <Link to="/login">Login</Link>}</b></p>
    </div>
  );
};

export default Main;