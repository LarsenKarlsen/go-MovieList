import React, { useState } from "react";

const AccSignForm = () => {
    const formSignin = 
        {
            width: "100%",
            maxWidth: "330px",
            padding: "15px",
            margin: "auto",
        } as React.CSSProperties;

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const submitFormHandler = async (e) => {
        e.preventDefault()
        const formData = {
            email: email,
            password: password
        }

        await fetch("/api/signup", {
            method: "POST",
            headers: {
                "content-type": "application/json" 
            },
            body: JSON.stringify(formData)
        })
        .then(response => response.json())
        .then(response => console.log(response))
    }


    return (
    <form style={formSignin}>
    <h1 className="h3 mb-3 fw-normal">Please sign in</h1>

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
    Sign in</button>
    <p className="mt-5 mb-3 text-muted">&copy; 2017–2022</p>
  </form>);
}

export default AccSignForm;