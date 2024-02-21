import React from 'react';
import ReactDOM from 'react-dom';
import './assets/style/index.css';
import { GoogleOAuthProvider } from '@react-oauth/google';
import { UserRoleProvider } from './utils/jwt/UserRole.jsx';
import App from './App.jsx'; 

const clientId = '429032937526-3um12rsfk8vh2dual7klnf69i05caoi7.apps.googleusercontent.com';

ReactDOM.render(
  <React.StrictMode>
    <GoogleOAuthProvider clientId={clientId}>
      <UserRoleProvider>
        <App /> {/* Render the AppRouter component */}
      </UserRoleProvider>
    </GoogleOAuthProvider>
  </React.StrictMode>,
  document.getElementById('root')
);
