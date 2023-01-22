import React from 'react';
import ReactDOM from 'react-dom/client';
import {BrowserRouter, Routes, Route} from "react-router-dom";

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap";

import Index from "./pages/Index";
import Login from './pages/Login';
import Signup from './pages/Signup';
import Main from './pages/Main';

function Application() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' index element={<Index />} />
        <Route path='/login' element={<Login />} />
        <Route path='/signup' element={<Signup />} />
        <Route path="/main" element={<Main />}/>
      </Routes>
    </BrowserRouter>
  )
}

const root = ReactDOM.createRoot(document.querySelector('#application')!);
root.render(<Application />);