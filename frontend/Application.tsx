import React from 'react';
import ReactDOM from 'react-dom/client';
import {BrowserRouter, Routes, Route} from "react-router-dom";

import Index from "./pages/Index";
import Login from './pages/Login';
import Signup from './pages/Signup';

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap";

function Application() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' index element={<Index />} />
        <Route path='/login' element={<Login />} />
        <Route path='/signup' element={<Signup />} />
      </Routes>
    </BrowserRouter>
  )
}

const root = ReactDOM.createRoot(document.querySelector('#application')!);
root.render(<Application />);