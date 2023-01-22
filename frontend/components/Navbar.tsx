import React from "react";
import LoginBtn from "./LoginBtn";
import LogoutBtn from "./LogoutBtn";

const Navbar = (props) => {
  return (
    <nav className="navbar navbar-dark bg-primary">
      <div className="container">
        <a className="navbar-brand" href="/main">MovieList</a>
        {props.auth
        ? <LogoutBtn />
        : <LoginBtn />
        }
      </div>
    </nav>
  )
}

export default Navbar