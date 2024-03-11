import React, { useState } from 'react';
import Styles from './CardProduct.module.css'
import { Link } from 'react-router-dom'
import { Icon } from '@iconify/react';

function CardProduct({ title, duration, price, description, detail }) {
  const [isCardMenuOpen, setCardMenuOpen] = useState(false);

  const toggleCardMenu = () => {
    setCardMenuOpen(!isCardMenuOpen);
  };
  return (
    <div className={Styles.card}>
      <div className={Styles.content}>
        <div className={Styles.buttonContainer}>
          <button className={Styles.button} onClick={toggleCardMenu}>
            <Icon icon="fluent:more-vertical-20-filled" width="1.5rem" height="1.5rem" />
          </button>
          {isCardMenuOpen && (
            <div className={Styles.cardMenu}>
              <ul>
                <Link to={detail}><li>Edit</li></Link>
                <li>Delete</li>
              </ul>
            </div>
          )}
        </div>
        <div className={Styles.title}>{title}</div>
        <div className={Styles.icon}>

        </div>
        <div className={Styles.features}>
          <ul>
            <li>{duration}</li>
            <li>{price}</li>
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