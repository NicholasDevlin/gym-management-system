import React, { useState, useEffect, useCallback } from "react";
import Layout from '../../../layout/MainLayout/Layout.jsx'
import Button from '../../../components/general/button/Button.jsx'
import Styles from './Absensi.module.css'
import { API_URLS } from '../../../apiConfig.js'
import { useAlert } from "react-alert";
import { Table } from "antd";
import column from "./Config.jsx";
import DatetimePicker from "../../../components/general/input/datetimePicker/DatetimePicker.jsx";
import AbsensiModal from "./AbsensiModal.jsx";
import { GetUsers } from "../../../controller/UserController.js";
import { formatDateTime, formatNumberWithCommas } from "../../../utils/CurrencyFormat/CurrencyFormat.jsx";
import * as XLSX from 'xlsx';

function Transaction() {
  const alert = useAlert()
  const [absensiData, setAbsensi] = useState([]);
  const [userOptions, setUserOptions] = useState([]);
  const [selected, setSelected] = useState(null);
  const [isOpen, setOpen] = useState(false);
  const [filter, setFilter] = useState({
    dateFrom: new Date(new Date().setDate(new Date().getDate() - 7)).toISOString(),
    dateTo: new Date().toISOString()
  });

  useEffect(() => {
    getAbsensi();
    fetchUsersOptions();
  }, []);

  const fetchUsersOptions = useCallback(async () => {
    const result = await GetUsers({active: true});
    if (result) {
      setUserOptions(
        [
          { name: "Please Select Member", value: "" },
          ...result.map((data) => ({
            name: ` ${data.name} (${data.phoneNumber})`,
            value: data.uuid,
          }))
        ]
      );
    }
  }, []);

  async function getAbsensi() {
    try {
      const queryParams = new URLSearchParams(filter);
      const response = await fetch(`${API_URLS.ABSENSI}?${queryParams.toString()}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      const responseData = await response.json();
      if (responseData.success) {
        setAbsensi(responseData.data);
      } else {
        throw new Error(responseData.message);
      }
    } catch (error) {
      alert.error(`${error}`);
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

  async function deleteAbsensi(uuid) {
    try {
      const response = await fetch(`${API_URLS.ABSENSI}/${uuid}`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
      });

      const responseData = await response.json();
      if (responseData.success) {
        const newData = absensiData.filter((item) => item.uuid !== uuid);
        setAbsensi(newData);
        alert.success("Delete Successful!");
      } else {
        alert.error("Delete Absensi unsuccessful");
      }
    } catch (error) {
      alert.error(`${error}`);
    }
  }

  const closeModal = () => {
    setOpen(false);
    getAbsensi();
    setSelected(null);
  }

  useEffect(() => {
    if (selected) {
      setOpen(true);
    }
  }, [selected])

  async function checkOutAbsensi(data) {
    try {
      data.checkOut = new Date();
      const response = await fetch(`${API_URLS.ABSENSI}/${data.uuid}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer ' + localStorage.getItem('authToken'),
        },
        body: JSON.stringify(data),
      });

      const responseData = await response.json();

      if (responseData.success) {
        alert.success('Check Out successfully');
        getAbsensi();
      } else {
        throw new Error(`${responseData.message}`);
      }
    } catch (error) {
      alert.error(`${error}`);
    }
  }

  const exportToExcel = () => {
    let data = [];

    absensiData.forEach((absensi) => {
      data.push({
        "Member": absensi.user.name,
        "Check In": formatDateTime(absensi.date),
        "Check Out": formatDateTime(absensi.checkOut),
      });
    });

    const workbook = XLSX.utils.book_new();
    const worksheet = XLSX.utils.json_to_sheet(data);
    XLSX.utils.book_append_sheet(workbook, worksheet, 'Sheet1');
    XLSX.writeFile(workbook, "CheckIn.xlsx");
  };

  return (
    <Layout>
      <AbsensiModal userOptions={userOptions} closeModal={closeModal} isOpen={isOpen} absensiData={selected} />
      <div className={Styles.container}>
        <div className="d-flex w-100">
          <div className="me-3">
            <DatetimePicker label="From" id="dateFrom" onChange={handleInputChange} value={filter.dateFrom} />
          </div>
          <div className="me-3">
            <DatetimePicker label="To" id="dateTo" onChange={handleInputChange} value={filter.dateTo} />
          </div>
          <div className="me-3 d-flex align-items-center mt-2">
            <Button text="Filter" onClick={() => getAbsensi()} />
          </div>
        </div>
        <Button className="me-2 px-3" onClick={exportToExcel} text={"Export to Excel"} />
        <Button text={"Check in"} onClick={() => setOpen(true)} />
      </div>
      <div>
        <Table
          columns={column(deleteAbsensi, setSelected, checkOutAbsensi)}
          dataSource={absensiData}
          className="h-100 m-3"
        />
      </div>
    </Layout>
  );
}

export default Transaction;