import React from "react"
import PaymentSingle from "./PaymentSingle";
import "./Payments.css"

const Payments = ({cartItems, totalPrice, makePay}) => {
    return (
        <div className="cart-items">
          
          {cartItems.length === 0 && (
              <div className="cart-items-empty"> No items added</div>
          )}
  
          {cartItems.map((cartItem) => (
                <PaymentSingle cartItem={cartItem}/>
            ))}
  
          {cartItems.length >= 1 && (<div>
            <h2> Total price: {totalPrice} zl</h2> 
            <button className="pay-button" onClick={() => makePay()}>Pay</button>
            </div>
          )}
        </div>
        
      )
}

export default Payments;
  