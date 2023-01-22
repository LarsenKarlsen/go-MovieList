import React from "react";

const Navbar = (props) => {
  return (
    <nav className="navbar navbar-dark bg-primary">
      <div className="container">
        <a className="navbar-brand" href="/main">MovieList</a>
        {props.auth
        ? <button className="btn btn-danger">Logout</button>
        : <button className="btn btn-success">Login</button>
        }
      </div>
    </nav>
  )
}

export default Navbar