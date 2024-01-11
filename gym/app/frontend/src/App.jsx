import React from 'react'
import Layout from './layout/MainLayout/Layout.jsx'
import DeleteButton from './components/general/button/DeleteButton.jsx'

function App() {
  return (
    <Layout>
      <div>Home</div>
      <DeleteButton />
    </Layout>
  );
}

export default App;
