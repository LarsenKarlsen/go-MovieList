import React from "react";
import { Link } from "react-router-dom";

const LogoutBtn = () => {
  const handleClick = () => {
    console.log("logout")
  }

  return (
    <Link to="/" className="btn btn-danger" onClick={()=>{handleClick()}}>Logout</Link>
  )
}

export default LogoutBtn