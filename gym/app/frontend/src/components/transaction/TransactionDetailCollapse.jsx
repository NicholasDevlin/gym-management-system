import Select from '../general/input/select/Select'
import Styles from './Transaction.module.css'
import NumericField from '../general/input/inputNumericField/NumericField.jsx'
import DeleteButton from '../general/button/DeleteButton.jsx'
import { useEffect, useRef, useState } from 'react';
import { Icon } from '@iconify/react';
import PrefixSuffixNumericField from '../general/input/inputNumericField/PrefixSuffixNumberField.jsx';

export default function TransactionDetailCollapse({ detail, index, setDetail, userOptions, membershipPlan }) {
  const detailRef = useRef(null);
  const [detailMembers, setDetailMembers] = useState([]);
  const [isCollapse, setIsCollapse] = useState(true);
  const [qty, setQty] = useState(detail.qty || 0);
  const [price, setPrice] = useState(0);
  const [membershipPlanOptions, setMembershipPlanOptions] = useState([]);

  const handleDelete = () => {
    if (detailRef.current) {
      detailRef.current.classList.add("d-none");
      setDetail(prevItems => {
        const newItems = [...prevItems];
        newItems[index] = { ...newItems[index], deleted: true };
        return newItems;
      });
    }
  };

  const handleInputChange = (e) => {
    const { id, value } = e.target || {};

    setDetail(prevItems => {
      const newItems = [...prevItems];
      newItems[index] = { ...newItems[index], [id]: value };
      return newItems;
    });
  };

  useEffect(() => {
    if (membershipPlan) {
      setMembershipPlanOptions(membershipPlan.map((data) => ({
        name: `${data.name} (${data.duration} Days)`,
        value: data.uuid,
      })));
    }
  }, [membershipPlan])

  useEffect(() => {
    setDetail(prevItems => {
      const newItems = [...prevItems];
      newItems[index] = { ...newItems[index], price: price, qty: qty, transactionMemberDetail: detailMembers };
      return newItems;
    });
  }, [detailMembers, qty, price]);

  const addQty = () => {
    setDetailMembers([...detailMembers, {}]);
    setIsCollapse(false);
  }

  const substracQty = () => {
    setDetailMembers(detailMembers.slice(0, -1));
  }

  const toggleCollapse = () => {
    setIsCollapse(!isCollapse)
  }

  const handleQtyChange = (value) => {
    setQty(value);
  }

  const handleMembershipOnSelect = (e) => {
    let selectedMembershipPlan = membershipPlan.find(x => x.uuid === e);
    setPrice(selectedMembershipPlan.price);
    setDetail(prevItems => {
      const newItems = [...prevItems];
      newItems[index] = { ...newItems[index], membershipPlanUUID: e };
      return newItems;
    });
  }

  return (
    <div ref={detailRef} className="row py-2">
      <div className={`${'card bg-dark p-0'} ${Styles.card}`}>
        <div className="card-header row pe-1">
          <div className="col-md-3 col-sm-12">
            <Select label={"Membership Plan"} options={membershipPlanOptions} name={"membershipPlanUUID"} placeholder={"Choose Membership Plan..."} onSelect={handleMembershipOnSelect} />
          </div>
          <div className="col-md-5 col-sm-6">
            <div className="row">
              <div className="col-md-4">
                <NumericField id={"price"} label={"Price"} value={price} onChange={(e) => {
                  setPrice(e.target.value);
                  handleInputChange(e);
                }} />
              </div>
              <div className="col-md-4">
                <PrefixSuffixNumericField id={"qty"} label={"Qty"} onChange={handleQtyChange} value={qty} prefixOnClick={substracQty} suffixOnClick={addQty} />
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
          {detailMembers.map((detailMember, index) => (
            <Member qty={qty} setQty={setQty} userOptions={userOptions} detailMember={detailMember} setDetailMember={setDetailMembers} index={index} />
          ))}
        </div>
      </div>
    </div>
  )
}

const Member = ({ qty, setQty, userOptions, setDetailMember, detailMember, index }) => {
  const memberRef = useRef(null);

  const handleDelete = () => {
    if (memberRef.current) {
      memberRef.current.classList.add("d-none");
      setQty(qty - 1);
      setDetailMember(prevItems => {
        const newItems = [...prevItems];
        newItems[index].deleted = true;
        return newItems;
      });
    }
  };

  const handleInputChange = (e) => {
    const { id, value } = e.target || {};

    setDetailMember(prevItems => {
      const newItems = [...prevItems];
      newItems[index] = { ...newItems[index], [id]: parseInt(value) };
      return newItems;
    });
  };

  const handleOnSelect = (v) => {
    setDetailMember(prevItems => {
      const newItems = [...prevItems];
      newItems[index] = { ...newItems[index], userUUID: v };
      return newItems;
    });
  }

  return (
    <div ref={memberRef} className="row">
      <div className="col-3">
        <div className='row'>
          <Select label={"Member"} options={userOptions} onSelect={handleOnSelect} name={"userUUID"} placeholder={"Choose Member"} />
        </div>
      </div>
      <div className='col-5'>
        <NumericField id={"additionalPrice"} onChange={handleInputChange} value={detailMember.additionalPrice || 0} label={"Additional Price"} />
      </div>
      <div className='col-4 d-flex align-items-center justify-content-end'>
        <DeleteButton onDelete={handleDelete} />
      </div>
    </div>
  )
}