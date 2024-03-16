import React, { useState, useEffect } from "react";
import { useParams } from 'react-router-dom';
import Layout from "../../../layout/MainLayout/Layout.jsx";
import Button from "../../../components/general/button/Button.jsx";
import { API_URLS } from "../../../apiConfig.js";
import { useAlert } from "react-alert";
import Styles from './Transaction.module.css'
import Select from "../../../components/general/input/select/Select.jsx";

function TransactionEditor() {
  const alert = useAlert();
  const { uuid } = useParams();

  const [transactionData, setTransactionData] = useState({
    transactionNo: '',
    transactionDate: '',
    status: '',
    total: ''
  });

  useEffect(() => {
    if (uuid) {
      getTransaction();
    }
  }, [uuid]);

  async function getTransaction() {
    try {
      const response = await fetch(`${API_URLS.TRANSACTION}/${uuid}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer ' + localStorage.getItem('authToken'),
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const responseData = await response.json();
      if (responseData.success) {
        setTransactionData(responseData.data);
      } else {
        alert.error('Get data unsuccessful');
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }

  // const handleInputChange = (e) => {
  //   const { id, value } = e.target || {};

  //   setTransactionData((prevData) => ({
  //     ...prevData,
  //     [id]: id === 'duration' || id === 'price' ? parseInt(value) : value,
  //   }));
  // };

  // async function saveMembershipPlan() {
  //   try {
  //     const apiUrl = uuid ? `${API_URLS.MEMBERSHIP_PLAN}/${uuid}` : API_URLS.MEMBERSHIP_PLAN;
  //     const method = uuid ? 'PUT' : 'POST';

  //     const response = await fetch(apiUrl, {
  //       method: method,
  //       headers: {
  //         'Content-Type': 'application/json',
  //         Authorization: 'Bearer ' + localStorage.getItem('authToken'),
  //       },
  //       body: JSON.stringify(transactionData),
  //     });

  //     if (!response.ok) {
  //       throw new Error(`HTTP error! Status: ${response.status}`);
  //     }

  //     const responseData = await response.json();

  //     if (responseData.success) {
  //       alert.success(uuid ? 'Membership plan updated successfully' : 'Membership plan created successfully');
  //     } else {
  //       alert.error('Save unsuccessful');
  //     }
  //   } catch (error) {
  //     alert.error(`Error: ${error}`);
  //   }
  // }

  const options = [
    { name: 'Swedish', value: 'sv' },
    { name: 'English', value: 'en' },
    {
      type: 'group',
      name: 'Group name',
      items: [
        { name: 'Spanish', value: 'es' },
      ]
    },
  ];

  return (
    <Layout>
      <div className="container my-5">
        <div className="row">
          <div className="row">
            <div className="col-6">
              <p className="text-start mb-0">Transaction No</p>
              <p className="text-start"><h4 className={Styles.dataHeader}>{transactionData.transactionNo} test</h4></p>
            </div>
            <div className="col-6">
              <p className="text-end mb-0">Transaction Date</p>
              <p className="text-end"><h4 className={Styles.dataHeader}>{transactionData.transactionDate} test</h4></p>
            </div>
          </div>
          <div className="row">
            <div className="col-6">
              <p className="text-start mb-0">Status</p>
              <p className="text-start"><h4 className={Styles.dataHeader}>{transactionData.status} test</h4></p>
            </div>
            <div className="col-6">
              <p className="text-end mb-0">Total</p>
              <p className="text-end"><h4 className={Styles.dataHeader}>{transactionData.total} test</h4></p>
            </div>
          </div>
        </div>
        <div className="row">
          <div className="card">
            <div className="card-header">
              <Select label={"Membership Plan"} options={options} name={"language"} placeholder={"Choose a language..."} />
            </div>
          </div>
        </div>
        <div className="d-flex justify-content-end">
          {/* <Button onClick={saveMembershipPlan} text={'Save'} /> */}
        </div>
      </div>
    </Layout>
  );
}

export default TransactionEditor;