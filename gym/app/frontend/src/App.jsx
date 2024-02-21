import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/superAdmin/home/Index.jsx';
import Membership from './pages/user/membership/Index.jsx';
import MembershipForAdmin from './pages/superAdmin/membership/Index.jsx';
import Authentication from './pages/authentication/Index.jsx';
import { useUserRole } from './utils/jwt/UserRole';

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
      </Routes>
    </Router>
  );
};

export default App;
