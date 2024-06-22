import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomeForAdmin from './pages/superAdmin/home/Index.jsx';
import HomeForMember from './pages/user/home/Index.jsx';
import Home from './pages/home/Index.jsx';
import Membership from './pages/user/membership/Index.jsx';
import MembershipForAdmin from './pages/superAdmin/membership/Index.jsx';
import Transaction from './pages/user/transaction/Index.jsx';
import TransactionForAdmin from './pages/superAdmin/transaction/Index.jsx';
import TransactionEditorForAdmin from './pages/superAdmin/transaction/Editor.jsx';
import Absensi from './pages/superAdmin/absensi/Index.jsx';
import User from './pages/superAdmin/user/Index.jsx';
import Authentication from './pages/authentication/Index.jsx';
import Profile from './pages/profile/Index.jsx';
import Help from './pages/user/help/Index.jsx';
import MembershipEditor from './pages/superAdmin/membership/Editor.jsx';
import { useUserData } from './utils/jwt/UserData.jsx';
import { ConfigProvider } from 'antd';

function App() {
  const { userData } = useUserData();

  return (
    <Router>
      <div>
        <ConfigProvider
          theme={{
            components: {
              Table: {
                headerBg: '#1f2124',
                headerColor: '#d8cdb9',
                colorBgContainer: '#43454e',
                colorText: '#d8cdb9',
                borderColor: '#6e6f73',
                fontWeightStrong: 600
              },
              Select: {
                selectorBg: '#1f2124',
                multipleItemBg: '#1f2124',
                optionSelectedBg: '#35373d'
              },
              Calendar: {
                colorFillSecondary: '#000',
                colorBgContainer: '#f0f0f0'
              }
            },
          }}
        >
          <Routes>
            {!userData ?
              <>
                <Route path="/" element={<Home />} />
                <Route path="/membership" element={<Membership />} />
              </>
              :
              <></>
            }
            <Route path="/authentication" element={<Authentication />} />
            {userData && userData.role === 'admin' ?
              <>
                <Route path="/" element={<HomeForAdmin />} />
                <Route path="/membership/editor/:uuid" element={<MembershipEditor />} />
                <Route path="/membership/editor" element={<MembershipEditor />} />
                <Route path="/transaction/editor" element={<TransactionEditorForAdmin />} />
                <Route path="/transaction/editor/:uuid" element={<TransactionEditorForAdmin />} />
                <Route path="/absensi" element={<Absensi />} />
                <Route path="/user" element={<User />} />
                <Route path="/membership" element={<MembershipForAdmin />} />
              </>
              :
              <></>
            }
            {userData && userData.role === 'user' ?
              <>
                <Route path="/" element={<HomeForMember />} />
                <Route path="/membership" element={<Membership />} />
              </>
              :
              <>
              </>
            }
            <Route path="/transaction" element={userData && userData.role === 'admin' ? <TransactionForAdmin /> : <Transaction />} />
            <Route path="/profile" element={<Profile />} />
            <Route path="/help" element={<Help />} />
          </Routes>
        </ConfigProvider>
      </div>
    </Router>
  );
};

export default App;
