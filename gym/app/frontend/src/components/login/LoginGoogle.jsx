import React from 'react';
import Styles from './Login.module.css'
import { GoogleOAuthProvider } from '@react-oauth/google';
// import { GoogleLogin } from 'react-google-login';

const GoogleLoginButton = ({ onSuccess, onFailure }) => {
  const clientId = '429032937526-3um12rsfk8vh2dual7klnf69i05caoi7.apps.googleusercontent.com';

  const responseGoogle = (response) => {
    if (response && response.profileObj) {
      const { googleId, email, name, imageUrl } = response.profileObj;
      console.log({ googleId, email, name, imageUrl });
    }
  };

  return (
    <GoogleOAuthProvider clientId={clientId}>...</GoogleOAuthProvider>
    // <GoogleLogin
    //   className={Styles.button}
    //   clientId={clientId}
    //   buttonText="Login with Google"
    //   onSuccess={responseGoogle}
    //   onFailure={onFailure}
    //   cookiePolicy={'single_host_origin'}
    // />
  );
};

export default GoogleLoginButton;
