import React from 'react'
import Layout from '../../../layout/MainLayout/Layout.jsx'
import CardProduct from '../../../components/general/card/CardProduct.jsx'
import Button from '../../../components/general/button/Button.jsx'
import Styles from './Membership.module.css'
import { Link } from 'react-router-dom'

function Membership() {
  return (
    <Layout>
      <div className={Styles.container}>
        <Link to='/membership/editor'><Button text={"Add new Membership Plan"}/></Link>
      </div>
      <div className={Styles.cardHolder}>
        <CardProduct title={"test"}></CardProduct>
        <CardProduct title={"test"}></CardProduct>
        <CardProduct title={"test"}></CardProduct>
      </div>
    </Layout>
  );
}

export default Membership;