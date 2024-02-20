import React from 'react'
import Layout from '../../../layout/MainLayout/Layout.jsx'
import CardProduct from '../../../components/general/card/CardProduct.jsx'
import Button from '../../../components/general/button/Button.jsx'
import Styles from './Membership.module.css'

function Membership() {
  return (
    <Layout>
      <div className={Styles.container}>
        <Button text={"Add new Membership Plan"}/>
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