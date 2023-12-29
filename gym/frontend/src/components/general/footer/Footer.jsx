import React from 'react'
import Styles from './Footer.module.css'
function Footer() {
  return (
    <footer className={Styles.footer}>
      <p>&copy;{new Date().getFullYear()} shychopath</p>
    </footer>
  );
}

export default Footer