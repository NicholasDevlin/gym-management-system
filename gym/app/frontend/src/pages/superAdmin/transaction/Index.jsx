import React, { useState, useEffect } from "react";
import Layout from '../../../layout/MainLayout/Layout.jsx'
import Card from '../../../components/general/card/Card.jsx'
import Button from '../../../components/general/button/Button.jsx'
import Styles from './Transaction.module.css'
import { API_URLS } from '../../../apiConfig.js'
import { useAlert } from "react-alert";

function Transaction() {
  const alert = useAlert()
  const [transactionData, setTransaction] = useState([]);

  useEffect(() => {
    getTransaction();
  }, []);

  async function getTransaction() {
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
        setTransaction(responseData.data);
      } else {
        alert.error("Get data unsuccessful");
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }
  return (
    <Layout>
      <div className={Styles.cardHolder}>
        {/* {transactionData.map((plan, index) => (

        ))} */}
      </div>
    </Layout>
  );
}

export default Transaction;