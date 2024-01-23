import React, { useEffect } from 'react';
import Styles from './Login.module.css'
import GoogleLoginButton from './LoginGoogle.jsx';

function Login() {
  const handleGoogleLoginFailure = (error) => {
    console.error('Google login failed:', error);
  };
  useEffect(() => {
    const root = document.getElementById('container');
    const eye = document.getElementById('eyeball');
    const beam = document.getElementById('beam');
    const passwordInput = document.getElementById('password');

    const handleMouseMove = (e) => {
      let rect = beam.getBoundingClientRect();
      let mouseX = rect.right + rect.width / 2;
      let mouseY = rect.top + rect.height / 2;
      let rad = Math.atan2(mouseX - e.pageX, mouseY - e.pageY);
      let degrees = (rad * (20 / Math.PI) * -1) - 350;

      root.style.setProperty('--beamDegrees', `${degrees}deg`);
    };

    const handleEyeClick = (e) => {
      e.preventDefault();
      root.classList.toggle(Styles.showPassword);
      passwordInput.type =
        passwordInput.type === 'password' ? 'text' : 'password';
      passwordInput.focus();
    };

    root.addEventListener('mousemove', handleMouseMove);
    eye.addEventListener('click', handleEyeClick);

    return () => {
      // Cleanup event listeners on component unmount
      root.removeEventListener('mousemove', handleMouseMove);
      eye.removeEventListener('click', handleEyeClick);
    };
  }, []);

  return (
    <div className={`${Styles.container}`} id="container">
      <form onSubmit={() => false}>
        <div className={Styles.formItem}>
          <label htmlFor="email">Email</label>
          <div className={Styles.inputWrapper}>
            <input
              type="text"
              id="email"
              autoComplete="off"
              autoCorrect="off"
              autoCapitalize="off"
              spellCheck="false"
              data-lpignore="true"
            />
          </div>
        </div>
        <div className={Styles.formItem}>
          <label htmlFor="password">Password</label>
          <div className={Styles.inputWrapper}>
            <input
              type="password"
              id="password"
              autoComplete="off"
              autoCorrect="off"
              autoCapitalize="off"
              spellCheck="false"
              data-lpignore="true"
              className={Styles.inputPassword}
            />
            <button type="button" id="eyeball">
              <div className={Styles.eye}></div>
            </button>
            <div id="beam"></div>
          </div>
        </div>
        <div className={Styles.row}>
          <GoogleLoginButton className={Styles.button} onFailure={handleGoogleLoginFailure} />
          <button className={Styles.button} type="submit" id="submit">
            Sign in
          </button>
        </div>
      </form>
    </div>
  );
};
// source https://codepen.io/hexagoncircle/pen/zYxzQqa
export default Login;
