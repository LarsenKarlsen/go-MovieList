import React from "react";
import { Link } from "react-router-dom"

const Index = () => {

  return (
    <ul>
      <li>
        <Link to="/login">Login</Link>
      </li>
    </ul>
  );
};

export default Index;