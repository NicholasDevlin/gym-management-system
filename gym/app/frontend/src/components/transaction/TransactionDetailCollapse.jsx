import Select from '../general/input/select/Select'
import Styles from './Transaction.module.css'
import NumericField from '../general/input/inputNumericField/NumericField.jsx'
import DeleteButton from '../general/button/DeleteButton.jsx'
import { useRef, useState } from 'react';
import { Icon } from '@iconify/react';
import PrefixSuffixNumericField from '../general/input/inputNumericField/PrefixSuffixNumberField.jsx';

export default function TransactionDetailCollapse({ detail }) {
  const detailRef = useRef(null);
  const [detailMember, setDetailMember] = useState([]);
  const [isCollapse, setIsCollapse] = useState(true);

  const handleDelete = () => {
    if (detailRef.current) {
      detailRef.current.remove();
    }
  };

  const options = [
    { name: 'Swedish', value: 'sv' },
    { name: 'English', value: 'en' },
    {
      type: 'group',
      name: 'Group name',
      items: [
        { name: 'Spanish', value: 'es' },
      ]
    },
  ];

  const addQty = () => {
    setDetailMember([...detailMember, {}]);
    setIsCollapse(false);
  }

  const substracQty = () => {
    setDetailMember(detailMember.slice(0, -1));
  }

  const toggleCollapse = () => {
    setIsCollapse(!isCollapse)
  }

  return (
    <div ref={detailRef} className="row py-2">
      <div className={`${'card bg-dark p-0'} ${Styles.card}`}>
        <div className="card-header row pe-1">
          <div className="col-md-3 col-sm-12">
            <Select label={"Membership Plan"} options={options} name={"membershipPlanUUID"} placeholder={"Choose Membership Plan..."} />
          </div>
          <div className="col-md-5 col-sm-6">
            <div className="row">
              <div className="col-md-4">
                <NumericField id={"price"} label={"Price"} />
              </div>
              <div className="col-md-4">
                <PrefixSuffixNumericField id={"qty"} label={"Qty"} prefixOnClick={substracQty} suffixOnClick={addQty} />
              </div>
              <div className="col-md-4">
                <NumericField id={"subtotal"} label={"Subtotal"} disabled={true} />
              </div>
            </div>
          </div>
          <div className="col-md-4 col-sm-6 d-flex justify-content-end pe-0 align-items-center">
            <DeleteButton onDelete={handleDelete} />
            <div className='h-100 d-flex align-items-center px-3 btn' onClick={toggleCollapse}>
              <Icon icon={`${isCollapse ? 'iconamoon:arrow-up-1' : 'iconamoon:arrow-down-1'}`} color="#d8cdb9" width="25" height="32" />
            </div>
          </div>
        </div>
        <div className={`card-body collapse ${isCollapse ? '' : 'show'}`}>
          {detailMember.map((detail) => (
            <Member />
          ))}
        </div>
      </div>
    </div>
  )
}

const Member = () => {
  const options = [
    { name: 'Swedish', value: 'sv' },
    { name: 'English', value: 'en' },
    {
      type: 'group',
      name: 'Group name',
      items: [
        { name: 'Spanish', value: 'es' },
      ]
    },
  ];
  return (
    <div className="row">
      <div className="col-6">
        <Select label={"Member"} options={options} name={"userUUID"} placeholder={"Choose Member"} />
      </div>
      <div className='col-6'>
        <NumericField id={"additionalPrice"} label={"Additional Price"} />
      </div>
    </div>
  )
}