import React from "react";
import { Link } from "react-router-dom";

const LogoutBtn = () => {
  const handleClick = () => {
    fetch('/api/logout')
    .then(response=>console.log(response.json()));
  }

  return (
    <Link to="/" className="btn btn-danger" onClick={()=>{handleClick()}}>Logout</Link>
  )
}

export default LogoutBtn