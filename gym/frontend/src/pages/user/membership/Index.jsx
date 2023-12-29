import React from 'react'
import Layout from '../../../layout/MainLayout/Layout.jsx'
import Card from '../../../components/general/card/Card.jsx'
import CardProduct from '../../../components/general/card/CardProduct.jsx'
import Button from '../../../components/general/button/Button.jsx'
import Styles from './Membership.module.css'

function Membership() {
  return (
    <Layout>
      <div className={Styles.container}>
        <Card title={"test"} body={"loremmmm "} button={Button} />
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