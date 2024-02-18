import Datetime from 'react-datetime';
import "react-datetime/css/react-datetime.css";
import Styles from './DatetimePicker.module.css'

function DatetimePicker({ id, onChange, label }) {
  return (
    <div className={Styles.formItem}>
      <label htmlFor={id}>{label}</label>
      <div className={Styles.inputWrapper}>
        <Datetime className={Styles.datetimePicker} inputProps={{ id: id, onChange: onChange }} />
      </div>
    </div>
  )
}

export default DatetimePicker;