import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { useUserData } from './utils/jwt/UserData.jsx';
import Home from './pages/superAdmin/home/Index.jsx';
import Membership from './pages/user/membership/Index.jsx';
import MembershipForAdmin from './pages/superAdmin/membership/Index.jsx';
import Authentication from './pages/authentication/Index.jsx';
import Profile from './pages/profile/Index.jsx';
import Help from './pages/user/help/Index.jsx';
import MembershipEditor from './pages/superAdmin/membership/Editor.jsx';

const App = () => {
  const userData = useUserData();

  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/authentication" element={<Authentication />} />
        {userData && (
          <Route
            path="/membership"
            element={userData.role === 'admin' ? <MembershipForAdmin /> : <Membership />}
          />
        )}
        <Route path="/profile" element={<Profile />} />
        <Route path="/membership/editor" element={<MembershipEditor />} />
        <Route path="/help" element={<Help />} />
      </Routes>
    </Router>
  );
};

export default App;
