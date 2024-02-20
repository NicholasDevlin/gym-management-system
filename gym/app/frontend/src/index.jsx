import React from 'react';
import ReactDOM from 'react-dom/client';
import './assets/style/index.css';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import App from './App.jsx';
import Membership from './pages/user/membership/Index.jsx'
import MembershipForAdmin from './pages/superAdmin/membership/Index.jsx'
import Authentication from './pages/authentication/Index.jsx'
import reportWebVitals from './reportWebVitals.js';
import { GoogleOAuthProvider } from '@react-oauth/google';
import { UserRoleProvider, useUserRole } from './utils/jwt/UserRole.jsx';

const AppRouter = () => {
  const { userRole } = useUserRole();

  return (
    <Router>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/authentication" element={<Authentication />} />
        <Route path="/membership" element={userRole === 'admin' ? <MembershipForAdmin /> : <Membership />} />
        <Route path="*" element={<Navigate to="/" />} /> {/* Fallback route */}
      </Routes>
    </Router>
  );
};

const clientId = '429032937526-3um12rsfk8vh2dual7klnf69i05caoi7.apps.googleusercontent.com';
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <GoogleOAuthProvider clientId={clientId}>
      <UserRoleProvider>
        <AppRouter />
      </UserRoleProvider>
    </GoogleOAuthProvider>
  </React.StrictMode>
);

reportWebVitals();
