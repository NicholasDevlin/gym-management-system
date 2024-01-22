import React from 'react'
import Styles from './Navbar.module.css'
import { Link } from 'react-router-dom'
import { Icon } from '@iconify/react';

function Navbar() {
  return (
    <nav className={Styles.navbar}>
      <div className={Styles.logo}><p>LOGO</p></div>
      <ul className={Styles.ul}>
        <Link to='/'><li className={Styles.li}>Home</li></Link>
        <Link to='/membership'><li className={Styles.li}>Membership</li></Link>
        <Link to='/help'><li className={Styles.li}>Help</li></Link>
        <Link to='/profile'><li className={Styles.li}><Icon icon="iconamoon:profile-circle-fill" color="#d8cdb9" width="32" height="32" /></li></Link>
      </ul>
    </nav>
  );
}

export default Navbar