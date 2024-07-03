import React, { useState } from 'react';
import Styles from './Navbar.module.css';
import { Link } from 'react-router-dom';
import { Icon } from '@iconify/react';
import { useUserData } from '../../../utils/jwt/UserData.jsx'; // Import useUserData hook

function Navbar() {
  const { handleLogout, userData } = useUserData();
  const [isProfileMenuOpen, setProfileMenuOpen] = useState(false);

  const toggleProfileMenu = () => {
    setProfileMenuOpen(!isProfileMenuOpen);
  };

  return (
    <nav className={Styles.navbar}>
      <div className={Styles.logo}><p className={Styles.bfc}>Brayan Fitness Centre</p><p className={Styles.smallBfc}>BFC</p></div>
      <ul className={Styles.ul}>
        <Link to='/'><li className={Styles.li}>Home</li></Link>
        <Link to='/membership'><li className={Styles.li}>Membership Plan</li></Link>
        {userData ?
          <>
            <Link to='/transaction'><li className={Styles.li}>Transaction</li></Link>
          </>
          :
          <></>}
        {userData && userData.role === "admin" ?
          <>
            <Link to='/absensi'><li className={Styles.li}>Check in</li></Link>
            <Link to='/user'><li className={Styles.li}>Users</li></Link>
          </>
          :
          <></>
        }
        <li className={Styles.li} onClick={toggleProfileMenu}>
          <Icon icon="iconamoon:profile-circle-fill" color="#d8cdb9" width="32" height="32" />
          {isProfileMenuOpen && (
            <div className={Styles.profileMenu}>
              <ul>
                {userData ? (<>
                  <Link to='/profile'><li>Profile</li></Link>
                  <Link to="/authentication"><li onClick={handleLogout}>Sign out</li></Link>
                </>
                ) : (
                  <Link to="/authentication"><li>Sign in</li></Link>
                )}
              </ul>
            </div>
          )}
        </li>
      </ul>
    </nav>
  );
}

export default Navbar;
