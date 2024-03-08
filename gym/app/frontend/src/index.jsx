import React from 'react';
import ReactDOM from 'react-dom';
import './assets/style/index.css';
import { GoogleOAuthProvider } from '@react-oauth/google';
import { UserDataProvider } from './utils/jwt/UserData.jsx';
import { transitions, positions, Provider as AlertProvider } from 'react-alert'
import AlertTemplate from 'react-alert-template-basic'
import App from './App.jsx';

const clientId = '429032937526-3um12rsfk8vh2dual7klnf69i05caoi7.apps.googleusercontent.com';
// alert configuration
const options = {
  position: positions.TOP_RIGHT,
  timeout: 5000,
  offset: '30px',
  transition: transitions.FADE,
  containerStyle: {
    color: 'white'
  },
}

ReactDOM.render(
  <React.StrictMode>
    <GoogleOAuthProvider clientId={clientId}>
      <UserDataProvider>
        <AlertProvider template={AlertTemplate} {...options}>
          <App /> {/* Render the AppRouter component */}
        </AlertProvider>
      </UserDataProvider>
    </GoogleOAuthProvider>
  </React.StrictMode>,
  document.getElementById('root')
);