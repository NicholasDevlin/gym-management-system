import { Modal } from "antd";
import Select from "../../../components/general/input/select/Select";
import DatetimePicker from "../../../components/general/input/datetimePicker/DatetimePicker";
import Button, { DangerButton } from "../../../components/general/button/Button";
import { API_URLS } from "../../../apiConfig";
import { useEffect, useState } from "react";
import { useAlert } from "react-alert";

export default function AbsensiModal({ isOpen, closeModal, userOptions, absensiData }) {
  const [absensi, setAbsensi] = useState({});
  const alert = useAlert();

  useEffect(() => {
    setAbsensi({
      uuid: absensiData != null ? absensiData.uuid : null,
      userUUID: absensiData != null ? absensiData.userUUID : null,
      date: absensiData != null ? absensiData.date : new Date()
    })
  }, [absensiData])

  async function saveAbsensi() {
    try {
      debugger
      const apiUrl = absensiData ? `${API_URLS.ABSENSI}/${absensiData.uuid}` : API_URLS.ABSENSI;
      const method = absensiData ? 'PUT' : 'POST';
      const response = await fetch(apiUrl, {
        method: method,
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer ' + localStorage.getItem('authToken'),
        },
        body: JSON.stringify(absensi),
      });

      const responseData = await response.json();

      if (responseData.success) {
        alert.success(absensiData ? 'Check in updated successfully' : 'Check in created successfully');
        closeModal();
      } else {
        throw new Error(`${responseData.message}`);
      }
    } catch (error) {
      alert.error(`${error}`);
    }
  }

  const handleMemberOnSelect = (e) => {
    setAbsensi({ ...absensi, userUUID: e });
  }

  const handleInputChange = (e) => {
    const { id, value } = e.target || e || {};
    setAbsensi((prevData) => ({
      ...prevData,
      [id]: value,
    }));
  };

  return (
    <>
      <Modal
        open={isOpen}
        onCancel={() => closeModal()}
        onClose={() => closeModal()}
        okButtonProps={{ className: "d-none" }}
        cancelButtonProps={{ className: "d-none" }}
      >
        <Select label="Member" options={userOptions} onSelect={handleMemberOnSelect} value={absensiData ? absensiData.userUUID : ""} />
        <DatetimePicker label="Date" onChange={handleInputChange} id="date" value={absensi.date} />
        <div className="d-flex justify-content-end">
          <DangerButton text="Cancel" onClick={() => closeModal()} />
          <Button className="ms-3" text="Save" onClick={() => saveAbsensi()} />
        </div>
      </Modal>
    </>
  );
}