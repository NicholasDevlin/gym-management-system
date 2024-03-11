import React, { useState, useEffect } from "react";
import { useParams } from 'react-router-dom'; 
import Layout from "../../../layout/MainLayout/Layout.jsx";
import TextField from "../../../components/general/input/inputTextField/TextField.jsx";
import NumericField from "../../../components/general/input/inputNumericField/NumericField.jsx";
import TextAreaField from "../../../components/general/input/inputTextAreaField/TextAreaField.jsx";
import Button from "../../../components/general/button/Button.jsx";
import { API_URLS } from "../../../apiConfig.js";
import { useAlert } from "react-alert";

function MembershipEditor() {
  const alert = useAlert();
  const { uuid } = useParams();

  const [membershipPlanData, setMembershipPlanData] = useState({
    name: '',
    duration: '',
    price: '',
    description: '',
  });

  useEffect(() => {
    if (uuid) {
      getMembershipPlan();
    }
  }, [uuid]);

  async function getMembershipPlan() {
    try {
      const response = await fetch(`${API_URLS.MEMBERSHIP_PLAN}/${uuid}`, {
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
        setMembershipPlanData(responseData.data);
        alert.success('Get data successful');
      } else {
        alert.error('Get data unsuccessful');
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }

  const handleInputChange = (e) => {
    const { id, value } = e.target || {};

    setMembershipPlanData((prevData) => ({
      ...prevData,
      [id]: id === 'duration' || id === 'price' ? parseInt(value) : value,
    }));
  };

  async function saveMembershipPlan() {
    try {
      const apiUrl = uuid ? `${API_URLS.MEMBERSHIP_PLAN}/${uuid}` : API_URLS.MEMBERSHIP_PLAN;
      const method = uuid ? 'PUT' : 'POST';

      const response = await fetch(apiUrl, {
        method: method,
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer ' + localStorage.getItem('authToken'),
        },
        body: JSON.stringify(membershipPlanData),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const responseData = await response.json();

      if (responseData.success) {
        alert.success(uuid ? 'Membership plan updated successfully' : 'Membership plan created successfully');
      } else {
        alert.error('Save unsuccessful');
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }

  return (
    <Layout>
      <div className="container my-5">
        <TextField id={'name'} value={membershipPlanData.name || ''} onChange={handleInputChange} label={'Membership Plan'} />
        <NumericField id={'duration'} onChange={handleInputChange} value={membershipPlanData.duration} label={'Duration'} />
        <NumericField id={'price'} onChange={handleInputChange} value={membershipPlanData.price} label={'Price'} />
        <TextAreaField id={'description'} value={membershipPlanData.description} onChange={handleInputChange} label={'Description'} />
        <div className="d-flex justify-content-end">
          <Button onClick={saveMembershipPlan} text={'Save'} />
        </div>
      </div>
    </Layout>
  );
}

export default MembershipEditor;