import React from "react";
import Layout from "../../../layout/MainLayout/Layout.jsx";
import TextField from "../../../components/general/input/inputTextField/TextField.jsx";
import NumericField from "../../../components/general/input/inputNumericField/NumericField.jsx";
import TextAreaField from "../../../components/general/input/inputTextAreaField/TextAreaField.jsx";
// import Button from '../../../components/general/button/Button.jsx'
// import Styles from './Membership.module.css'

function MembershipEditor() {
  return (
    <Layout>
      <div className="container my-5">
        <TextField id={"membershipPlan"} label={"Membership Plan"} />
        <NumericField id={"duration"} label={"Duration"} />
        <NumericField id={"price"} label={"Price"} />
        <TextAreaField id={"description"} label={"Description"} />
      </div>
    </Layout>
  );
}

export default MembershipEditor;
