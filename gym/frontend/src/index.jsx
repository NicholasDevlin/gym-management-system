import React from 'react';
import ReactDOM from 'react-dom/client';
import './assets/style/index.css';
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import App from './App';
import Membership from './pages/user/membership/Index.jsx'
import Authentication from './pages/authentication/Index.jsx'
import reportWebVitals from './reportWebVitals';

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
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

reportWebVitals();
