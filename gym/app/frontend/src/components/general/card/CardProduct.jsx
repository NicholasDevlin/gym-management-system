import React, { useState } from 'react';
import Styles from './CardProduct.module.css'
import { Link } from 'react-router-dom'
import { Icon } from '@iconify/react';
import { useUserData } from '../../../utils/jwt/UserData.jsx';
import CurrencyFormat from '../../../utils/CurrencyFormat/CurrencyFormat.jsx';

function CardProduct({ title, duration, price, description, detail, deleteProduct }) {
  const [isCardMenuOpen, setCardMenuOpen] = useState(false);
  const { userData } = useUserData();

  const toggleCardMenu = () => {
    setCardMenuOpen(!isCardMenuOpen);
  };
  return (
    <div className={Styles.card}>
      <div className={Styles.content}>
        <div className={Styles.buttonContainer}>
          {userData.role === 'admin' && (
            <button className={Styles.button} onClick={toggleCardMenu}>
              <Icon icon="fluent:more-vertical-20-filled" width="1.5rem" height="1.5rem" />
            </button>
          )}
          {isCardMenuOpen && (
            <div className={Styles.cardMenu}>
              <ul>
                <Link to={detail}><li>Edit</li></Link>
                <li onClick={deleteProduct} >Delete</li>
              </ul>
            </div>
          )}
        </div>
        <div className={Styles.title}>{title}</div>
        <div className={Styles.icon}>

        </div>
        <div className={Styles.features}>
          <ul className='m-0'>
            <li>{duration}</li>
            <li><CurrencyFormat value={price} /></li>
          </ul>
        </div>
        <div className={`'row' ${Styles.description}`}>
          <p>{description}</p>
        </div>
      </div>
      <div className={Styles.bottom}>
        <Link className={Styles.btn}>Check it out</Link>
      </div>
    </div>
  )
}

export default CardProduct;