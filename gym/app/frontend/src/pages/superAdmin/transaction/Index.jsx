import React, { useState, useEffect } from "react";
import Layout from '../../../layout/MainLayout/Layout.jsx'
import Button from '../../../components/general/button/Button.jsx'
import Styles from './Transaction.module.css'
import { API_URLS } from '../../../apiConfig.js'
import { Link } from 'react-router-dom'
import { useAlert } from "react-alert";
import { Table } from "antd";
import column from "./config.jsx";
import DatetimePicker from "../../../components/general/input/datetimePicker/DatetimePicker.jsx";

function Transaction() {
  const alert = useAlert()
  const [transactionData, setTransaction] = useState([]);

  useEffect(() => {
    getTransaction();
  });

  async function getTransaction() {
    try {
      const response = await fetch(`${API_URLS.TRANSACTION}`, {
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
        debugger
        setTransaction(responseData.data);
      } else {
        alert.error("Get data unsuccessful");
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }


  async function deleteTransaction(uuid) {
    try {
      const response = await fetch(`${API_URLS.TRANSACTION}/${uuid}`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}\n${response.message}`);
      }

      const responseData = await response.json();
      if (responseData.success) {
        const newData = transactionData.filter((item) => item.uuid !== uuid);
        setTransaction(newData);
        alert.success("Delete Successful!");
      } else {
        alert.error("Delete Transaction unsuccessful");
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }

  return (
    <Layout>
      <div className={Styles.container}>
        <div className="d-flex w-100">
          <div className="me-3">
            <DatetimePicker label="From" />
          </div>
          <div className="me-3">
            <DatetimePicker label="To" />
          </div>
          <div className="me-3 d-flex align-items-center mt-2">
            <Button text="Filter" />
          </div>
        </div>
        <Link to='/transaction/editor'><Button text={"Add new Transaction"} /></Link>
      </div>
      <div>
        <Table
          columns={column(deleteTransaction)}
          dataSource={transactionData}
          className="h-100 m-3"
        />
      </div>
    </Layout>
  );
}

export default Transaction;