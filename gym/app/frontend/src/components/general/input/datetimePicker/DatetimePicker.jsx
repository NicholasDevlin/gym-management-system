import React from 'react';
import Datetime from 'react-datetime';
import "react-datetime/css/react-datetime.css";
import Styles from '../Input.module.css'

function DatetimePicker({ id, onChange, label, value, className }) {
  const handleChange = (selectedDate) => {
    onChange({ id, value: selectedDate });
  };

  return (
    <div className={Styles.formItem}>
      {label? <label htmlFor={id}>{label}</label> : <></>}
      <div className={Styles.inputWrapper}>
        <Datetime
          dateFormat="DD MMM YYYY" timeFormat={false}
          value={value}
          initialValue={new Date()}
          className={`${Styles.datetimePicker} ${className}`}
          inputProps={{ id: id }}
          onChange={handleChange}
        />
      </div>
    </div>
  );
}

export default DatetimePicker;
