import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/superAdmin/home/Index.jsx';
import Membership from './pages/user/membership/Index.jsx';
import MembershipForAdmin from './pages/superAdmin/membership/Index.jsx';
import Transaction from './pages/user/transaction/Index.jsx';
import TransactionForAdmin from './pages/superAdmin/transaction/Index.jsx';
import Authentication from './pages/authentication/Index.jsx';
import Profile from './pages/profile/Index.jsx';
import Help from './pages/user/help/Index.jsx';
import MembershipEditor from './pages/superAdmin/membership/Editor.jsx';
import { useUserData } from './utils/jwt/UserData.jsx';

function App() {
  const { userData } = useUserData();

  return (
    <Router>
      <div>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/authentication" element={<Authentication />} />
          <Route path="/membership" element={userData && userData.role === 'admin' ? <MembershipForAdmin /> : <Membership />} />
          <Route path="/transaction" element={userData && userData.role === 'admin' ? <TransactionForAdmin /> : <Transaction />} />
          <Route path="/profile" element={<Profile />} />
          <Route path="/membership/editor/:uuid" element={<MembershipEditor />} />
          <Route path="/membership/editor" element={<MembershipEditor />} />
          <Route path="/help" element={<Help />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;
