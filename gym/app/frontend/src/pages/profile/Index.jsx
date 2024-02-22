import React from 'react'
import Layout from '../../layout/MainLayout/Layout.jsx';
import Styles from './Profile.module.css'
import { Icon } from '@iconify/react';

function Profile() {
  return (
    <Layout>
      <div className='d-flex justify-content-center mt-4'>
        <div className={Styles.imageContainer}>
          <img alt='profile' className={`rounded-circle ${Styles.image}`} src='https://w7.pngwing.com/pngs/79/184/png-transparent-mannequin-head-dummy-model-face-male-fashion-bold-thumbnail.png' />
          <div className={Styles.iconContainer}>
            <label htmlFor='inputFile'>
              <Icon className={Styles.icon} icon="mdi:pencil-outline" width="25" height="25" />
            </label>
            <input id='inputFile' type="file" className='d-none' />
          </div>
        </div>
      </div>
    </Layout>
  );
}

export default Profile;