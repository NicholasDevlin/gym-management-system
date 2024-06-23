import React from 'react'
import Layout from '../../layout/MainLayout/Layout.jsx'
import Styles from './Home.module.css'
import { Carousel } from 'antd';
import gymInside from '../../assets/images/inside-gym.jpg'
import bfc from '../../assets/images/brayan-fitness-centre.jpeg'
import gymOutside from '../../assets/images/outside-gym.png'

function Home() {


  return (
    <Layout>
      <div className={`container-fluid bg-light ${Styles.container}`}>
        <Carousel autoplay >
          <div className={`py-3 ${Styles.imageHolder}`}>
            <img src={bfc} className={Styles.images} alt='brayan fitness centre' />
          </div>
          <div className={`py-3 ${Styles.imageHolder}`}>
            <img src={gymOutside} className={Styles.images} alt='gym outside' />
          </div>
          <div className={`py-3 ${Styles.imageHolder}`}>
            <img src={gymInside} className={Styles.images} alt='gym inside' />
          </div>
        </Carousel>
        <div className={`text-centre mx-5 py-3 px-5 ${Styles.content}`}>
          <h4>Selamat Datang di Brayan Fitness Centre: Gerbang Menuju Tubuh Sehat dan Bugar!</h4>
          <br />
          <h4>Apa itu Gym?</h4>
          <p className={Styles.content}>
            Gym, atau pusat kebugaran, adalah tempat yang menyediakan berbagai peralatan olahraga dan layanan untuk membantu Anda mencapai tujuan kebugaran Anda. Baik Anda ingin menurunkan berat badan, membangun otot, meningkatkan stamina, atau hanya ingin bersenang-senang dan berolahraga, gym memiliki sesuatu untuk semua orang.
          </p>
          <br />
          <h4>Jam Operasional:</h4>
          <div className='row'>
            <div className='col-4 text-dark'>
              <h5>Senin - Jumat</h5>
              <h5>Sabtu</h5>
              <h5>Minggu atau Tanggal merah</h5>
            </div>
            <div className='col-1 text-dark'>
              <h5>:</h5>
              <h5>:</h5>
              <h5>:</h5>
            </div>
            <div className='col-7 text-dark'>
              <h5>08:00 - 21:00</h5>
              <h5>08:00 - 19:00</h5>
              <h5>Tutup</h5>
            </div>
          </div>
          <br />
          <h4>Kontak:</h4>
          <div className='d-flex align-items-center mb-2'>
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="#000" class="bi bi-instagram" viewBox="0 0 16 16">
              <path d="M8 0C5.829 0 5.556.01 4.703.048 3.85.088 3.269.222 2.76.42a3.9 3.9 0 0 0-1.417.923A3.9 3.9 0 0 0 .42 2.76C.222 3.268.087 3.85.048 4.7.01 5.555 0 5.827 0 8.001c0 2.172.01 2.444.048 3.297.04.852.174 1.433.372 1.942.205.526.478.972.923 1.417.444.445.89.719 1.416.923.51.198 1.09.333 1.942.372C5.555 15.99 5.827 16 8 16s2.444-.01 3.298-.048c.851-.04 1.434-.174 1.943-.372a3.9 3.9 0 0 0 1.416-.923c.445-.445.718-.891.923-1.417.197-.509.332-1.09.372-1.942C15.99 10.445 16 10.173 16 8s-.01-2.445-.048-3.299c-.04-.851-.175-1.433-.372-1.941a3.9 3.9 0 0 0-.923-1.417A3.9 3.9 0 0 0 13.24.42c-.51-.198-1.092-.333-1.943-.372C10.443.01 10.172 0 7.998 0zm-.717 1.442h.718c2.136 0 2.389.007 3.232.046.78.035 1.204.166 1.486.275.373.145.64.319.92.599s.453.546.598.92c.11.281.24.705.275 1.485.039.843.047 1.096.047 3.231s-.008 2.389-.047 3.232c-.035.78-.166 1.203-.275 1.485a2.5 2.5 0 0 1-.599.919c-.28.28-.546.453-.92.598-.28.11-.704.24-1.485.276-.843.038-1.096.047-3.232.047s-2.39-.009-3.233-.047c-.78-.036-1.203-.166-1.485-.276a2.5 2.5 0 0 1-.92-.598 2.5 2.5 0 0 1-.6-.92c-.109-.281-.24-.705-.275-1.485-.038-.843-.046-1.096-.046-3.233s.008-2.388.046-3.231c.036-.78.166-1.204.276-1.486.145-.373.319-.64.599-.92s.546-.453.92-.598c.282-.11.705-.24 1.485-.276.738-.034 1.024-.044 2.515-.045zm4.988 1.328a.96.96 0 1 0 0 1.92.96.96 0 0 0 0-1.92m-4.27 1.122a4.109 4.109 0 1 0 0 8.217 4.109 4.109 0 0 0 0-8.217m0 1.441a2.667 2.667 0 1 1 0 5.334 2.667 2.667 0 0 1 0-5.334" />
            </svg>
            <h5 className='ms-2 text-dark my-0'> : <a href='https://www.instagram.com/brayanfitnesscentre?utm_source=ig_web_button_share_sheet&igsh=ZDNlZDc0MzIxNw=='>brayanfitnesscentre</a> </h5>
          </div>
          <div className='d-flex align-items-center'>
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="#000" class="bi bi-phone" viewBox="0 0 16 16">
              <path d="M11 1a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1zM5 0a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2z" />
              <path d="M8 14a1 1 0 1 0 0-2 1 1 0 0 0 0 2" />
            </svg>
            <h5 className='ms-2 text-dark my-0'> : 085277471888</h5>
          </div>
          <br />
          <h4>Alamat:</h4>
          <div className='d-flex justify-content-center'>
            <iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d672.8834356272563!2d98.66663594029978!3d3.628289391862243!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x3031320434c1e23f%3A0xfff90b1a725f4205!2sBrayan%20Fitness%20Centre!5e0!3m2!1sen!2sid!4v1719139977033!5m2!1sen!2sid" width="600" height="450" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>
          </div>
        </div>
      </div>
    </Layout>
  );
}

export default Home;