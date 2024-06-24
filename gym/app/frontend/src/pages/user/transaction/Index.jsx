import React, { useState, useEffect } from "react";
import Layout from '../../../layout/MainLayout/Layout.jsx'
import Card from '../../../components/general/card/Card.jsx'
import Styles from './Transaction.module.css'
import { API_URLS } from '../../../apiConfig.js'
import { useAlert } from "react-alert";
import { useUserData } from "../../../utils/jwt/UserData.jsx";
import CurrencyFormat from "../../../utils/CurrencyFormat/CurrencyFormat.jsx";

function Transaction() {
  const alert = useAlert()
  const { userData } = useUserData();
  const [transactionData, setTransaction] = useState([]);

  useEffect(() => {
    getTransaction();
  }, []);

  async function getTransaction() {
    try {
      const queryParams = new URLSearchParams({ memberUUID: userData.uuid });
      const response = await fetch(`${API_URLS.TRANSACTION}?${queryParams.toString()}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      const responseData = await response.json();
      if (responseData.success) {
        setTransaction(responseData.data);
      } else {
        alert.error("Get data unsuccessful");
      }
    } catch (error) {
      alert.error(`${error}`);
    }
  }

  function getTransactionByUserUUID(data, userUUID) {
    for (const transactionDetail of data.transactionDetail) {
      for (const transactionMemberDetail of transactionDetail.transactionMemberDetail) {
        if (transactionMemberDetail.userUUID === userUUID) {
          return transactionDetail;
        }
      }
    }
    return null;
  }

  const CardTitle = (data) => {
    let options = { day: 'numeric', month: 'long', year: 'numeric' };
    let date = new Date(data.transactionDate);
    let getTransactionDetail = getTransactionByUserUUID(data, userData.uuid);
    let getAdditionalPrice = getTransactionDetail.transactionMemberDetail.find(x => x.userUUID === userData.uuid);
    return (<>
      <div className="d-flex mx-3 justify-content-between">
        <div>
          <h6>Transaction No</h6>
          <h5>{data.transactionNo}</h5>
        </div>
        <div>
          <h6 className="text-end">Transaction Date</h6>
          <h5 className="text-end">{date.toLocaleDateString("id-ID", options)}</h5>
        </div>
      </div>
      <div className="d-flex mx-3 mt-3 justify-content-between">
        <div>
          <h6>Membership Plan</h6>
          <h5>{getTransactionDetail.membershipPlan.name}</h5>
        </div>
        <div>
          <h6 className="text-end">Price</h6>
          <h5><CurrencyFormat value={getAdditionalPrice.additionalPrice + getTransactionDetail.membershipPlan.price} /></h5>
        </div>
      </div>
    </>
    );
  }

  return (
    <Layout>
      <div className={Styles.cardHolder}>
        {transactionData.map((data) => (
          <Card title={CardTitle(data)} />
        ))}
      </div>
    </Layout>
  );
}

export default Transaction;