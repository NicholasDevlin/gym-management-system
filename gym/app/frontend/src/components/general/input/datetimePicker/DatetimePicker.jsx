import React from 'react';
import Datetime from 'react-datetime';
import "react-datetime/css/react-datetime.css";
import Styles from '../Input.module.css'

function DatetimePicker({ id, onChange, label }) {
  const handleChange = (selectedDate) => {
    onChange({ id, value: selectedDate });
  };

  return (
    <div className={Styles.formItem}>
      <label htmlFor={id}>{label}</label>
      <div className={Styles.inputWrapper}>
        <Datetime
          className={Styles.datetimePicker}
          inputProps={{ id: id }}
          onChange={handleChange}
        />
      </div>
    </div>
  );
}

export default DatetimePicker;
