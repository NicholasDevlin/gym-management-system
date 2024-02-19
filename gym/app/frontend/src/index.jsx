import React from 'react';
import ReactDOM from 'react-dom/client';
import './assets/style/index.css';
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import App from './App.jsx';
import Membership from './pages/user/membership/Index.jsx'
import Authentication from './pages/authentication/Index.jsx'
import reportWebVitals from './reportWebVitals.js';
import { GoogleOAuthProvider } from '@react-oauth/google';

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />
  },
  {
    path: "membership",
    element: <Membership />
  },
  {
    path: "authentication",
    element: <Authentication />
  }
])
const clientId = '429032937526-3um12rsfk8vh2dual7klnf69i05caoi7.apps.googleusercontent.com';
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <GoogleOAuthProvider clientId={clientId}>
      <RouterProvider router={router} />
    </GoogleOAuthProvider>
  </React.StrictMode>
);

reportWebVitals();
