import React from 'react';
// import Styles from './Login.module.css'
// import GoogleLoginButton from './loginWithGoogle/LoginGoogle.jsx';
// import { API_URLS } from '../../apiConfig.js';
// import PasswordFied from './inputPasswordField/PasswordField.jsx';
// import TextField from './inputTextField/TextField.jsx';
// import { useNavigate } from 'react-router-dom';

function Register({loginOnClick}) {
  return (
    <>
    <h1>Register</h1>
    <button onClick={loginOnClick}>Register</button>
    </>
  )
}

export default Register;