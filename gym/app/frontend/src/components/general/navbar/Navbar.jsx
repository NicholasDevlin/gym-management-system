import React, { useState } from 'react';
import Styles from './Navbar.module.css'
import { Link } from 'react-router-dom'
import { Icon } from '@iconify/react';

function Navbar() {
  const [isProfileMenuOpen, setProfileMenuOpen] = useState(false);

  const toggleProfileMenu = () => {
    setProfileMenuOpen(!isProfileMenuOpen);
  };
  return (
    <nav className={Styles.navbar}>
      <div className={Styles.logo}><p>LOGO</p></div>
      <ul className={Styles.ul}>
        <Link to='/'><li className={Styles.li}>Home</li></Link>
        <Link to='/membership'><li className={Styles.li}>Membership</li></Link>
        <Link to='/help'><li className={Styles.li}>Help</li></Link>
        <li className={Styles.li} onClick={toggleProfileMenu}>
          <Icon icon="iconamoon:profile-circle-fill" color="#d8cdb9" width="32" height="32" />
          {isProfileMenuOpen && (
            <div className={Styles.profileMenu}>
              <ul>
                <li><Link to='/profile'>Profile</Link></li>
                <li>Sign out</li>
              </ul>
            </div>
          )}
        </li>
      </ul>
    </nav>
  );
}

export default Navbar