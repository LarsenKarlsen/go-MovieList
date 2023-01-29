import React, {useState} from "react";
import {Navigate} from "react-router-dom";

const AccLogForm = () =>{
  const formSignin = 
      {
          width: "100%",
          maxWidth: "330px",
          padding: "15px",
          margin: "auto",
      } as React.CSSProperties;

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [redirectToMain, setRedirectToMain] = useState(false);

  const submitFormHandler = async (e) => {
      e.preventDefault()
      const formData = {
          email: email,
          password: password
      }

      await fetch("/api/user/login", {
          method: "POST",
          headers: {
              "content-type": "application/json" 
          },
          body: JSON.stringify(formData)
      })
      .then(async response => {
        if (response.status===200){
            setRedirectToMain(true);
        }
    })
  }

  if (redirectToMain){
    return <Navigate to="/main"/>
  }


  return (
  <form style={formSignin}>
  <h1 className="h3 mb-3 fw-normal">Please login</h1>

  <div className="form-control">
    <label>
      Email address
      <input 
      type="email"
      className="form-control"
      name="email"
      placeholder="name@example.com"
      onChange={(e)=>{setEmail(e.target.value)}}
      ></input>
      </label>
  </div>
  <div className="form-control">
    <label>
      Password
      <input
      type="password" 
      className="form-control"
      name="password"
      placeholder="Password"
      onChange={(e)=>{setPassword(e.target.value)}}>
      </input>
      </label>
  </div>

  <button
  className="w-100 btn btn-lg btn-primary"
  onClick={submitFormHandler}>
  Login</button>
  <p className="mt-5 mb-3 text-muted">&copy; 2017–2022</p>
</form>);
}

export default AccLogForm;