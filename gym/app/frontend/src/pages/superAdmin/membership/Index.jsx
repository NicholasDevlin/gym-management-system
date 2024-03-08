import React, { useState, useEffect } from "react";
import Layout from '../../../layout/MainLayout/Layout.jsx'
import CardProduct from '../../../components/general/card/CardProduct.jsx'
import Button from '../../../components/general/button/Button.jsx'
import Styles from './Membership.module.css'
import { Link } from 'react-router-dom'
import { API_URLS } from '../../../apiConfig.js'
import { useAlert } from "react-alert";

function Membership() {
  const alert = useAlert()
  const [membershipPlanData, setMembershipPlanData] = useState([]);

  useEffect(() => {
    getMembershipPlan();
  }, []);

  async function getMembershipPlan() {
    try {
      const response = await fetch(`${API_URLS.MEMBERSHIP_PLAN}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const responseData = await response.json();
      if (responseData.success) {
        setMembershipPlanData(responseData.data);
      } else {
        alert.error("Get data unsuccessful");
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }
  return (
    <Layout>
      <div className={Styles.container}>
        <Link to='/membership/editor'><Button text={"Add new Membership Plan"} /></Link>
      </div>
      <div className={Styles.cardHolder}>
        {membershipPlanData.map((plan, index) => (
          <CardProduct title={plan.name} duration={plan.duration} price={plan.price} description={plan.description} />
        ))}
      </div>
    </Layout>
  );
}

export default Membership;