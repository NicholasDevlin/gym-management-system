import React, { useState } from 'react'
import Login from '../../components/login/Login.jsx'
import Register from '../../components/login/Register.jsx'

function Authentication() {
  const [activeTab, setActiveTab] = useState('login');

  const handleTabChange = (tab) => {
    setActiveTab(tab);
  };

  return (
    <>
      {activeTab === 'register' && <Register loginOnClick={() => handleTabChange('login')} />}
      {activeTab === 'login' && <Login registerOnClick={() => handleTabChange('register')} />}
    </>
  );
}

export default Authentication;