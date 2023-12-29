import React from 'react'
import Navbar from '../../components/general/navbar/Navbar.jsx'
import Footer from '../../components/general/footer/Footer.jsx'
import Styles from './Layout.module.css'

function MainLayout({ children }) {
  return (
    <div className={Styles.parent}>
      <div className={Styles.content}>
        <Navbar title="home" />
        {children}
      </div>
      <div className={Styles.footer}>
        <Footer />
      </div>
    </div>
  );
}

export default MainLayout