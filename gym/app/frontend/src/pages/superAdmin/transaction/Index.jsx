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
import { formatDateTime, formatNumberWithCommas } from "../../../utils/CurrencyFormat/CurrencyFormat.jsx";
import * as XLSX from 'xlsx';

function Transaction() {
  const alert = useAlert()
  const [transactionData, setTransaction] = useState([]);
  const [filter, setFilter] = useState({
    transactionDateFrom: new Date(new Date().setDate(new Date().getDate() - 7)).toISOString(),
    transactionDateTo: new Date().toISOString()
  });

  useEffect(() => {
    getTransaction();
  }, []);

  async function getTransaction() {
    try {
      const queryParams = new URLSearchParams(filter);
      const response = await fetch(`${API_URLS.TRANSACTION}?${queryParams.toString()}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      const responseData = await response.json();
      if (responseData.success) {
        setTransaction(responseData.data ?? []);
      } else {
        throw new Error(responseData.message);
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }

  const handleInputChange = (e) => {
    const { id, value } = e.target || e || {};
    const originalDate = new Date(value.toString());
    const modifiedTime = new Date(
      originalDate.getFullYear(),
      originalDate.getMonth(),
      originalDate.getDate(),
      12,
      0,
      0
    );
    setFilter((prevData) => ({
      ...prevData,
      [id]: modifiedTime.toISOString(),
    }));
  };

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
      alert.error(`${error}`);
    }
  }

  const exportToExcel = () => {
    let data = [];

    transactionData.forEach((transaction) => {
      transaction.transactionDetail?.forEach((detail) => {
        detail.transactionMemberDetail?.forEach((member) => {
          data.push({
            "Transaction No": transaction.transactionNo,
            "Transaction Date": formatDateTime(transaction.transactionDate),
            "Status": transaction.status,
            "Member": member.user.name,
            "Membership Plan": detail.membershipPlan.name,
            "Price": formatNumberWithCommas(detail.price + (member.additionalPrice ?? 0)),
          });
        });
      });
    });

    const workbook = XLSX.utils.book_new();
    const worksheet = XLSX.utils.json_to_sheet(data);
    XLSX.utils.book_append_sheet(workbook, worksheet, 'Sheet1');
    XLSX.writeFile(workbook, "Transaksi.xlsx");
  };

  return (
    <Layout>
      <div className={Styles.container}>
        <div className="d-flex w-100">
          <div className="me-3">
            <DatetimePicker label="From" id="transactionDateFrom" onChange={handleInputChange} value={filter.transactionDateFrom} />
          </div>
          <div className="me-3">
            <DatetimePicker label="To" id="transactionDateTo" onChange={handleInputChange} value={filter.transactionDateTo} />
          </div>
          <div className="me-3 d-flex align-items-center mt-2">
            <Button text="Filter" onClick={() => getTransaction()} />
          </div>
        </div>
        <div className="d-flex">
          <Button className="me-2 px-3" onClick={exportToExcel} text={"Export to Excel"} />
          <Link to='/transaction/editor'><Button text={"Add new Transaction"} /></Link>
        </div>
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