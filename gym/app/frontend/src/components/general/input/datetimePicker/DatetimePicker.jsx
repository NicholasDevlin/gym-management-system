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
// Doc: https://github.com/arqex/react-datetime/blob/2a83208452ac5e41c43fea31ef47c65efba0bb56/README.md
export default DatetimePicker;