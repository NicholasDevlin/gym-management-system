import React from 'react'
import Styles from './Footer.module.css'
function Footer() {
  return (
    <footer className={Styles.footer}>
      <p>Copyright {new Date().getFullYear()} &copy; shychopath</p>
    </footer>
  );
}

export default Footer