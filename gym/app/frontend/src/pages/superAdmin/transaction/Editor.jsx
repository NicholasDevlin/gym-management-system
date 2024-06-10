import React, { useState, useEffect, useCallback } from "react";
import { useParams } from 'react-router-dom';
import Layout from "../../../layout/MainLayout/Layout.jsx";
import Button from "../../../components/general/button/Button.jsx";
import { API_URLS } from "../../../apiConfig.js";
import { useAlert } from "react-alert";
import Styles from './Transaction.module.css'
import TransactionDetailCollapse from "../../../components/transaction/TransactionDetailCollapse.jsx";
import { GetUsers } from "../../../controller/UserController.js";
import { GetMembershipPlan } from "../../../controller/MembershipPlanController.js";
import NumericField from "../../../components/general/input/inputNumericField/NumericField.jsx";
import Select from "../../../components/general/input/select/Select.jsx";
import DatetimePicker from "../../../components/general/input/datetimePicker/DatetimePicker.jsx";
import TextField from "../../../components/general/input/inputTextField/TextField.jsx";

function TransactionEditor() {
  const alert = useAlert();
  const { uuid } = useParams();
  const [details, setDetails] = useState([]);
  const [isFetched, setIsFetched] = useState(false);
  const [userOptions, setUserOptions] = useState([]);
  const [membershipPlan, setMembershipPlan] = useState([]);
  const [transactionData, setTransactionData] = useState({
    transactionNo: "",
    transactionDate: new Date(),
    status: "",
    total: "",
    transactionDetail: []
  });

  const statusOptions = [
    {
      name: "Waiting for payment",
      value: "Waiting for payment"
    },
    {
      name: "Transaction complete",
      value: "Transaction complete"
    }
  ];

  const fetchUsersOptions = useCallback(async () => {
    const result = await GetUsers();
    if (result) {
      setUserOptions(result.map((data) => ({
        name: ` ${data.name} (${data.phoneNumber})`,
        value: data.uuid,
      })));
    }
  }, []);

  const fetchMembershipPlanOptions = useCallback(async () => {
    const result = await GetMembershipPlan();
    if (result) {
      setMembershipPlan(result);
    }
  }, []);

  useEffect(() => {
    if (uuid) {
      getTransaction();
    }
    if (!isFetched) {
      fetchMembershipPlanOptions();
      fetchUsersOptions();
      setIsFetched(true);
    }
  }, [uuid]);

  useEffect(() => {
    setTransactionData(prevData => ({
      ...prevData,
      transactionDetail: [...details]
    }));
  }, [details])

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
        setDetails(response.data.transactionDetail);
      } else {
        alert.error('Get data unsuccessful');
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }

  const handleInputChange = (e) => {
    const { id, value } = e.target || {};

    setTransactionData((prevData) => ({
      ...prevData,
      [id]: value,
    }));
  };

  async function saveTransaction() {
    debugger
    try {
      const apiUrl = uuid ? `${API_URLS.TRANSACTION}/${uuid}` : API_URLS.TRANSACTION;
      const method = uuid ? 'PUT' : 'POST';

      const response = await fetch(apiUrl, {
        method: method,
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer ' + localStorage.getItem('authToken'),
        },
        body: JSON.stringify(transactionData),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const responseData = await response.json();

      if (responseData.success) {
        alert.success(uuid ? 'Transaction updated successfully' : 'Transaction created successfully');
      } else {
        alert.error('Save unsuccessful');
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }

  const addTransactionDetail = () => {
    setDetails([...details, {}]);
  }

  return (
    <Layout>
      <div className="container my-5">
        <div className="row">
          <div className="col-6">
            <p className="text-start mb-0">Transaction No</p>
            <div className="d-flex justify-content-start">
              <TextField value={transactionData.transactionNo} disabled={true} />
            </div>
          </div>
          <div className="col-6">
            <p className="text-end mb-0">Transaction Date</p>
            <div className="d-flex justify-content-end">
              <DatetimePicker onChange={handleInputChange} id="transactionDate" className={Styles.textEnd} />
            </div>
          </div>
        </div>
        <div className="row">
          <div className="col-6">
            <p className="text-start mb-0">Status</p>
            <Select options={statusOptions} />
          </div>
          <div className="col-6">
            <p className="text-end mb-0">Total</p>
            <div className="d-flex justify-content-end">
              <NumericField className={`${Styles.dataHeader} text-end`} name="total" value={transactionData.total} disabled={true} />
            </div>
          </div>
        </div>
        <div className="d-flex justify-content-between my-3">
          <Button text={"+ Add Transaction Detail"} onClick={addTransactionDetail} />
          <Button text={"Save"} onClick={saveTransaction} />
        </div>
        <div className="my-3" id="transaction-details">
          {details.map((detail, index) => (
            <TransactionDetailCollapse detail={detail} index={index} setDetail={setDetails} userOptions={userOptions} membershipPlan={membershipPlan} />
          ))}
        </div>
      </div>
    </Layout >
  );
}

export default TransactionEditor;