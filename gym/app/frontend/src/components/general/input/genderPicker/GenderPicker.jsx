import "react-datetime/css/react-datetime.css";
import Styles from "./GenderPicker.module.css";

function GenderPicker({ id, onChange, value }) {
  return (
    <div className={Styles.container}>
      <label>Gender</label>
      <div className="row">
        <input type="radio" name="gender" value="male" data-icon="" id={id} onChange={onChange} checked={value === 'male'} />
        <input type="radio" name="gender" value="female" data-icon="" id={id} onChange={onChange} checked={value === 'female'} />
      </div>
    </div>
  );
}
//source : https://codepen.io/morten-olsen/pen/QbvBYy
export default GenderPicker;
