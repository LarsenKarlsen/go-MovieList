import React from "react";
import { Link } from "react-router-dom";

const LoginBtn = () => {
  return (
    <Link to="/login" className="btn btn-success">Login</Link>
  )
}

export default LoginBtn