import React, { useState } from "react";
import Styles from "./Login.module.css";
import GoogleLoginButton from "./loginWithGoogle/LoginGoogle.jsx";
import { API_URLS } from "../../apiConfig.js";
import PasswordFied from "./inputPasswordField/PasswordField.jsx";
import TextField from "../general/input/inputTextField/TextField.jsx";
import { useNavigate } from "react-router-dom";

function Login({ registerOnClick }) {
  const navigate = useNavigate();
  const [loginData, setLoginData] = useState({
    email: "",
    password: "",
  });

  const handleInputChange = (e) => {
    const { id, value } = e.target;
    setLoginData((prevData) => ({
      ...prevData,
      [id]: value,
    }));
  };

  const handleLoginSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch(API_URLS.LOGIN, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(loginData),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const responseData = await response.json();
      if (responseData.success) {
        navigate("/");
      }
      localStorage.setItem("authToken", responseData.data.token);
      console.log("Login successful. Response:", responseData);
    } catch (error) {
      console.error("Error during login:", error);
    }
  };

  return (
    <div className={Styles.container} id="container">
      <form className={Styles.form} onSubmit={handleLoginSubmit}>
        <TextField id={"email"} label={"Email"} onChange={handleInputChange} />
        <PasswordFied id={"password"} onChange={handleInputChange} />
        <div className={Styles.row}>
          <button className={Styles.button} type="submit" id="submit">
            Sign in
          </button>
        </div>
        <p className={Styles.textColor}>or login with:</p>
        <div className={Styles.centerContent}>
          <GoogleLoginButton />
        </div>
        <div className={Styles.register}>
          <span className={Styles.textColor}>Don't have an account? </span>
          <button
            className={`${Styles.linkButton} ${Styles.textColor}`}
            type="button"
            onClick={registerOnClick}
          >
            Register new account
          </button>
        </div>
      </form>
    </div>
  );
}
// source https://codepen.io/hexagoncircle/pen/zYxzQqa
export default Login;
