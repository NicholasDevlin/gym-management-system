import React, { useState } from 'react'
import Login from '../../components/login/Login.jsx'
import Register from '../../components/login/Register.jsx'

function Authentication() {
  const [activeTab, setActiveTab] = useState('login');

  const handleTabChange = (tab) => {
    setActiveTab(tab);
  };

  return (
    // <div className={Styles.container}>
    //   <div className={Styles.card}>
    //     <div className={Styles.nav}>
    //       <button className={Styles.tab} onClick={() => handleTabChange('register')}>Register</button>
    //       <button className={Styles.tab} onClick={() => handleTabChange('login')}>Login</button>
    //     </div>
        <div>
          {activeTab === 'register' && <Register loginOnClick={() => handleTabChange('login')} />}
          {activeTab === 'login' && <Login registerOnClick={() => handleTabChange('register')} />}
        </div>
    //   </div>
    // </div>
  );
}

export default Authentication;