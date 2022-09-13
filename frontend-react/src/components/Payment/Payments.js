import React from "react"
import PaymentSingle from "./PaymentSingle";
import "./Payments.css"
import { Elements } from "@stripe/react-stripe-js";
import { loadStripe } from "@stripe/stripe-js";
import { CheckoutForm } from "../Checkout/Checkout.js";

const stripePromise = loadStripe("pk_test_51LhMCXG0dnjPPiJiWgVK2u1ozxZs1Z71QQSCROSRrzvIdJHjTlEog0fI1HnhVCtItNaz915AFSWnxrjMmQctH3GN00XKtArsI0");

const Payments = ({cartItems, totalPrice, makePayment, preparePayment, clientSecret}) => { 
  
  if (clientSecret == "" || clientSecret == null) {
    preparePayment()
  }

  const appearance = {
      theme: 'stripe',
    };
    const options = {
      clientSecret,
      appearance,
    };
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
            {/* <button className="pay-button" onClick={() => makePay()}>Pay</button> */}
            </div>
          )}

      {clientSecret && (
        <Elements options={options} stripe={stripePromise}>
          <CheckoutForm makePayment={makePayment} />
        </Elements>
      )}
        </div>
        
      )
}

export default Payments;
  