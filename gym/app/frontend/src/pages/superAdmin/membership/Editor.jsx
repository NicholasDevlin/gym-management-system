import React, { useState, useEffect } from "react";
import Layout from "../../../layout/MainLayout/Layout.jsx";
import TextField from "../../../components/general/input/inputTextField/TextField.jsx";
import NumericField from "../../../components/general/input/inputNumericField/NumericField.jsx";
import TextAreaField from "../../../components/general/input/inputTextAreaField/TextAreaField.jsx";
import Button from "../../../components/general/button/Button.jsx";
import { API_URLS } from "../../../apiConfig.js";
import { useAlert } from "react-alert";

function MembershipEditor() {
  const alert = useAlert()
  const [membershipPlanData, setMembershipPlanData] = useState({
    name: "",
    duration: "",
    price: "",
    description: "",
  });

  const handleInputChange = (e) => {
    const { id, value } = e.target || {};

    setMembershipPlanData((prevData) => ({
      ...prevData,
      [id]: id === 'duration' || id === 'price' ? parseInt(value) : value,
    }));
  };

  async function saveMembershipPlan() {
    try {
      const response = await fetch(`${API_URLS.MEMBERSHIP_PLAN}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": "Bearer " + localStorage.getItem('authToken')
        },
        body: JSON.stringify(membershipPlanData),
      });

      debugger;
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const responseData = await response.json();

      if (responseData.success) {
        alert.success("Membership plan created successful")
      } else {
        alert.error("Save unsuccessful");
      }
    } catch (error) {
      alert.error(`Error: ${error}`);
    }
  }

  return (
    <Layout>
      <div className="container my-5">
        <TextField id={"name"} value={membershipPlanData.name || ''} onChange={handleInputChange} label={"Membership Plan"} />
        <NumericField id={"duration"} onChange={handleInputChange} value={membershipPlanData.duration} label={"Duration"} />
        <NumericField id={"price"} onChange={handleInputChange} value={membershipPlanData.price} label={"Price"} />
        <TextAreaField id={"description"} value={membershipPlanData.description} onChange={handleInputChange} label={"Description"} />
        <div className="d-flex justify-content-end">
          <Button onClick={saveMembershipPlan} text={"Save"} />
        </div>
      </div>
    </Layout>
  );
}

export default MembershipEditor;
