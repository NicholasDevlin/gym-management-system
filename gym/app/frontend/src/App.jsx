import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { useUserRole } from './utils/jwt/UserRole';
import Home from './pages/superAdmin/home/Index.jsx';
import Membership from './pages/user/membership/Index.jsx';
import MembershipForAdmin from './pages/superAdmin/membership/Index.jsx';
import Authentication from './pages/authentication/Index.jsx';
import Profile from './pages/profile/Index.jsx';
import Help from './pages/user/help/Index.jsx';
import MembershipEditor from './pages/superAdmin/membership/Editor.jsx';

const App = () => {
  const { userRole } = useUserRole();

  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/authentication" element={<Authentication />} />
        <Route
          path="/membership"
          element={userRole === 'admin' ? <MembershipForAdmin /> : <Membership />}
        />
        <Route path="/profile" element={<Profile />} />
        <Route path="/membership/editor" element={<MembershipEditor />} />
        <Route path="/help" element={<Help />} />
      </Routes>
    </Router>
  );
};

export default App;
